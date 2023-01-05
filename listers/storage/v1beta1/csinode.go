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

package v1beta1

import (
	v1beta1 "github.com/spotmaxtech/k8s-api-v02217/storage/v1beta1"
	"github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/api/errors"
	"github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/labels"
	"github.com/spotmaxtech/k8s-client-go-v02217/tools/cache"
)

// CSINodeLister helps list CSINodes.
// All objects returned here must be treated as read-only.
type CSINodeLister interface {
	// List lists all CSINodes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.CSINode, err error)
	// Get retrieves the CSINode from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.CSINode, error)
	CSINodeListerExpansion
}

// cSINodeLister implements the CSINodeLister interface.
type cSINodeLister struct {
	indexer cache.Indexer
}

// NewCSINodeLister returns a new CSINodeLister.
func NewCSINodeLister(indexer cache.Indexer) CSINodeLister {
	return &cSINodeLister{indexer: indexer}
}

// List lists all CSINodes in the indexer.
func (s *cSINodeLister) List(selector labels.Selector) (ret []*v1beta1.CSINode, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.CSINode))
	})
	return ret, err
}

// Get retrieves the CSINode from the index for a given name.
func (s *cSINodeLister) Get(name string) (*v1beta1.CSINode, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("csinode"), name)
	}
	return obj.(*v1beta1.CSINode), nil
}
