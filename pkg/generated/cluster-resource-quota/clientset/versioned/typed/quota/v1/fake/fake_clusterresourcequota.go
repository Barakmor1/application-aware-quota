/*
Copyright 2023 The AAQ Authors.

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

	v1 "github.com/openshift/api/quota/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterResourceQuotas implements ClusterResourceQuotaInterface
type FakeClusterResourceQuotas struct {
	Fake *FakeQuotaV1
}

var clusterresourcequotasResource = v1.SchemeGroupVersion.WithResource("clusterresourcequotas")

var clusterresourcequotasKind = v1.SchemeGroupVersion.WithKind("ClusterResourceQuota")

// Get takes name of the clusterResourceQuota, and returns the corresponding clusterResourceQuota object, and an error if there is any.
func (c *FakeClusterResourceQuotas) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterresourcequotasResource, name), &v1.ClusterResourceQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterResourceQuota), err
}

// List takes label and field selectors, and returns the list of ClusterResourceQuotas that match those selectors.
func (c *FakeClusterResourceQuotas) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterResourceQuotaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterresourcequotasResource, clusterresourcequotasKind, opts), &v1.ClusterResourceQuotaList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ClusterResourceQuotaList{ListMeta: obj.(*v1.ClusterResourceQuotaList).ListMeta}
	for _, item := range obj.(*v1.ClusterResourceQuotaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterResourceQuotas.
func (c *FakeClusterResourceQuotas) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterresourcequotasResource, opts))
}

// Create takes the representation of a clusterResourceQuota and creates it.  Returns the server's representation of the clusterResourceQuota, and an error, if there is any.
func (c *FakeClusterResourceQuotas) Create(ctx context.Context, clusterResourceQuota *v1.ClusterResourceQuota, opts metav1.CreateOptions) (result *v1.ClusterResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterresourcequotasResource, clusterResourceQuota), &v1.ClusterResourceQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterResourceQuota), err
}

// Update takes the representation of a clusterResourceQuota and updates it. Returns the server's representation of the clusterResourceQuota, and an error, if there is any.
func (c *FakeClusterResourceQuotas) Update(ctx context.Context, clusterResourceQuota *v1.ClusterResourceQuota, opts metav1.UpdateOptions) (result *v1.ClusterResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterresourcequotasResource, clusterResourceQuota), &v1.ClusterResourceQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterResourceQuota), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterResourceQuotas) UpdateStatus(ctx context.Context, clusterResourceQuota *v1.ClusterResourceQuota, opts metav1.UpdateOptions) (*v1.ClusterResourceQuota, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterresourcequotasResource, "status", clusterResourceQuota), &v1.ClusterResourceQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterResourceQuota), err
}

// Delete takes name of the clusterResourceQuota and deletes it. Returns an error if one occurs.
func (c *FakeClusterResourceQuotas) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterresourcequotasResource, name, opts), &v1.ClusterResourceQuota{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterResourceQuotas) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterresourcequotasResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ClusterResourceQuotaList{})
	return err
}

// Patch applies the patch and returns the patched clusterResourceQuota.
func (c *FakeClusterResourceQuotas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterresourcequotasResource, name, pt, data, subresources...), &v1.ClusterResourceQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterResourceQuota), err
}
