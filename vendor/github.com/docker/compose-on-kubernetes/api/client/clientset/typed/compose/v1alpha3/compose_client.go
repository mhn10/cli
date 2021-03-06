package v1alpha3

import (
	"github.com/docker/compose-on-kubernetes/api/client/clientset/scheme"
	v1alpha3 "github.com/docker/compose-on-kubernetes/api/compose/v1alpha3"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

// ComposeV1alpha3Interface defines the methods a compose v1alpha3 client has
type ComposeV1alpha3Interface interface {
	RESTClient() rest.Interface
	StacksGetter
}

// ComposeV1alpha3Client is used to interact with features provided by the compose.docker.com group.
type ComposeV1alpha3Client struct {
	restClient rest.Interface
}

// Stacks returns a stack client
func (c *ComposeV1alpha3Client) Stacks(namespace string) StackInterface {
	return newStacks(c, namespace)
}

// NewForConfig creates a new ComposeV1alpha3Client for the given config.
func NewForConfig(c *rest.Config) (*ComposeV1alpha3Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ComposeV1alpha3Client{client}, nil
}

// NewForConfigOrDie creates a new ComposeV1alpha3Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ComposeV1alpha3Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ComposeV1alpha3Client for the given RESTClient.
func New(c rest.Interface) *ComposeV1alpha3Client {
	return &ComposeV1alpha3Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha3.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ComposeV1alpha3Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
