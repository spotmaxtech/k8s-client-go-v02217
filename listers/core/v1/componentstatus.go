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

package v1

import (
	v1 "github.com/spotmaxtech/k8s-api-v02217/core/v1"
	"github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/api/errors"
	"github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/labels"
	"github.com/spotmaxtech/k8s-client-go-v02217/tools/cache"
)

// ComponentStatusLister helps list ComponentStatuses.
// All objects returned here must be treated as read-only.
type ComponentStatusLister interface {
	// List lists all ComponentStatuses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ComponentStatus, err error)
	// Get retrieves the ComponentStatus from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ComponentStatus, error)
	ComponentStatusListerExpansion
}

// componentStatusLister implements the ComponentStatusLister interface.
type componentStatusLister struct {
	indexer cache.Indexer
}

// NewComponentStatusLister returns a new ComponentStatusLister.
func NewComponentStatusLister(indexer cache.Indexer) ComponentStatusLister {
	return &componentStatusLister{indexer: indexer}
}

// List lists all ComponentStatuses in the indexer.
func (s *componentStatusLister) List(selector labels.Selector) (ret []*v1.ComponentStatus, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ComponentStatus))
	})
	return ret, err
}

// Get retrieves the ComponentStatus from the index for a given name.
func (s *componentStatusLister) Get(name string) (*v1.ComponentStatus, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("componentstatus"), name)
	}
	return obj.(*v1.ComponentStatus), nil
}
