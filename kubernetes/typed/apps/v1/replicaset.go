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
	autoscalingv1 "github.com/spotmaxtech/k8s-api-v02217/autoscaling/v1"
	metav1 "github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/apis/meta/v1"
	types "github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/types"
	watch "github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/watch"
	appsv1 "github.com/spotmaxtech/k8s-client-go-v02217/applyconfigurations/apps/v1"
	applyconfigurationsautoscalingv1 "github.com/spotmaxtech/k8s-client-go-v02217/applyconfigurations/autoscaling/v1"
	scheme "github.com/spotmaxtech/k8s-client-go-v02217/kubernetes/scheme"
	rest "github.com/spotmaxtech/k8s-client-go-v02217/rest"
)

// ReplicaSetsGetter has a method to return a ReplicaSetInterface.
// A group's client should implement this interface.
type ReplicaSetsGetter interface {
	ReplicaSets(namespace string) ReplicaSetInterface
}

// ReplicaSetInterface has methods to work with ReplicaSet resources.
type ReplicaSetInterface interface {
	Create(ctx context.Context, replicaSet *v1.ReplicaSet, opts metav1.CreateOptions) (*v1.ReplicaSet, error)
	Update(ctx context.Context, replicaSet *v1.ReplicaSet, opts metav1.UpdateOptions) (*v1.ReplicaSet, error)
	UpdateStatus(ctx context.Context, replicaSet *v1.ReplicaSet, opts metav1.UpdateOptions) (*v1.ReplicaSet, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ReplicaSet, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ReplicaSetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ReplicaSet, err error)
	Apply(ctx context.Context, replicaSet *appsv1.ReplicaSetApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ReplicaSet, err error)
	ApplyStatus(ctx context.Context, replicaSet *appsv1.ReplicaSetApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ReplicaSet, err error)
	GetScale(ctx context.Context, replicaSetName string, options metav1.GetOptions) (*autoscalingv1.Scale, error)
	UpdateScale(ctx context.Context, replicaSetName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (*autoscalingv1.Scale, error)
	ApplyScale(ctx context.Context, replicaSetName string, scale *applyconfigurationsautoscalingv1.ScaleApplyConfiguration, opts metav1.ApplyOptions) (*autoscalingv1.Scale, error)

	ReplicaSetExpansion
}

// replicaSets implements ReplicaSetInterface
type replicaSets struct {
	client rest.Interface
	ns     string
}

// newReplicaSets returns a ReplicaSets
func newReplicaSets(c *AppsV1Client, namespace string) *replicaSets {
	return &replicaSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the replicaSet, and returns the corresponding replicaSet object, and an error if there is any.
func (c *replicaSets) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ReplicaSet, err error) {
	result = &v1.ReplicaSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("replicasets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ReplicaSets that match those selectors.
func (c *replicaSets) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ReplicaSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ReplicaSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("replicasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested replicaSets.
func (c *replicaSets) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("replicasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a replicaSet and creates it.  Returns the server's representation of the replicaSet, and an error, if there is any.
func (c *replicaSets) Create(ctx context.Context, replicaSet *v1.ReplicaSet, opts metav1.CreateOptions) (result *v1.ReplicaSet, err error) {
	result = &v1.ReplicaSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("replicasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(replicaSet).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a replicaSet and updates it. Returns the server's representation of the replicaSet, and an error, if there is any.
func (c *replicaSets) Update(ctx context.Context, replicaSet *v1.ReplicaSet, opts metav1.UpdateOptions) (result *v1.ReplicaSet, err error) {
	result = &v1.ReplicaSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("replicasets").
		Name(replicaSet.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(replicaSet).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *replicaSets) UpdateStatus(ctx context.Context, replicaSet *v1.ReplicaSet, opts metav1.UpdateOptions) (result *v1.ReplicaSet, err error) {
	result = &v1.ReplicaSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("replicasets").
		Name(replicaSet.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(replicaSet).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the replicaSet and deletes it. Returns an error if one occurs.
func (c *replicaSets) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("replicasets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *replicaSets) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("replicasets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched replicaSet.
func (c *replicaSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ReplicaSet, err error) {
	result = &v1.ReplicaSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("replicasets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied replicaSet.
func (c *replicaSets) Apply(ctx context.Context, replicaSet *appsv1.ReplicaSetApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ReplicaSet, err error) {
	if replicaSet == nil {
		return nil, fmt.Errorf("replicaSet provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(replicaSet)
	if err != nil {
		return nil, err
	}
	name := replicaSet.Name
	if name == nil {
		return nil, fmt.Errorf("replicaSet.Name must be provided to Apply")
	}
	result = &v1.ReplicaSet{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("replicasets").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *replicaSets) ApplyStatus(ctx context.Context, replicaSet *appsv1.ReplicaSetApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ReplicaSet, err error) {
	if replicaSet == nil {
		return nil, fmt.Errorf("replicaSet provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(replicaSet)
	if err != nil {
		return nil, err
	}

	name := replicaSet.Name
	if name == nil {
		return nil, fmt.Errorf("replicaSet.Name must be provided to Apply")
	}

	result = &v1.ReplicaSet{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("replicasets").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// GetScale takes name of the replicaSet, and returns the corresponding autoscalingv1.Scale object, and an error if there is any.
func (c *replicaSets) GetScale(ctx context.Context, replicaSetName string, options metav1.GetOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("replicasets").
		Name(replicaSetName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *replicaSets) UpdateScale(ctx context.Context, replicaSetName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("replicasets").
		Name(replicaSetName).
		SubResource("scale").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scale).
		Do(ctx).
		Into(result)
	return
}

// ApplyScale takes top resource name and the apply declarative configuration for scale,
// applies it and returns the applied scale, and an error, if there is any.
func (c *replicaSets) ApplyScale(ctx context.Context, replicaSetName string, scale *applyconfigurationsautoscalingv1.ScaleApplyConfiguration, opts metav1.ApplyOptions) (result *autoscalingv1.Scale, err error) {
	if scale == nil {
		return nil, fmt.Errorf("scale provided to ApplyScale must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(scale)
	if err != nil {
		return nil, err
	}

	result = &autoscalingv1.Scale{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("replicasets").
		Name(replicaSetName).
		SubResource("scale").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
