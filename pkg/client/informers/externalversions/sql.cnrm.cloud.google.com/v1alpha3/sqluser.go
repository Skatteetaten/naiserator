// Code generated by informer-gen. DO NOT EDIT.

package v1alpha3

import (
	time "time"

	sqlcnrmcloudgooglecomv1alpha3 "github.com/nais/naiserator/pkg/apis/sql.cnrm.cloud.google.com/v1alpha3"
	versioned "github.com/nais/naiserator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nais/naiserator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha3 "github.com/nais/naiserator/pkg/client/listers/sql.cnrm.cloud.google.com/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SqlUserInformer provides access to a shared informer and lister for
// SqlUsers.
type SqlUserInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha3.SqlUserLister
}

type sqlUserInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSqlUserInformer constructs a new informer for SqlUser type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSqlUserInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSqlUserInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSqlUserInformer constructs a new informer for SqlUser type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSqlUserInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SqlV1alpha3().SqlUsers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SqlV1alpha3().SqlUsers(namespace).Watch(options)
			},
		},
		&sqlcnrmcloudgooglecomv1alpha3.SqlUser{},
		resyncPeriod,
		indexers,
	)
}

func (f *sqlUserInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSqlUserInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sqlUserInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sqlcnrmcloudgooglecomv1alpha3.SqlUser{}, f.defaultInformer)
}

func (f *sqlUserInformer) Lister() v1alpha3.SqlUserLister {
	return v1alpha3.NewSqlUserLister(f.Informer().GetIndexer())
}
