/*
Copyright 2018 Google LLC

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
	v1beta1 "github.com/grafeas/kritis/pkg/kritis/apis/kritis/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAttestors implements AttestorInterface
type FakeAttestors struct {
	Fake *FakeKritisV1beta1
	ns   string
}

var attestorsResource = schema.GroupVersionResource{Group: "kritis", Version: "v1beta1", Resource: "attestors"}

var attestorsKind = schema.GroupVersionKind{Group: "kritis", Version: "v1beta1", Kind: "Attestor"}

// Get takes name of the attestor, and returns the corresponding attestor object, and an error if there is any.
func (c *FakeAttestors) Get(name string, options v1.GetOptions) (result *v1beta1.Attestor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(attestorsResource, c.ns, name), &v1beta1.Attestor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Attestor), err
}

// List takes label and field selectors, and returns the list of Attestors that match those selectors.
func (c *FakeAttestors) List(opts v1.ListOptions) (result *v1beta1.AttestorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(attestorsResource, attestorsKind, c.ns, opts), &v1beta1.AttestorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.AttestorList{}
	for _, item := range obj.(*v1beta1.AttestorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested attestors.
func (c *FakeAttestors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(attestorsResource, c.ns, opts))

}

// Create takes the representation of a attestor and creates it.  Returns the server's representation of the attestor, and an error, if there is any.
func (c *FakeAttestors) Create(attestor *v1beta1.Attestor) (result *v1beta1.Attestor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(attestorsResource, c.ns, attestor), &v1beta1.Attestor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Attestor), err
}

// Update takes the representation of a attestor and updates it. Returns the server's representation of the attestor, and an error, if there is any.
func (c *FakeAttestors) Update(attestor *v1beta1.Attestor) (result *v1beta1.Attestor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(attestorsResource, c.ns, attestor), &v1beta1.Attestor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Attestor), err
}

// Delete takes name of the attestor and deletes it. Returns an error if one occurs.
func (c *FakeAttestors) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(attestorsResource, c.ns, name), &v1beta1.Attestor{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAttestors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(attestorsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.AttestorList{})
	return err
}

// Patch applies the patch and returns the patched attestor.
func (c *FakeAttestors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Attestor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(attestorsResource, c.ns, name, data, subresources...), &v1beta1.Attestor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Attestor), err
}
