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

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/spotmaxtech/k8s-api-v02217/apps/v1"
	metav1 "github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/apis/meta/v1"
	types "github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/types"
	watch "github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/watch"
	appsv1 "github.com/spotmaxtech/k8s-client-go-v02217/applyconfigurations/apps/v1"
	scheme "github.com/spotmaxtech/k8s-client-go-v02217/kubernetes/scheme"
	rest "github.com/spotmaxtech/k8s-client-go-v02217/rest"
)

// DaemonSetsGetter has a method to return a DaemonSetInterface.
// A group's client should implement this interface.
type DaemonSetsGetter interface {
	DaemonSets(namespace string) DaemonSetInterface
}

// DaemonSetInterface has methods to work with DaemonSet resources.
type DaemonSetInterface interface {
	Create(ctx context.Context, daemonSet *v1.DaemonSet, opts metav1.CreateOptions) (*v1.DaemonSet, error)
	Update(ctx context.Context, daemonSet *v1.DaemonSet, opts metav1.UpdateOptions) (*v1.DaemonSet, error)
	UpdateStatus(ctx context.Context, daemonSet *v1.DaemonSet, opts metav1.UpdateOptions) (*v1.DaemonSet, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DaemonSet, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DaemonSetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DaemonSet, err error)
	Apply(ctx context.Context, daemonSet *appsv1.DaemonSetApplyConfiguration, opts metav1.ApplyOptions) (result *v1.DaemonSet, err error)
	ApplyStatus(ctx context.Context, daemonSet *appsv1.DaemonSetApplyConfiguration, opts metav1.ApplyOptions) (result *v1.DaemonSet, err error)
	DaemonSetExpansion
}

// daemonSets implements DaemonSetInterface
type daemonSets struct {
	client rest.Interface
	ns     string
}

// newDaemonSets returns a DaemonSets
func newDaemonSets(c *AppsV1Client, namespace string) *daemonSets {
	return &daemonSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the daemonSet, and returns the corresponding daemonSet object, and an error if there is any.
func (c *daemonSets) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DaemonSet, err error) {
	result = &v1.DaemonSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("daemonsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DaemonSets that match those selectors.
func (c *daemonSets) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DaemonSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DaemonSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("daemonsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested daemonSets.
func (c *daemonSets) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("daemonsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a daemonSet and creates it.  Returns the server's representation of the daemonSet, and an error, if there is any.
func (c *daemonSets) Create(ctx context.Context, daemonSet *v1.DaemonSet, opts metav1.CreateOptions) (result *v1.DaemonSet, err error) {
	result = &v1.DaemonSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("daemonsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(daemonSet).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a daemonSet and updates it. Returns the server's representation of the daemonSet, and an error, if there is any.
func (c *daemonSets) Update(ctx context.Context, daemonSet *v1.DaemonSet, opts metav1.UpdateOptions) (result *v1.DaemonSet, err error) {
	result = &v1.DaemonSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("daemonsets").
		Name(daemonSet.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(daemonSet).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *daemonSets) UpdateStatus(ctx context.Context, daemonSet *v1.DaemonSet, opts metav1.UpdateOptions) (result *v1.DaemonSet, err error) {
	result = &v1.DaemonSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("daemonsets").
		Name(daemonSet.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(daemonSet).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the daemonSet and deletes it. Returns an error if one occurs.
func (c *daemonSets) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("daemonsets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *daemonSets) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("daemonsets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched daemonSet.
func (c *daemonSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DaemonSet, err error) {
	result = &v1.DaemonSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("daemonsets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied daemonSet.
func (c *daemonSets) Apply(ctx context.Context, daemonSet *appsv1.DaemonSetApplyConfiguration, opts metav1.ApplyOptions) (result *v1.DaemonSet, err error) {
	if daemonSet == nil {
		return nil, fmt.Errorf("daemonSet provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(daemonSet)
	if err != nil {
		return nil, err
	}
	name := daemonSet.Name
	if name == nil {
		return nil, fmt.Errorf("daemonSet.Name must be provided to Apply")
	}
	result = &v1.DaemonSet{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("daemonsets").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *daemonSets) ApplyStatus(ctx context.Context, daemonSet *appsv1.DaemonSetApplyConfiguration, opts metav1.ApplyOptions) (result *v1.DaemonSet, err error) {
	if daemonSet == nil {
		return nil, fmt.Errorf("daemonSet provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(daemonSet)
	if err != nil {
		return nil, err
	}

	name := daemonSet.Name
	if name == nil {
		return nil, fmt.Errorf("daemonSet.Name must be provided to Apply")
	}

	result = &v1.DaemonSet{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("daemonsets").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
