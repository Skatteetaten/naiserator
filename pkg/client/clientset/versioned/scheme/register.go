// Code generated by client-gen. DO NOT EDIT.

package scheme

import (
	iamv1beta1 "github.com/nais/naiserator/pkg/apis/iam.cnrm.cloud.google.com/v1beta1"
	naisv1 "github.com/nais/naiserator/pkg/apis/nais.io/v1"
	naisv1alpha1 "github.com/nais/naiserator/pkg/apis/nais.io/v1alpha1"
	networkingv1alpha3 "github.com/nais/naiserator/pkg/apis/networking.istio.io/v1alpha3"
	sqlv1beta1 "github.com/nais/naiserator/pkg/apis/sql.cnrm.cloud.google.com/v1beta1"
	storagev1beta1 "github.com/nais/naiserator/pkg/apis/storage.cnrm.cloud.google.com/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)
var localSchemeBuilder = runtime.SchemeBuilder{
	iamv1beta1.AddToScheme,
	naisv1.AddToScheme,
	naisv1alpha1.AddToScheme,
	networkingv1alpha3.AddToScheme,
	sqlv1beta1.AddToScheme,
	storagev1beta1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   _ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(Scheme))
}
