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
	v1 "k8s.io/api/core/v1"
	metav1 "github.com/spotmaxtech/k8s-client-go-v02217/applyconfigurations/meta/v1"
)

// PersistentVolumeClaimSpecApplyConfiguration represents an declarative configuration of the PersistentVolumeClaimSpec type for use
// with apply.
type PersistentVolumeClaimSpecApplyConfiguration struct {
	AccessModes      []v1.PersistentVolumeAccessMode              `json:"accessModes,omitempty"`
	Selector         *metav1.LabelSelectorApplyConfiguration      `json:"selector,omitempty"`
	Resources        *ResourceRequirementsApplyConfiguration      `json:"resources,omitempty"`
	VolumeName       *string                                      `json:"volumeName,omitempty"`
	StorageClassName *string                                      `json:"storageClassName,omitempty"`
	VolumeMode       *v1.PersistentVolumeMode                     `json:"volumeMode,omitempty"`
	DataSource       *TypedLocalObjectReferenceApplyConfiguration `json:"dataSource,omitempty"`
	DataSourceRef    *TypedLocalObjectReferenceApplyConfiguration `json:"dataSourceRef,omitempty"`
}

// PersistentVolumeClaimSpecApplyConfiguration constructs an declarative configuration of the PersistentVolumeClaimSpec type for use with
// apply.
func PersistentVolumeClaimSpec() *PersistentVolumeClaimSpecApplyConfiguration {
	return &PersistentVolumeClaimSpecApplyConfiguration{}
}

// WithAccessModes adds the given value to the AccessModes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AccessModes field.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithAccessModes(values ...v1.PersistentVolumeAccessMode) *PersistentVolumeClaimSpecApplyConfiguration {
	for i := range values {
		b.AccessModes = append(b.AccessModes, values[i])
	}
	return b
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithSelector(value *metav1.LabelSelectorApplyConfiguration) *PersistentVolumeClaimSpecApplyConfiguration {
	b.Selector = value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithResources(value *ResourceRequirementsApplyConfiguration) *PersistentVolumeClaimSpecApplyConfiguration {
	b.Resources = value
	return b
}

// WithVolumeName sets the VolumeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeName field is set to the value of the last call.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithVolumeName(value string) *PersistentVolumeClaimSpecApplyConfiguration {
	b.VolumeName = &value
	return b
}

// WithStorageClassName sets the StorageClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageClassName field is set to the value of the last call.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithStorageClassName(value string) *PersistentVolumeClaimSpecApplyConfiguration {
	b.StorageClassName = &value
	return b
}

// WithVolumeMode sets the VolumeMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeMode field is set to the value of the last call.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithVolumeMode(value v1.PersistentVolumeMode) *PersistentVolumeClaimSpecApplyConfiguration {
	b.VolumeMode = &value
	return b
}

// WithDataSource sets the DataSource field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DataSource field is set to the value of the last call.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithDataSource(value *TypedLocalObjectReferenceApplyConfiguration) *PersistentVolumeClaimSpecApplyConfiguration {
	b.DataSource = value
	return b
}

// WithDataSourceRef sets the DataSourceRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DataSourceRef field is set to the value of the last call.
func (b *PersistentVolumeClaimSpecApplyConfiguration) WithDataSourceRef(value *TypedLocalObjectReferenceApplyConfiguration) *PersistentVolumeClaimSpecApplyConfiguration {
	b.DataSourceRef = value
	return b
}
