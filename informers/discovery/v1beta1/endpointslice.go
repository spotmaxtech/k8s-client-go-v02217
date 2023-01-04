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

	discoveryv1beta1 "k8s.io/api/discovery/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "github.com/spotmaxtech/k8s-client-go-v02217/informers/internalinterfaces"
	kubernetes "github.com/spotmaxtech/k8s-client-go-v02217/kubernetes"
	v1beta1 "github.com/spotmaxtech/k8s-client-go-v02217/listers/discovery/v1beta1"
	cache "github.com/spotmaxtech/k8s-client-go-v02217/tools/cache"
)

// EndpointSliceInformer provides access to a shared informer and lister for
// EndpointSlices.
type EndpointSliceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.EndpointSliceLister
}

type endpointSliceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEndpointSliceInformer constructs a new informer for EndpointSlice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEndpointSliceInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEndpointSliceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEndpointSliceInformer constructs a new informer for EndpointSlice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEndpointSliceInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DiscoveryV1beta1().EndpointSlices(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DiscoveryV1beta1().EndpointSlices(namespace).Watch(context.TODO(), options)
			},
		},
		&discoveryv1beta1.EndpointSlice{},
		resyncPeriod,
		indexers,
	)
}

func (f *endpointSliceInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEndpointSliceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *endpointSliceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&discoveryv1beta1.EndpointSlice{}, f.defaultInformer)
}

func (f *endpointSliceInformer) Lister() v1beta1.EndpointSliceLister {
	return v1beta1.NewEndpointSliceLister(f.Informer().GetIndexer())
}
