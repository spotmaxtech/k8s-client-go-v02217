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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "k8s.io/api/flowcontrol/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	flowcontrolv1alpha1 "github.com/spotmaxtech/k8s-client-go-v02217/applyconfigurations/flowcontrol/v1alpha1"
	scheme "github.com/spotmaxtech/k8s-client-go-v02217/kubernetes/scheme"
	rest "github.com/spotmaxtech/k8s-client-go-v02217/rest"
)

// PriorityLevelConfigurationsGetter has a method to return a PriorityLevelConfigurationInterface.
// A group's client should implement this interface.
type PriorityLevelConfigurationsGetter interface {
	PriorityLevelConfigurations() PriorityLevelConfigurationInterface
}

// PriorityLevelConfigurationInterface has methods to work with PriorityLevelConfiguration resources.
type PriorityLevelConfigurationInterface interface {
	Create(ctx context.Context, priorityLevelConfiguration *v1alpha1.PriorityLevelConfiguration, opts v1.CreateOptions) (*v1alpha1.PriorityLevelConfiguration, error)
	Update(ctx context.Context, priorityLevelConfiguration *v1alpha1.PriorityLevelConfiguration, opts v1.UpdateOptions) (*v1alpha1.PriorityLevelConfiguration, error)
	UpdateStatus(ctx context.Context, priorityLevelConfiguration *v1alpha1.PriorityLevelConfiguration, opts v1.UpdateOptions) (*v1alpha1.PriorityLevelConfiguration, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PriorityLevelConfiguration, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PriorityLevelConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PriorityLevelConfiguration, err error)
	Apply(ctx context.Context, priorityLevelConfiguration *flowcontrolv1alpha1.PriorityLevelConfigurationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PriorityLevelConfiguration, err error)
	ApplyStatus(ctx context.Context, priorityLevelConfiguration *flowcontrolv1alpha1.PriorityLevelConfigurationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PriorityLevelConfiguration, err error)
	PriorityLevelConfigurationExpansion
}

// priorityLevelConfigurations implements PriorityLevelConfigurationInterface
type priorityLevelConfigurations struct {
	client rest.Interface
}

// newPriorityLevelConfigurations returns a PriorityLevelConfigurations
func newPriorityLevelConfigurations(c *FlowcontrolV1alpha1Client) *priorityLevelConfigurations {
	return &priorityLevelConfigurations{
		client: c.RESTClient(),
	}
}

// Get takes name of the priorityLevelConfiguration, and returns the corresponding priorityLevelConfiguration object, and an error if there is any.
func (c *priorityLevelConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PriorityLevelConfiguration, err error) {
	result = &v1alpha1.PriorityLevelConfiguration{}
	err = c.client.Get().
		Resource("prioritylevelconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PriorityLevelConfigurations that match those selectors.
func (c *priorityLevelConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PriorityLevelConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PriorityLevelConfigurationList{}
	err = c.client.Get().
		Resource("prioritylevelconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested priorityLevelConfigurations.
func (c *priorityLevelConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("prioritylevelconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a priorityLevelConfiguration and creates it.  Returns the server's representation of the priorityLevelConfiguration, and an error, if there is any.
func (c *priorityLevelConfigurations) Create(ctx context.Context, priorityLevelConfiguration *v1alpha1.PriorityLevelConfiguration, opts v1.CreateOptions) (result *v1alpha1.PriorityLevelConfiguration, err error) {
	result = &v1alpha1.PriorityLevelConfiguration{}
	err = c.client.Post().
		Resource("prioritylevelconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(priorityLevelConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a priorityLevelConfiguration and updates it. Returns the server's representation of the priorityLevelConfiguration, and an error, if there is any.
func (c *priorityLevelConfigurations) Update(ctx context.Context, priorityLevelConfiguration *v1alpha1.PriorityLevelConfiguration, opts v1.UpdateOptions) (result *v1alpha1.PriorityLevelConfiguration, err error) {
	result = &v1alpha1.PriorityLevelConfiguration{}
	err = c.client.Put().
		Resource("prioritylevelconfigurations").
		Name(priorityLevelConfiguration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(priorityLevelConfiguration).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *priorityLevelConfigurations) UpdateStatus(ctx context.Context, priorityLevelConfiguration *v1alpha1.PriorityLevelConfiguration, opts v1.UpdateOptions) (result *v1alpha1.PriorityLevelConfiguration, err error) {
	result = &v1alpha1.PriorityLevelConfiguration{}
	err = c.client.Put().
		Resource("prioritylevelconfigurations").
		Name(priorityLevelConfiguration.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(priorityLevelConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the priorityLevelConfiguration and deletes it. Returns an error if one occurs.
func (c *priorityLevelConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("prioritylevelconfigurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *priorityLevelConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("prioritylevelconfigurations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched priorityLevelConfiguration.
func (c *priorityLevelConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PriorityLevelConfiguration, err error) {
	result = &v1alpha1.PriorityLevelConfiguration{}
	err = c.client.Patch(pt).
		Resource("prioritylevelconfigurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied priorityLevelConfiguration.
func (c *priorityLevelConfigurations) Apply(ctx context.Context, priorityLevelConfiguration *flowcontrolv1alpha1.PriorityLevelConfigurationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PriorityLevelConfiguration, err error) {
	if priorityLevelConfiguration == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(priorityLevelConfiguration)
	if err != nil {
		return nil, err
	}
	name := priorityLevelConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration.Name must be provided to Apply")
	}
	result = &v1alpha1.PriorityLevelConfiguration{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("prioritylevelconfigurations").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *priorityLevelConfigurations) ApplyStatus(ctx context.Context, priorityLevelConfiguration *flowcontrolv1alpha1.PriorityLevelConfigurationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PriorityLevelConfiguration, err error) {
	if priorityLevelConfiguration == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(priorityLevelConfiguration)
	if err != nil {
		return nil, err
	}

	name := priorityLevelConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration.Name must be provided to Apply")
	}

	result = &v1alpha1.PriorityLevelConfiguration{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("prioritylevelconfigurations").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
