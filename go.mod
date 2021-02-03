module github.com/vmware-tanzu/carvel-kapp-controller

go 1.15

require (
	github.com/ghodss/yaml v1.0.0
	github.com/go-logr/logr v0.4.0
	github.com/vmware-tanzu/carvel-vendir v0.15.0
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	k8s.io/code-generator v0.20.2
	sigs.k8s.io/controller-runtime v0.8.1
	sigs.k8s.io/yaml v1.2.0
)
