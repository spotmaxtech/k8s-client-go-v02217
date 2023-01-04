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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "k8s.io/api/storage/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"github.com/spotmaxtech/k8s-client-go-v02217/tools/cache"
)

// CSIStorageCapacityLister helps list CSIStorageCapacities.
// All objects returned here must be treated as read-only.
type CSIStorageCapacityLister interface {
	// List lists all CSIStorageCapacities in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CSIStorageCapacity, err error)
	// CSIStorageCapacities returns an object that can list and get CSIStorageCapacities.
	CSIStorageCapacities(namespace string) CSIStorageCapacityNamespaceLister
	CSIStorageCapacityListerExpansion
}

// cSIStorageCapacityLister implements the CSIStorageCapacityLister interface.
type cSIStorageCapacityLister struct {
	indexer cache.Indexer
}

// NewCSIStorageCapacityLister returns a new CSIStorageCapacityLister.
func NewCSIStorageCapacityLister(indexer cache.Indexer) CSIStorageCapacityLister {
	return &cSIStorageCapacityLister{indexer: indexer}
}

// List lists all CSIStorageCapacities in the indexer.
func (s *cSIStorageCapacityLister) List(selector labels.Selector) (ret []*v1alpha1.CSIStorageCapacity, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CSIStorageCapacity))
	})
	return ret, err
}

// CSIStorageCapacities returns an object that can list and get CSIStorageCapacities.
func (s *cSIStorageCapacityLister) CSIStorageCapacities(namespace string) CSIStorageCapacityNamespaceLister {
	return cSIStorageCapacityNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CSIStorageCapacityNamespaceLister helps list and get CSIStorageCapacities.
// All objects returned here must be treated as read-only.
type CSIStorageCapacityNamespaceLister interface {
	// List lists all CSIStorageCapacities in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CSIStorageCapacity, err error)
	// Get retrieves the CSIStorageCapacity from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CSIStorageCapacity, error)
	CSIStorageCapacityNamespaceListerExpansion
}

// cSIStorageCapacityNamespaceLister implements the CSIStorageCapacityNamespaceLister
// interface.
type cSIStorageCapacityNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CSIStorageCapacities in the indexer for a given namespace.
func (s cSIStorageCapacityNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CSIStorageCapacity, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CSIStorageCapacity))
	})
	return ret, err
}

// Get retrieves the CSIStorageCapacity from the indexer for a given namespace and name.
func (s cSIStorageCapacityNamespaceLister) Get(name string) (*v1alpha1.CSIStorageCapacity, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("csistoragecapacity"), name)
	}
	return obj.(*v1alpha1.CSIStorageCapacity), nil
}
