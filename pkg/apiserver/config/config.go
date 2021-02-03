package config

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path"
	"time"

	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	genericopenapi "k8s.io/apiserver/pkg/endpoints/openapi"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/kube-aggregator/pkg/generated/openapi"
)

const (
	bindAddress = "0.0.0.0"
	bindPort    = 10349
	TokenPath   = "/var/run/kapp-ctrl/apiserver/loopback-client-token"
)

var (
	scheme = runtime.NewScheme()
	Codecs = serializer.NewCodecFactory(scheme)
)

func init() {
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
}

func NewConfig() (*genericapiserver.Config, error) {
	secureServing := genericoptions.NewSecureServingOptions().WithLoopback()
	authentication := genericoptions.NewDelegatingAuthenticationOptions()
	authorization := genericoptions.NewDelegatingAuthorizationOptions()

	secureServing.BindAddress = net.ParseIP(bindAddress)
	secureServing.BindPort = bindPort

	serverConfig := genericapiserver.NewConfig(Codecs)
	if err := secureServing.ApplyTo(&serverConfig.SecureServing, &serverConfig.LoopbackClientConfig); err != nil {
		return nil, err
	}
	if err := authentication.ApplyTo(&serverConfig.Authentication, serverConfig.SecureServing, nil); err != nil {
		return nil, err
	}
	if err := authorization.ApplyTo(&serverConfig.Authorization); err != nil {
		return nil, err
	}

	if err := os.MkdirAll(path.Dir(TokenPath), os.ModeDir); err != nil {
		return nil, fmt.Errorf("error when creating dirs of token file: %v", err)
	}

	if err := ioutil.WriteFile(TokenPath, []byte(serverConfig.LoopbackClientConfig.BearerToken), 0600); err != nil {
		return nil, fmt.Errorf("error when writing loopback access token to file: %v", err)
	}

	serverConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(
		openapi.GetOpenAPIDefinitions,
		genericopenapi.NewDefinitionNamer(scheme))
	serverConfig.OpenAPIConfig.Info.Title = "Kapp-controller"
	serverConfig.MinRequestTimeout = int((2 * time.Hour).Seconds())

	return serverConfig, nil
}
