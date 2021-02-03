// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePkgs implements PkgInterface
type FakePkgs struct {
	Fake *FakeKappctrlV1alpha1
}

var pkgsResource = schema.GroupVersionResource{Group: "kappctrl", Version: "v1alpha1", Resource: "pkgs"}

var pkgsKind = schema.GroupVersionKind{Group: "kappctrl", Version: "v1alpha1", Kind: "Pkg"}

// Get takes name of the pkg, and returns the corresponding pkg object, and an error if there is any.
func (c *FakePkgs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Pkg, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(pkgsResource, name), &v1alpha1.Pkg{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Pkg), err
}

// List takes label and field selectors, and returns the list of Pkgs that match those selectors.
func (c *FakePkgs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PkgList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(pkgsResource, pkgsKind, opts), &v1alpha1.PkgList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PkgList{ListMeta: obj.(*v1alpha1.PkgList).ListMeta}
	for _, item := range obj.(*v1alpha1.PkgList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pkgs.
func (c *FakePkgs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(pkgsResource, opts))
}

// Create takes the representation of a pkg and creates it.  Returns the server's representation of the pkg, and an error, if there is any.
func (c *FakePkgs) Create(ctx context.Context, pkg *v1alpha1.Pkg, opts v1.CreateOptions) (result *v1alpha1.Pkg, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(pkgsResource, pkg), &v1alpha1.Pkg{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Pkg), err
}

// Update takes the representation of a pkg and updates it. Returns the server's representation of the pkg, and an error, if there is any.
func (c *FakePkgs) Update(ctx context.Context, pkg *v1alpha1.Pkg, opts v1.UpdateOptions) (result *v1alpha1.Pkg, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(pkgsResource, pkg), &v1alpha1.Pkg{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Pkg), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePkgs) UpdateStatus(ctx context.Context, pkg *v1alpha1.Pkg, opts v1.UpdateOptions) (*v1alpha1.Pkg, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(pkgsResource, "status", pkg), &v1alpha1.Pkg{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Pkg), err
}

// Delete takes name of the pkg and deletes it. Returns an error if one occurs.
func (c *FakePkgs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(pkgsResource, name), &v1alpha1.Pkg{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePkgs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(pkgsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PkgList{})
	return err
}

// Patch applies the patch and returns the patched pkg.
func (c *FakePkgs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Pkg, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(pkgsResource, name, pt, data, subresources...), &v1alpha1.Pkg{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Pkg), err
}
