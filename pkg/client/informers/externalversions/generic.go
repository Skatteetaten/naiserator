// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/nais/naiserator/pkg/apis/iam.cnrm.cloud.google.com/v1alpha1"
	naisiov1alpha1 "github.com/nais/naiserator/pkg/apis/nais.io/v1alpha1"
	v1alpha3 "github.com/nais/naiserator/pkg/apis/networking.istio.io/v1alpha3"
	rbacistioiov1alpha1 "github.com/nais/naiserator/pkg/apis/rbac.istio.io/v1alpha1"
	sqlcnrmcloudgooglecomv1alpha3 "github.com/nais/naiserator/pkg/apis/sql.cnrm.cloud.google.com/v1alpha3"
	v1alpha2 "github.com/nais/naiserator/pkg/apis/storage.cnrm.cloud.google.com/v1alpha2"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=iam.cnrm.cloud.google.com, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("iampolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1alpha1().IAMPolicies().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("iamserviceaccounts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1alpha1().IAMServiceAccounts().Informer()}, nil

		// Group=naiserator.nais.io, Version=v1alpha1
	case naisiov1alpha1.SchemeGroupVersion.WithResource("applications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Naiserator().V1alpha1().Applications().Informer()}, nil

		// Group=networking.istio.io, Version=v1alpha3
	case v1alpha3.SchemeGroupVersion.WithResource("serviceentries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().ServiceEntries().Informer()}, nil
	case v1alpha3.SchemeGroupVersion.WithResource("virtualservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().VirtualServices().Informer()}, nil

		// Group=rbac.istio.io, Version=v1alpha1
	case rbacistioiov1alpha1.SchemeGroupVersion.WithResource("serviceroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1alpha1().ServiceRoles().Informer()}, nil
	case rbacistioiov1alpha1.SchemeGroupVersion.WithResource("servicerolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().V1alpha1().ServiceRoleBindings().Informer()}, nil

		// Group=sql.cnrm.cloud.google.com, Version=v1alpha3
	case sqlcnrmcloudgooglecomv1alpha3.SchemeGroupVersion.WithResource("sqldatabases"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Sql().V1alpha3().SqlDatabases().Informer()}, nil
	case sqlcnrmcloudgooglecomv1alpha3.SchemeGroupVersion.WithResource("sqlinstances"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Sql().V1alpha3().SqlInstances().Informer()}, nil
	case sqlcnrmcloudgooglecomv1alpha3.SchemeGroupVersion.WithResource("sqlusers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Sql().V1alpha3().SqlUsers().Informer()}, nil

		// Group=storage.cnrm.cloud.google.com, Version=v1alpha2
	case v1alpha2.SchemeGroupVersion.WithResource("googlestoragebuckets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1alpha2().GoogleStorageBuckets().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("googlestoragebucketaccesscontrols"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1alpha2().GoogleStorageBucketAccessControls().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
