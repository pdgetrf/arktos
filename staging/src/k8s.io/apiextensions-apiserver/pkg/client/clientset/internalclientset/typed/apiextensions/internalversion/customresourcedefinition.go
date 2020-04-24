/*
Copyright The Kubernetes Authors.
Copyright 2020 Authors of Arktos - file modified.

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

package internalversion

import (
	"time"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	scheme "k8s.io/apiextensions-apiserver/pkg/client/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CustomResourceDefinitionsGetter has a method to return a CustomResourceDefinitionInterface.
// A group's client should implement this interface.
type CustomResourceDefinitionsGetter interface {
	CustomResourceDefinitions() CustomResourceDefinitionInterface
	CustomResourceDefinitionsWithMultiTenancy(tenant string) CustomResourceDefinitionInterface
}

// CustomResourceDefinitionInterface has methods to work with CustomResourceDefinition resources.
type CustomResourceDefinitionInterface interface {
	Create(*apiextensions.CustomResourceDefinition) (*apiextensions.CustomResourceDefinition, error)
	Update(*apiextensions.CustomResourceDefinition) (*apiextensions.CustomResourceDefinition, error)
	UpdateStatus(*apiextensions.CustomResourceDefinition) (*apiextensions.CustomResourceDefinition, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*apiextensions.CustomResourceDefinition, error)
	List(opts v1.ListOptions) (*apiextensions.CustomResourceDefinitionList, error)
	Watch(opts v1.ListOptions) watch.AggregatedWatchInterface
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apiextensions.CustomResourceDefinition, err error)
	CustomResourceDefinitionExpansion
}

// customResourceDefinitions implements CustomResourceDefinitionInterface
type customResourceDefinitions struct {
	client  rest.Interface
	clients []rest.Interface
	te      string
}

// newCustomResourceDefinitions returns a CustomResourceDefinitions
func newCustomResourceDefinitions(c *ApiextensionsClient) *customResourceDefinitions {
	return newCustomResourceDefinitionsWithMultiTenancy(c, "default")
}

func newCustomResourceDefinitionsWithMultiTenancy(c *ApiextensionsClient, tenant string) *customResourceDefinitions {
	return &customResourceDefinitions{
		client:  c.RESTClient(),
		clients: c.RESTClients(),
		te:      tenant,
	}
}

// Get takes name of the customResourceDefinition, and returns the corresponding customResourceDefinition object, and an error if there is any.
func (c *customResourceDefinitions) Get(name string, options v1.GetOptions) (result *apiextensions.CustomResourceDefinition, err error) {
	result = &apiextensions.CustomResourceDefinition{}
	err = c.client.Get().
		Tenant(c.te).
		Resource("customresourcedefinitions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)

	return
}

// List takes label and field selectors, and returns the list of CustomResourceDefinitions that match those selectors.
func (c *customResourceDefinitions) List(opts v1.ListOptions) (result *apiextensions.CustomResourceDefinitionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &apiextensions.CustomResourceDefinitionList{}
	err = c.client.Get().
		Tenant(c.te).
		Resource("customresourcedefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)

	return
}

// Watch returns a watch.Interface that watches the requested customResourceDefinitions.
func (c *customResourceDefinitions) Watch(opts v1.ListOptions) watch.AggregatedWatchInterface {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	aggWatch := watch.NewAggregatedWatcher()
	for _, client := range c.clients {
		watcher, err := client.Get().
			Tenant(c.te).
			Resource("customresourcedefinitions").
			VersionedParams(&opts, scheme.ParameterCodec).
			Timeout(timeout).
			Watch()
		aggWatch.AddWatchInterface(watcher, err)
	}
	return aggWatch
}

// Create takes the representation of a customResourceDefinition and creates it.  Returns the server's representation of the customResourceDefinition, and an error, if there is any.
func (c *customResourceDefinitions) Create(customResourceDefinition *apiextensions.CustomResourceDefinition) (result *apiextensions.CustomResourceDefinition, err error) {
	result = &apiextensions.CustomResourceDefinition{}

	err = c.client.Post().
		Tenant(c.te).
		Resource("customresourcedefinitions").
		Body(customResourceDefinition).
		Do().
		Into(result)

	return
}

// Update takes the representation of a customResourceDefinition and updates it. Returns the server's representation of the customResourceDefinition, and an error, if there is any.
func (c *customResourceDefinitions) Update(customResourceDefinition *apiextensions.CustomResourceDefinition) (result *apiextensions.CustomResourceDefinition, err error) {
	result = &apiextensions.CustomResourceDefinition{}

	err = c.client.Put().
		Tenant(c.te).
		Resource("customresourcedefinitions").
		Name(customResourceDefinition.Name).
		Body(customResourceDefinition).
		Do().
		Into(result)

	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *customResourceDefinitions) UpdateStatus(customResourceDefinition *apiextensions.CustomResourceDefinition) (result *apiextensions.CustomResourceDefinition, err error) {
	result = &apiextensions.CustomResourceDefinition{}

	err = c.client.Put().
		Tenant(c.te).
		Resource("customresourcedefinitions").
		Name(customResourceDefinition.Name).
		SubResource("status").
		Body(customResourceDefinition).
		Do().
		Into(result)

	return
}

// Delete takes name of the customResourceDefinition and deletes it. Returns an error if one occurs.
func (c *customResourceDefinitions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Tenant(c.te).
		Resource("customresourcedefinitions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *customResourceDefinitions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Tenant(c.te).
		Resource("customresourcedefinitions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched customResourceDefinition.
func (c *customResourceDefinitions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apiextensions.CustomResourceDefinition, err error) {
	result = &apiextensions.CustomResourceDefinition{}
	err = c.client.Patch(pt).
		Tenant(c.te).
		Resource("customresourcedefinitions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)

	return
}
