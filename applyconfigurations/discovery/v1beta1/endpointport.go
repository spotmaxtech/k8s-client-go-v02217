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
	v1 "github.com/spotmaxtech/k8s-api-v02217/core/v1"
)

// EndpointPortApplyConfiguration represents an declarative configuration of the EndpointPort type for use
// with apply.
type EndpointPortApplyConfiguration struct {
	Name        *string      `json:"name,omitempty"`
	Protocol    *v1.Protocol `json:"protocol,omitempty"`
	Port        *int32       `json:"port,omitempty"`
	AppProtocol *string      `json:"appProtocol,omitempty"`
}

// EndpointPortApplyConfiguration constructs an declarative configuration of the EndpointPort type for use with
// apply.
func EndpointPort() *EndpointPortApplyConfiguration {
	return &EndpointPortApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *EndpointPortApplyConfiguration) WithName(value string) *EndpointPortApplyConfiguration {
	b.Name = &value
	return b
}

// WithProtocol sets the Protocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Protocol field is set to the value of the last call.
func (b *EndpointPortApplyConfiguration) WithProtocol(value v1.Protocol) *EndpointPortApplyConfiguration {
	b.Protocol = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *EndpointPortApplyConfiguration) WithPort(value int32) *EndpointPortApplyConfiguration {
	b.Port = &value
	return b
}

// WithAppProtocol sets the AppProtocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AppProtocol field is set to the value of the last call.
func (b *EndpointPortApplyConfiguration) WithAppProtocol(value string) *EndpointPortApplyConfiguration {
	b.AppProtocol = &value
	return b
}
