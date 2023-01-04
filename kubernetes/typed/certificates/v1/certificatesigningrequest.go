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

	v1 "k8s.io/api/certificates/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	certificatesv1 "github.com/spotmaxtech/k8s-client-go-v02217/applyconfigurations/certificates/v1"
	scheme "github.com/spotmaxtech/k8s-client-go-v02217/kubernetes/scheme"
	rest "github.com/spotmaxtech/k8s-client-go-v02217/rest"
)

// CertificateSigningRequestsGetter has a method to return a CertificateSigningRequestInterface.
// A group's client should implement this interface.
type CertificateSigningRequestsGetter interface {
	CertificateSigningRequests() CertificateSigningRequestInterface
}

// CertificateSigningRequestInterface has methods to work with CertificateSigningRequest resources.
type CertificateSigningRequestInterface interface {
	Create(ctx context.Context, certificateSigningRequest *v1.CertificateSigningRequest, opts metav1.CreateOptions) (*v1.CertificateSigningRequest, error)
	Update(ctx context.Context, certificateSigningRequest *v1.CertificateSigningRequest, opts metav1.UpdateOptions) (*v1.CertificateSigningRequest, error)
	UpdateStatus(ctx context.Context, certificateSigningRequest *v1.CertificateSigningRequest, opts metav1.UpdateOptions) (*v1.CertificateSigningRequest, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CertificateSigningRequest, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CertificateSigningRequestList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CertificateSigningRequest, err error)
	Apply(ctx context.Context, certificateSigningRequest *certificatesv1.CertificateSigningRequestApplyConfiguration, opts metav1.ApplyOptions) (result *v1.CertificateSigningRequest, err error)
	ApplyStatus(ctx context.Context, certificateSigningRequest *certificatesv1.CertificateSigningRequestApplyConfiguration, opts metav1.ApplyOptions) (result *v1.CertificateSigningRequest, err error)
	UpdateApproval(ctx context.Context, certificateSigningRequestName string, certificateSigningRequest *v1.CertificateSigningRequest, opts metav1.UpdateOptions) (*v1.CertificateSigningRequest, error)

	CertificateSigningRequestExpansion
}

// certificateSigningRequests implements CertificateSigningRequestInterface
type certificateSigningRequests struct {
	client rest.Interface
}

// newCertificateSigningRequests returns a CertificateSigningRequests
func newCertificateSigningRequests(c *CertificatesV1Client) *certificateSigningRequests {
	return &certificateSigningRequests{
		client: c.RESTClient(),
	}
}

// Get takes name of the certificateSigningRequest, and returns the corresponding certificateSigningRequest object, and an error if there is any.
func (c *certificateSigningRequests) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CertificateSigningRequest, err error) {
	result = &v1.CertificateSigningRequest{}
	err = c.client.Get().
		Resource("certificatesigningrequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CertificateSigningRequests that match those selectors.
func (c *certificateSigningRequests) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CertificateSigningRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CertificateSigningRequestList{}
	err = c.client.Get().
		Resource("certificatesigningrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested certificateSigningRequests.
func (c *certificateSigningRequests) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("certificatesigningrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a certificateSigningRequest and creates it.  Returns the server's representation of the certificateSigningRequest, and an error, if there is any.
func (c *certificateSigningRequests) Create(ctx context.Context, certificateSigningRequest *v1.CertificateSigningRequest, opts metav1.CreateOptions) (result *v1.CertificateSigningRequest, err error) {
	result = &v1.CertificateSigningRequest{}
	err = c.client.Post().
		Resource("certificatesigningrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateSigningRequest).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a certificateSigningRequest and updates it. Returns the server's representation of the certificateSigningRequest, and an error, if there is any.
func (c *certificateSigningRequests) Update(ctx context.Context, certificateSigningRequest *v1.CertificateSigningRequest, opts metav1.UpdateOptions) (result *v1.CertificateSigningRequest, err error) {
	result = &v1.CertificateSigningRequest{}
	err = c.client.Put().
		Resource("certificatesigningrequests").
		Name(certificateSigningRequest.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateSigningRequest).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *certificateSigningRequests) UpdateStatus(ctx context.Context, certificateSigningRequest *v1.CertificateSigningRequest, opts metav1.UpdateOptions) (result *v1.CertificateSigningRequest, err error) {
	result = &v1.CertificateSigningRequest{}
	err = c.client.Put().
		Resource("certificatesigningrequests").
		Name(certificateSigningRequest.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateSigningRequest).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the certificateSigningRequest and deletes it. Returns an error if one occurs.
func (c *certificateSigningRequests) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("certificatesigningrequests").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *certificateSigningRequests) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("certificatesigningrequests").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched certificateSigningRequest.
func (c *certificateSigningRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CertificateSigningRequest, err error) {
	result = &v1.CertificateSigningRequest{}
	err = c.client.Patch(pt).
		Resource("certificatesigningrequests").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied certificateSigningRequest.
func (c *certificateSigningRequests) Apply(ctx context.Context, certificateSigningRequest *certificatesv1.CertificateSigningRequestApplyConfiguration, opts metav1.ApplyOptions) (result *v1.CertificateSigningRequest, err error) {
	if certificateSigningRequest == nil {
		return nil, fmt.Errorf("certificateSigningRequest provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(certificateSigningRequest)
	if err != nil {
		return nil, err
	}
	name := certificateSigningRequest.Name
	if name == nil {
		return nil, fmt.Errorf("certificateSigningRequest.Name must be provided to Apply")
	}
	result = &v1.CertificateSigningRequest{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("certificatesigningrequests").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *certificateSigningRequests) ApplyStatus(ctx context.Context, certificateSigningRequest *certificatesv1.CertificateSigningRequestApplyConfiguration, opts metav1.ApplyOptions) (result *v1.CertificateSigningRequest, err error) {
	if certificateSigningRequest == nil {
		return nil, fmt.Errorf("certificateSigningRequest provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(certificateSigningRequest)
	if err != nil {
		return nil, err
	}

	name := certificateSigningRequest.Name
	if name == nil {
		return nil, fmt.Errorf("certificateSigningRequest.Name must be provided to Apply")
	}

	result = &v1.CertificateSigningRequest{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("certificatesigningrequests").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// UpdateApproval takes the top resource name and the representation of a certificateSigningRequest and updates it. Returns the server's representation of the certificateSigningRequest, and an error, if there is any.
func (c *certificateSigningRequests) UpdateApproval(ctx context.Context, certificateSigningRequestName string, certificateSigningRequest *v1.CertificateSigningRequest, opts metav1.UpdateOptions) (result *v1.CertificateSigningRequest, err error) {
	result = &v1.CertificateSigningRequest{}
	err = c.client.Put().
		Resource("certificatesigningrequests").
		Name(certificateSigningRequestName).
		SubResource("approval").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateSigningRequest).
		Do(ctx).
		Into(result)
	return
}
