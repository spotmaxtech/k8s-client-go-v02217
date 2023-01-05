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

package v2beta1

import (
	v2beta1 "github.com/spotmaxtech/k8s-api-v02217/autoscaling/v2beta1"
)

// MetricStatusApplyConfiguration represents an declarative configuration of the MetricStatus type for use
// with apply.
type MetricStatusApplyConfiguration struct {
	Type              *v2beta1.MetricSourceType                        `json:"type,omitempty"`
	Object            *ObjectMetricStatusApplyConfiguration            `json:"object,omitempty"`
	Pods              *PodsMetricStatusApplyConfiguration              `json:"pods,omitempty"`
	Resource          *ResourceMetricStatusApplyConfiguration          `json:"resource,omitempty"`
	ContainerResource *ContainerResourceMetricStatusApplyConfiguration `json:"containerResource,omitempty"`
	External          *ExternalMetricStatusApplyConfiguration          `json:"external,omitempty"`
}

// MetricStatusApplyConfiguration constructs an declarative configuration of the MetricStatus type for use with
// apply.
func MetricStatus() *MetricStatusApplyConfiguration {
	return &MetricStatusApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *MetricStatusApplyConfiguration) WithType(value v2beta1.MetricSourceType) *MetricStatusApplyConfiguration {
	b.Type = &value
	return b
}

// WithObject sets the Object field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Object field is set to the value of the last call.
func (b *MetricStatusApplyConfiguration) WithObject(value *ObjectMetricStatusApplyConfiguration) *MetricStatusApplyConfiguration {
	b.Object = value
	return b
}

// WithPods sets the Pods field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Pods field is set to the value of the last call.
func (b *MetricStatusApplyConfiguration) WithPods(value *PodsMetricStatusApplyConfiguration) *MetricStatusApplyConfiguration {
	b.Pods = value
	return b
}

// WithResource sets the Resource field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resource field is set to the value of the last call.
func (b *MetricStatusApplyConfiguration) WithResource(value *ResourceMetricStatusApplyConfiguration) *MetricStatusApplyConfiguration {
	b.Resource = value
	return b
}

// WithContainerResource sets the ContainerResource field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ContainerResource field is set to the value of the last call.
func (b *MetricStatusApplyConfiguration) WithContainerResource(value *ContainerResourceMetricStatusApplyConfiguration) *MetricStatusApplyConfiguration {
	b.ContainerResource = value
	return b
}

// WithExternal sets the External field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the External field is set to the value of the last call.
func (b *MetricStatusApplyConfiguration) WithExternal(value *ExternalMetricStatusApplyConfiguration) *MetricStatusApplyConfiguration {
	b.External = value
	return b
}
