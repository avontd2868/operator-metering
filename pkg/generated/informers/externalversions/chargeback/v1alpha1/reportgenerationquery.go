// This file was automatically generated by informer-gen

package v1alpha1

import (
	time "time"

	chargeback_v1alpha1 "github.com/coreos-inc/kube-chargeback/pkg/apis/chargeback/v1alpha1"
	versioned "github.com/coreos-inc/kube-chargeback/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/coreos-inc/kube-chargeback/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/coreos-inc/kube-chargeback/pkg/generated/listers/chargeback/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ReportGenerationQueryInformer provides access to a shared informer and lister for
// ReportGenerationQueries.
type ReportGenerationQueryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ReportGenerationQueryLister
}

type reportGenerationQueryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewReportGenerationQueryInformer constructs a new informer for ReportGenerationQuery type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewReportGenerationQueryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredReportGenerationQueryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredReportGenerationQueryInformer constructs a new informer for ReportGenerationQuery type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredReportGenerationQueryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ChargebackV1alpha1().ReportGenerationQueries(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ChargebackV1alpha1().ReportGenerationQueries(namespace).Watch(options)
			},
		},
		&chargeback_v1alpha1.ReportGenerationQuery{},
		resyncPeriod,
		indexers,
	)
}

func (f *reportGenerationQueryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredReportGenerationQueryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *reportGenerationQueryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&chargeback_v1alpha1.ReportGenerationQuery{}, f.defaultInformer)
}

func (f *reportGenerationQueryInformer) Lister() v1alpha1.ReportGenerationQueryLister {
	return v1alpha1.NewReportGenerationQueryLister(f.Informer().GetIndexer())
}
