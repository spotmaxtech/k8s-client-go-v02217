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

package v2beta2

import (
	resource "github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/api/resource"
)

// MetricValueStatusApplyConfiguration represents an declarative configuration of the MetricValueStatus type for use
// with apply.
type MetricValueStatusApplyConfiguration struct {
	Value              *resource.Quantity `json:"value,omitempty"`
	AverageValue       *resource.Quantity `json:"averageValue,omitempty"`
	AverageUtilization *int32             `json:"averageUtilization,omitempty"`
}

// MetricValueStatusApplyConfiguration constructs an declarative configuration of the MetricValueStatus type for use with
// apply.
func MetricValueStatus() *MetricValueStatusApplyConfiguration {
	return &MetricValueStatusApplyConfiguration{}
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *MetricValueStatusApplyConfiguration) WithValue(value resource.Quantity) *MetricValueStatusApplyConfiguration {
	b.Value = &value
	return b
}

// WithAverageValue sets the AverageValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AverageValue field is set to the value of the last call.
func (b *MetricValueStatusApplyConfiguration) WithAverageValue(value resource.Quantity) *MetricValueStatusApplyConfiguration {
	b.AverageValue = &value
	return b
}

// WithAverageUtilization sets the AverageUtilization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AverageUtilization field is set to the value of the last call.
func (b *MetricValueStatusApplyConfiguration) WithAverageUtilization(value int32) *MetricValueStatusApplyConfiguration {
	b.AverageUtilization = &value
	return b
}
