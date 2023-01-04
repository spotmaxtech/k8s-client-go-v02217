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

package v1beta2

import (
	corev1 "github.com/spotmaxtech/k8s-client-go-v02217/applyconfigurations/core/v1"
	v1 "github.com/spotmaxtech/k8s-client-go-v02217/applyconfigurations/meta/v1"
)

// ReplicaSetSpecApplyConfiguration represents an declarative configuration of the ReplicaSetSpec type for use
// with apply.
type ReplicaSetSpecApplyConfiguration struct {
	Replicas        *int32                                    `json:"replicas,omitempty"`
	MinReadySeconds *int32                                    `json:"minReadySeconds,omitempty"`
	Selector        *v1.LabelSelectorApplyConfiguration       `json:"selector,omitempty"`
	Template        *corev1.PodTemplateSpecApplyConfiguration `json:"template,omitempty"`
}

// ReplicaSetSpecApplyConfiguration constructs an declarative configuration of the ReplicaSetSpec type for use with
// apply.
func ReplicaSetSpec() *ReplicaSetSpecApplyConfiguration {
	return &ReplicaSetSpecApplyConfiguration{}
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *ReplicaSetSpecApplyConfiguration) WithReplicas(value int32) *ReplicaSetSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithMinReadySeconds sets the MinReadySeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinReadySeconds field is set to the value of the last call.
func (b *ReplicaSetSpecApplyConfiguration) WithMinReadySeconds(value int32) *ReplicaSetSpecApplyConfiguration {
	b.MinReadySeconds = &value
	return b
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *ReplicaSetSpecApplyConfiguration) WithSelector(value *v1.LabelSelectorApplyConfiguration) *ReplicaSetSpecApplyConfiguration {
	b.Selector = value
	return b
}

// WithTemplate sets the Template field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Template field is set to the value of the last call.
func (b *ReplicaSetSpecApplyConfiguration) WithTemplate(value *corev1.PodTemplateSpecApplyConfiguration) *ReplicaSetSpecApplyConfiguration {
	b.Template = value
	return b
}
