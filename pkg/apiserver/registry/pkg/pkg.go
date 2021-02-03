package pkg

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/packages/v1alpha1"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
)

type REST struct {
	cache cache.Indexer
}

var (
	_ rest.Storage = &REST{}
	_ rest.Scoper  = &REST{}
	_ rest.Lister  = &REST{}
	_ rest.Getter  = &REST{}
)

func NewREST() *REST {
	keyFunc := func(obj interface{}) (string, error) {
		pkg, ok := obj.(*v1alpha1.Pkg)
		if !ok {
			return "", fmt.Errorf("object is not a Pkg: %v", obj)
		}
		return pkg.Name, nil
	}
	return &REST{cache.NewIndexer(keyFunc, make(cache.Indexers))}
}

// Storage
func (r *REST) New() runtime.Object {
	return &v1alpha1.Pkg{}
}

func (r *REST) NewList() runtime.Object {
	return &v1alpha1.PkgList{}
}

func (r *REST) NamespaceScoped() bool { return false }

func (r *REST) List(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
	labelSelector := labels.Everything()
	if options != nil && options.LabelSelector != nil {
		labelSelector = options.LabelSelector
	}
	pkgs := r.cache.List()
	items := make([]v1alpha1.Pkg, 0, len(pkgs))
	for i := range pkgs {
		item := (pkgs[i].(*v1alpha1.Pkg)).DeepCopy()
		if labelSelector.Matches(labels.Set(item.Labels)) {
			items = append(items, *item)
		}
	}
	list := &v1alpha1.PkgList{Items: items}
	return list, nil
}

func (r *REST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	pkgObj, exists, err := r.cache.Get(name)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pkg"), name)
	}

	return (pkgObj.(*v1alpha1.Pkg)).DeepCopy(), nil
}

func (r *REST) ConvertToTable(ctx context.Context, obj runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	return rest.NewDefaultTableConvertor(v1alpha1.Resource("pkg")).ConvertToTable(ctx, obj, tableOptions)
}
