package maskinporten

import (
	nais_io_v1 "github.com/nais/liberator/pkg/apis/nais.io/v1"
	nais_io_v1alpha1 "github.com/nais/liberator/pkg/apis/nais.io/v1alpha1"
	"github.com/nais/liberator/pkg/namegen"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation"

	"github.com/nais/naiserator/pkg/resourcecreator/pod"
	"github.com/nais/naiserator/pkg/resourcecreator/resource"
)

type Source interface {
	resource.Source
	GetMaskinporten() *nais_io_v1.Maskinporten
}

type Config interface {
	IsDigdiratorEnabled() bool
}

func secretName(name string) string {
	return namegen.PrefixedRandShortName("maskinporten", name, validation.DNS1035LabelMaxLength)
}

func client(objectMeta metav1.ObjectMeta, naisMaskinporten *nais_io_v1.Maskinporten) (*nais_io_v1.MaskinportenClient, error) {
	return &nais_io_v1.MaskinportenClient{
		TypeMeta: metav1.TypeMeta{
			Kind:       "MaskinportenClient",
			APIVersion: "nais.io/v1",
		},
		ObjectMeta: objectMeta,
		Spec: nais_io_v1.MaskinportenClientSpec{
			Scopes: nais_io_v1.MaskinportenScope{
				ConsumedScopes: naisMaskinporten.Scopes.ConsumedScopes,
				ExposedScopes:  naisMaskinporten.Scopes.ExposedScopes,
			},
			SecretName: secretName(objectMeta.Name),
		},
	}, nil
}

func Create(source Source, ast *resource.Ast, cfg Config) error {
	maskinporten := source.GetMaskinporten()

	if !cfg.IsDigdiratorEnabled() || maskinporten == nil || !maskinporten.Enabled {
		return nil
	}

	maskinportenClient, err := client(resource.CreateObjectMeta(source), maskinporten)
	if err != nil {
		return err
	}

	ast.AppendOperation(resource.OperationCreateOrUpdate, maskinportenClient)
	pod.WithAdditionalSecret(ast, maskinportenClient.Spec.SecretName, nais_io_v1alpha1.DefaultDigdiratorMaskinportenMountPath)
	pod.WithAdditionalEnvFromSecret(ast, maskinportenClient.Spec.SecretName)

	return nil
}
