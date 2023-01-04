/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	coordinationv1beta1 "k8s.io/api/coordination/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "github.com/spotmaxtech/k8s-client-go-v02217/informers/internalinterfaces"
	kubernetes "github.com/spotmaxtech/k8s-client-go-v02217/kubernetes"
	v1beta1 "github.com/spotmaxtech/k8s-client-go-v02217/listers/coordination/v1beta1"
	cache "github.com/spotmaxtech/k8s-client-go-v02217/tools/cache"
)

// LeaseInformer provides access to a shared informer and lister for
// Leases.
type LeaseInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.LeaseLister
}

type leaseInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLeaseInformer constructs a new informer for Lease type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLeaseInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLeaseInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLeaseInformer constructs a new informer for Lease type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLeaseInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoordinationV1beta1().Leases(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoordinationV1beta1().Leases(namespace).Watch(context.TODO(), options)
			},
		},
		&coordinationv1beta1.Lease{},
		resyncPeriod,
		indexers,
	)
}

func (f *leaseInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLeaseInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *leaseInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&coordinationv1beta1.Lease{}, f.defaultInformer)
}

func (f *leaseInformer) Lister() v1beta1.LeaseLister {
	return v1beta1.NewLeaseLister(f.Informer().GetIndexer())
}
