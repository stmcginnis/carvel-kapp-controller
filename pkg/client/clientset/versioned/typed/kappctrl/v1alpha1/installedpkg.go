// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	scheme "github.com/vmware-tanzu/carvel-kapp-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// InstalledPkgsGetter has a method to return a InstalledPkgInterface.
// A group's client should implement this interface.
type InstalledPkgsGetter interface {
	InstalledPkgs(namespace string) InstalledPkgInterface
}

// InstalledPkgInterface has methods to work with InstalledPkg resources.
type InstalledPkgInterface interface {
	Create(ctx context.Context, installedPkg *v1alpha1.InstalledPkg, opts v1.CreateOptions) (*v1alpha1.InstalledPkg, error)
	Update(ctx context.Context, installedPkg *v1alpha1.InstalledPkg, opts v1.UpdateOptions) (*v1alpha1.InstalledPkg, error)
	UpdateStatus(ctx context.Context, installedPkg *v1alpha1.InstalledPkg, opts v1.UpdateOptions) (*v1alpha1.InstalledPkg, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.InstalledPkg, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.InstalledPkgList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InstalledPkg, err error)
	InstalledPkgExpansion
}

// installedPkgs implements InstalledPkgInterface
type installedPkgs struct {
	client rest.Interface
	ns     string
}

// newInstalledPkgs returns a InstalledPkgs
func newInstalledPkgs(c *KappctrlV1alpha1Client, namespace string) *installedPkgs {
	return &installedPkgs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the installedPkg, and returns the corresponding installedPkg object, and an error if there is any.
func (c *installedPkgs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.InstalledPkg, err error) {
	result = &v1alpha1.InstalledPkg{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("installedpkgs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of InstalledPkgs that match those selectors.
func (c *installedPkgs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InstalledPkgList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.InstalledPkgList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("installedpkgs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested installedPkgs.
func (c *installedPkgs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("installedpkgs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a installedPkg and creates it.  Returns the server's representation of the installedPkg, and an error, if there is any.
func (c *installedPkgs) Create(ctx context.Context, installedPkg *v1alpha1.InstalledPkg, opts v1.CreateOptions) (result *v1alpha1.InstalledPkg, err error) {
	result = &v1alpha1.InstalledPkg{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("installedpkgs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(installedPkg).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a installedPkg and updates it. Returns the server's representation of the installedPkg, and an error, if there is any.
func (c *installedPkgs) Update(ctx context.Context, installedPkg *v1alpha1.InstalledPkg, opts v1.UpdateOptions) (result *v1alpha1.InstalledPkg, err error) {
	result = &v1alpha1.InstalledPkg{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("installedpkgs").
		Name(installedPkg.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(installedPkg).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *installedPkgs) UpdateStatus(ctx context.Context, installedPkg *v1alpha1.InstalledPkg, opts v1.UpdateOptions) (result *v1alpha1.InstalledPkg, err error) {
	result = &v1alpha1.InstalledPkg{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("installedpkgs").
		Name(installedPkg.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(installedPkg).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the installedPkg and deletes it. Returns an error if one occurs.
func (c *installedPkgs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("installedpkgs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *installedPkgs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("installedpkgs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched installedPkg.
func (c *installedPkgs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InstalledPkg, err error) {
	result = &v1alpha1.InstalledPkg{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("installedpkgs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
