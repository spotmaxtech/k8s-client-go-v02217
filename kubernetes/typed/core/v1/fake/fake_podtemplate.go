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

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationscorev1 "github.com/spotmaxtech/k8s-client-go-v02217/applyconfigurations/core/v1"
	testing "github.com/spotmaxtech/k8s-client-go-v02217/testing"
)

// FakePodTemplates implements PodTemplateInterface
type FakePodTemplates struct {
	Fake *FakeCoreV1
	ns   string
}

var podtemplatesResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "podtemplates"}

var podtemplatesKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "PodTemplate"}

// Get takes name of the podTemplate, and returns the corresponding podTemplate object, and an error if there is any.
func (c *FakePodTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.PodTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(podtemplatesResource, c.ns, name), &corev1.PodTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PodTemplate), err
}

// List takes label and field selectors, and returns the list of PodTemplates that match those selectors.
func (c *FakePodTemplates) List(ctx context.Context, opts v1.ListOptions) (result *corev1.PodTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(podtemplatesResource, podtemplatesKind, c.ns, opts), &corev1.PodTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.PodTemplateList{ListMeta: obj.(*corev1.PodTemplateList).ListMeta}
	for _, item := range obj.(*corev1.PodTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podTemplates.
func (c *FakePodTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(podtemplatesResource, c.ns, opts))

}

// Create takes the representation of a podTemplate and creates it.  Returns the server's representation of the podTemplate, and an error, if there is any.
func (c *FakePodTemplates) Create(ctx context.Context, podTemplate *corev1.PodTemplate, opts v1.CreateOptions) (result *corev1.PodTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podtemplatesResource, c.ns, podTemplate), &corev1.PodTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PodTemplate), err
}

// Update takes the representation of a podTemplate and updates it. Returns the server's representation of the podTemplate, and an error, if there is any.
func (c *FakePodTemplates) Update(ctx context.Context, podTemplate *corev1.PodTemplate, opts v1.UpdateOptions) (result *corev1.PodTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(podtemplatesResource, c.ns, podTemplate), &corev1.PodTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PodTemplate), err
}

// Delete takes name of the podTemplate and deletes it. Returns an error if one occurs.
func (c *FakePodTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(podtemplatesResource, c.ns, name), &corev1.PodTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(podtemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.PodTemplateList{})
	return err
}

// Patch applies the patch and returns the patched podTemplate.
func (c *FakePodTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.PodTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podtemplatesResource, c.ns, name, pt, data, subresources...), &corev1.PodTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PodTemplate), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied podTemplate.
func (c *FakePodTemplates) Apply(ctx context.Context, podTemplate *applyconfigurationscorev1.PodTemplateApplyConfiguration, opts v1.ApplyOptions) (result *corev1.PodTemplate, err error) {
	if podTemplate == nil {
		return nil, fmt.Errorf("podTemplate provided to Apply must not be nil")
	}
	data, err := json.Marshal(podTemplate)
	if err != nil {
		return nil, err
	}
	name := podTemplate.Name
	if name == nil {
		return nil, fmt.Errorf("podTemplate.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podtemplatesResource, c.ns, *name, types.ApplyPatchType, data), &corev1.PodTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PodTemplate), err
}
