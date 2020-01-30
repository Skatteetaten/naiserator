package resourcecreator

import (
	"fmt"

	nais "github.com/nais/naiserator/pkg/apis/nais.io/v1alpha1"
	istio "istio.io/api/security/v1beta1"
	"istio.io/api/type/v1beta1"
	istio_security_client "istio.io/client-go/pkg/apis/security/v1beta1"
	k8s_meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func AuthorizationPolicy(app *nais.Application) *istio_security_client.AuthorizationPolicy {
	objectMeta := app.CreateObjectMeta()
	var rules []*istio.Rule

	rules = append(rules, createRuleForPrometheus())

	if len(app.Spec.Ingresses) > 0 {
		rules = append(rules, createRuleForIngressGateway())
	}

	if len(app.Spec.AccessPolicy.Inbound.Rules) > 0 {
		rules = append(rules, createAccessPolicyPrincipals(app))
	}

	return &istio_security_client.AuthorizationPolicy{
		TypeMeta: k8s_meta.TypeMeta{
			Kind:       "AuthorizationPolicy",
			APIVersion: IstioAuthorizationPolicyVersion,
		},
		ObjectMeta: objectMeta,
		Spec: istio.AuthorizationPolicy{
			Selector: &v1beta1.WorkloadSelector{
				MatchLabels: map[string]string{"app": app.Name},
			},
			Rules: rules,
		},
	}
}

func createRuleForIngressGateway() *istio.Rule {
	return &istio.Rule{
		From: []*istio.Rule_From{
			&istio.Rule_From{
				Source: &istio.Source{
					Principals: []string{fmt.Sprintf("cluster.local/ns/%s/sa/%s", IstioNamespace, IstioIngressGatewayServiceAccount)},
				},
			},
		},
		To: []*istio.Rule_To{
			{
				Operation: &istio.Operation{
					Methods: []string{"*"},
					Paths:   []string{"*"},
				},
			},
		},
	}
}

func createRuleForPrometheus() *istio.Rule {
	return &istio.Rule{
		From: []*istio.Rule_From{
			&istio.Rule_From{
				Source: &istio.Source{
					Namespaces: []string{IstioNamespace},
				},
			},
		},
		To: []*istio.Rule_To{
			{
				Operation: &istio.Operation{
					Ports: []string{IstioPrometheusPort},
				},
			},
		},
	}
}

func createAccessPolicyPrincipals(app *nais.Application) *istio.Rule {
	var principals []string

	for _, rule := range app.Spec.AccessPolicy.Inbound.Rules {
		var namespace string
		if rule.Namespace == "" {
			namespace = app.Namespace
		} else {
			namespace = rule.Namespace
		}
		tmp := fmt.Sprintf("cluster.local/ns/%s/sa/%s", namespace, rule.Application)
		principals = append(principals, tmp)
	}

	return &istio.Rule{
		From: []*istio.Rule_From{
			&istio.Rule_From{
				Source: &istio.Source{
					Principals: principals,
				},
			},
		},
		To: []*istio.Rule_To{
			{
				Operation: &istio.Operation{
					Methods: []string{"*"},
					Paths:   []string{"*"},
				},
			},
		},
	}
}
