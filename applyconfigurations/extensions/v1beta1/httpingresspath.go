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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/spotmaxtech/k8s-api-v02217/extensions/v1beta1"
)

// HTTPIngressPathApplyConfiguration represents an declarative configuration of the HTTPIngressPath type for use
// with apply.
type HTTPIngressPathApplyConfiguration struct {
	Path     *string                           `json:"path,omitempty"`
	PathType *v1beta1.PathType                 `json:"pathType,omitempty"`
	Backend  *IngressBackendApplyConfiguration `json:"backend,omitempty"`
}

// HTTPIngressPathApplyConfiguration constructs an declarative configuration of the HTTPIngressPath type for use with
// apply.
func HTTPIngressPath() *HTTPIngressPathApplyConfiguration {
	return &HTTPIngressPathApplyConfiguration{}
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *HTTPIngressPathApplyConfiguration) WithPath(value string) *HTTPIngressPathApplyConfiguration {
	b.Path = &value
	return b
}

// WithPathType sets the PathType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PathType field is set to the value of the last call.
func (b *HTTPIngressPathApplyConfiguration) WithPathType(value v1beta1.PathType) *HTTPIngressPathApplyConfiguration {
	b.PathType = &value
	return b
}

// WithBackend sets the Backend field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Backend field is set to the value of the last call.
func (b *HTTPIngressPathApplyConfiguration) WithBackend(value *IngressBackendApplyConfiguration) *HTTPIngressPathApplyConfiguration {
	b.Backend = value
	return b
}
