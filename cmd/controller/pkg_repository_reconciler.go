package controller

import (
	"context"

	"github.com/go-logr/logr"
	kcclient "github.com/vmware-tanzu/carvel-kapp-controller/pkg/client/clientset/versioned"
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/pkgrepository"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type PkgRepositoryReconciler struct {
	client kcclient.Interface
	log    logr.Logger
}

var _ reconcile.Reconciler = &PkgRepositoryReconciler{}

func (r *PkgRepositoryReconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	log := r.log.WithValues("request", request)

	// TODO currently we've decided to get a fresh copy of app so
	// that we do not operate on stale copy for efficiency reasons
	existingPkgRepository, err := r.client.KappctrlV1alpha1().PkgRepositories().Get(ctx, request.Name, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			log.Info("Could not find PkgRepository", "name", request.Name)
			return reconcile.Result{}, nil // No requeue
		}

		log.Error(err, "Could not fetch PkgRepository")
		return reconcile.Result{}, err
	}

	return pkgrepository.NewPkgRepositoryCR(existingPkgRepository, log, r.client).Reconcile()
}
