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

	v1 "github.com/spotmaxtech/k8s-api-v02217/core/v1"
	metav1 "github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/apis/meta/v1"
	types "github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/types"
	watch "github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/watch"
	corev1 "github.com/spotmaxtech/k8s-client-go-v02217/applyconfigurations/core/v1"
	scheme "github.com/spotmaxtech/k8s-client-go-v02217/kubernetes/scheme"
	rest "github.com/spotmaxtech/k8s-client-go-v02217/rest"
)

// EventsGetter has a method to return a EventInterface.
// A group's client should implement this interface.
type EventsGetter interface {
	Events(namespace string) EventInterface
}

// EventInterface has methods to work with Event resources.
type EventInterface interface {
	Create(ctx context.Context, event *v1.Event, opts metav1.CreateOptions) (*v1.Event, error)
	Update(ctx context.Context, event *v1.Event, opts metav1.UpdateOptions) (*v1.Event, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Event, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.EventList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Event, err error)
	Apply(ctx context.Context, event *corev1.EventApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Event, err error)
	EventExpansion
}

// events implements EventInterface
type events struct {
	client rest.Interface
	ns     string
}

// newEvents returns a Events
func newEvents(c *CoreV1Client, namespace string) *events {
	return &events{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the event, and returns the corresponding event object, and an error if there is any.
func (c *events) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Event, err error) {
	result = &v1.Event{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("events").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Events that match those selectors.
func (c *events) List(ctx context.Context, opts metav1.ListOptions) (result *v1.EventList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.EventList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("events").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested events.
func (c *events) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("events").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a event and creates it.  Returns the server's representation of the event, and an error, if there is any.
func (c *events) Create(ctx context.Context, event *v1.Event, opts metav1.CreateOptions) (result *v1.Event, err error) {
	result = &v1.Event{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("events").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(event).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a event and updates it. Returns the server's representation of the event, and an error, if there is any.
func (c *events) Update(ctx context.Context, event *v1.Event, opts metav1.UpdateOptions) (result *v1.Event, err error) {
	result = &v1.Event{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("events").
		Name(event.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(event).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the event and deletes it. Returns an error if one occurs.
func (c *events) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("events").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *events) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("events").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched event.
func (c *events) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Event, err error) {
	result = &v1.Event{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("events").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied event.
func (c *events) Apply(ctx context.Context, event *corev1.EventApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Event, err error) {
	if event == nil {
		return nil, fmt.Errorf("event provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	name := event.Name
	if name == nil {
		return nil, fmt.Errorf("event.Name must be provided to Apply")
	}
	result = &v1.Event{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("events").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
