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

package v1

import (
	v1 "github.com/spotmaxtech/k8s-api-v02217/core/v1"
)

// ContainerPortApplyConfiguration represents an declarative configuration of the ContainerPort type for use
// with apply.
type ContainerPortApplyConfiguration struct {
	Name          *string      `json:"name,omitempty"`
	HostPort      *int32       `json:"hostPort,omitempty"`
	ContainerPort *int32       `json:"containerPort,omitempty"`
	Protocol      *v1.Protocol `json:"protocol,omitempty"`
	HostIP        *string      `json:"hostIP,omitempty"`
}

// ContainerPortApplyConfiguration constructs an declarative configuration of the ContainerPort type for use with
// apply.
func ContainerPort() *ContainerPortApplyConfiguration {
	return &ContainerPortApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ContainerPortApplyConfiguration) WithName(value string) *ContainerPortApplyConfiguration {
	b.Name = &value
	return b
}

// WithHostPort sets the HostPort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostPort field is set to the value of the last call.
func (b *ContainerPortApplyConfiguration) WithHostPort(value int32) *ContainerPortApplyConfiguration {
	b.HostPort = &value
	return b
}

// WithContainerPort sets the ContainerPort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ContainerPort field is set to the value of the last call.
func (b *ContainerPortApplyConfiguration) WithContainerPort(value int32) *ContainerPortApplyConfiguration {
	b.ContainerPort = &value
	return b
}

// WithProtocol sets the Protocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Protocol field is set to the value of the last call.
func (b *ContainerPortApplyConfiguration) WithProtocol(value v1.Protocol) *ContainerPortApplyConfiguration {
	b.Protocol = &value
	return b
}

// WithHostIP sets the HostIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostIP field is set to the value of the last call.
func (b *ContainerPortApplyConfiguration) WithHostIP(value string) *ContainerPortApplyConfiguration {
	b.HostIP = &value
	return b
}
