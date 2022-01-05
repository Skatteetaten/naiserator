package aiven

import (
	"fmt"

	aiven_nais_io_v1 "github.com/nais/liberator/pkg/apis/aiven.nais.io/v1"
	nais_io_v1 "github.com/nais/liberator/pkg/apis/nais.io/v1"
	nais_io_v1alpha1 "github.com/nais/liberator/pkg/apis/nais.io/v1alpha1"
	"github.com/nais/liberator/pkg/namegen"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/validation"

	"github.com/nais/naiserator/pkg/resourcecreator/pod"
	"github.com/nais/naiserator/pkg/resourcecreator/resource"
)

const (
	aivenCredentialFilesVolumeName = "aiven-credentials"
)

type Source interface {
	resource.Source
	GetInflux() *nais_io_v1.Influx
	GetKafka() *nais_io_v1.Kafka
	GetElastic() *nais_io_v1.Elastic
	GetOpenSearch() *nais_io_v1.OpenSearch
}

type Config interface {
	IsKafkaratorEnabled() bool
}

func generateAivenSecretName(name string) string {
	secretName := namegen.RandShortName(fmt.Sprintf("aiven-%s", name), validation.DNS1035LabelMaxLength)

	return secretName
}

func Create(source Source, ast *resource.Ast, config Config) error {
	secretName := generateAivenSecretName(source.GetName())
	aivenApp := aiven_nais_io_v1.NewAivenApplicationBuilder(source.GetName(), source.GetNamespace()).
		WithSpec(aiven_nais_io_v1.AivenApplicationSpec{
			SecretName: secretName,
		}).
		Build()
	aivenApp.ObjectMeta = resource.CreateObjectMeta(source)

	Influx(ast, source.GetInflux(), &aivenApp)
	kafkaKeyPaths := Kafka(source, ast, config, source.GetKafka(), &aivenApp)

	elasticEnabled, err := Elastic(ast, source.GetElastic(), &aivenApp)
	if err != nil {
		return err
	}

	openSearchEnabled, err := OpenSearch(ast, source.GetOpenSearch(), &aivenApp)
	if err != nil {
		return err
	}

	if len(kafkaKeyPaths) > 0 {
		credentialFilesVolume := pod.FromFilesSecretVolume(aivenCredentialFilesVolumeName, secretName, kafkaKeyPaths)

		ast.Volumes = append(ast.Volumes, credentialFilesVolume)
		ast.VolumeMounts = append(ast.VolumeMounts, pod.FromFilesVolumeMount(credentialFilesVolume.Name, nais_io_v1alpha1.DefaultKafkaratorMountPath, "", true))
	}

	if len(kafkaKeyPaths) > 0 || elasticEnabled || openSearchEnabled {
		ast.AppendOperation(resource.OperationCreateOrUpdate, &aivenApp)
	}
	return nil
}

func makeSecretEnvVar(key, secretName string) v1.EnvVar {
	return v1.EnvVar{
		Name: key,
		ValueFrom: &v1.EnvVarSource{
			SecretKeyRef: &v1.SecretKeySelector{
				LocalObjectReference: v1.LocalObjectReference{
					Name: secretName,
				},
				Key: key,
			},
		},
	}
}
