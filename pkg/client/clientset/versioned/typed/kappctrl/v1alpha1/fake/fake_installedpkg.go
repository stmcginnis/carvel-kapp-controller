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

// FakeInstalledPkgs implements InstalledPkgInterface
type FakeInstalledPkgs struct {
	Fake *FakeKappctrlV1alpha1
	ns   string
}

var installedpkgsResource = schema.GroupVersionResource{Group: "kappctrl", Version: "v1alpha1", Resource: "installedpkgs"}

var installedpkgsKind = schema.GroupVersionKind{Group: "kappctrl", Version: "v1alpha1", Kind: "InstalledPkg"}

// Get takes name of the installedPkg, and returns the corresponding installedPkg object, and an error if there is any.
func (c *FakeInstalledPkgs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.InstalledPkg, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(installedpkgsResource, c.ns, name), &v1alpha1.InstalledPkg{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstalledPkg), err
}

// List takes label and field selectors, and returns the list of InstalledPkgs that match those selectors.
func (c *FakeInstalledPkgs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InstalledPkgList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(installedpkgsResource, installedpkgsKind, c.ns, opts), &v1alpha1.InstalledPkgList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InstalledPkgList{ListMeta: obj.(*v1alpha1.InstalledPkgList).ListMeta}
	for _, item := range obj.(*v1alpha1.InstalledPkgList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested installedPkgs.
func (c *FakeInstalledPkgs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(installedpkgsResource, c.ns, opts))

}

// Create takes the representation of a installedPkg and creates it.  Returns the server's representation of the installedPkg, and an error, if there is any.
func (c *FakeInstalledPkgs) Create(ctx context.Context, installedPkg *v1alpha1.InstalledPkg, opts v1.CreateOptions) (result *v1alpha1.InstalledPkg, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(installedpkgsResource, c.ns, installedPkg), &v1alpha1.InstalledPkg{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstalledPkg), err
}

// Update takes the representation of a installedPkg and updates it. Returns the server's representation of the installedPkg, and an error, if there is any.
func (c *FakeInstalledPkgs) Update(ctx context.Context, installedPkg *v1alpha1.InstalledPkg, opts v1.UpdateOptions) (result *v1alpha1.InstalledPkg, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(installedpkgsResource, c.ns, installedPkg), &v1alpha1.InstalledPkg{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstalledPkg), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInstalledPkgs) UpdateStatus(ctx context.Context, installedPkg *v1alpha1.InstalledPkg, opts v1.UpdateOptions) (*v1alpha1.InstalledPkg, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(installedpkgsResource, "status", c.ns, installedPkg), &v1alpha1.InstalledPkg{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstalledPkg), err
}

// Delete takes name of the installedPkg and deletes it. Returns an error if one occurs.
func (c *FakeInstalledPkgs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(installedpkgsResource, c.ns, name), &v1alpha1.InstalledPkg{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInstalledPkgs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(installedpkgsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InstalledPkgList{})
	return err
}

// Patch applies the patch and returns the patched installedPkg.
func (c *FakeInstalledPkgs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InstalledPkg, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(installedpkgsResource, c.ns, name, pt, data, subresources...), &v1alpha1.InstalledPkg{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstalledPkg), err
}
