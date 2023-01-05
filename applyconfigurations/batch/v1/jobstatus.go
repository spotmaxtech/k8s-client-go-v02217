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
	metav1 "github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/apis/meta/v1"
)

// JobStatusApplyConfiguration represents an declarative configuration of the JobStatus type for use
// with apply.
type JobStatusApplyConfiguration struct {
	Conditions              []JobConditionApplyConfiguration           `json:"conditions,omitempty"`
	StartTime               *metav1.Time                               `json:"startTime,omitempty"`
	CompletionTime          *metav1.Time                               `json:"completionTime,omitempty"`
	Active                  *int32                                     `json:"active,omitempty"`
	Succeeded               *int32                                     `json:"succeeded,omitempty"`
	Failed                  *int32                                     `json:"failed,omitempty"`
	CompletedIndexes        *string                                    `json:"completedIndexes,omitempty"`
	UncountedTerminatedPods *UncountedTerminatedPodsApplyConfiguration `json:"uncountedTerminatedPods,omitempty"`
}

// JobStatusApplyConfiguration constructs an declarative configuration of the JobStatus type for use with
// apply.
func JobStatus() *JobStatusApplyConfiguration {
	return &JobStatusApplyConfiguration{}
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *JobStatusApplyConfiguration) WithConditions(values ...*JobConditionApplyConfiguration) *JobStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithStartTime sets the StartTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StartTime field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithStartTime(value metav1.Time) *JobStatusApplyConfiguration {
	b.StartTime = &value
	return b
}

// WithCompletionTime sets the CompletionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CompletionTime field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithCompletionTime(value metav1.Time) *JobStatusApplyConfiguration {
	b.CompletionTime = &value
	return b
}

// WithActive sets the Active field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Active field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithActive(value int32) *JobStatusApplyConfiguration {
	b.Active = &value
	return b
}

// WithSucceeded sets the Succeeded field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Succeeded field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithSucceeded(value int32) *JobStatusApplyConfiguration {
	b.Succeeded = &value
	return b
}

// WithFailed sets the Failed field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Failed field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithFailed(value int32) *JobStatusApplyConfiguration {
	b.Failed = &value
	return b
}

// WithCompletedIndexes sets the CompletedIndexes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CompletedIndexes field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithCompletedIndexes(value string) *JobStatusApplyConfiguration {
	b.CompletedIndexes = &value
	return b
}

// WithUncountedTerminatedPods sets the UncountedTerminatedPods field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UncountedTerminatedPods field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithUncountedTerminatedPods(value *UncountedTerminatedPodsApplyConfiguration) *JobStatusApplyConfiguration {
	b.UncountedTerminatedPods = value
	return b
}
