package rest

import (
	"context"
	"fmt"
	"net/http"

	"codeup.aliyun.com/6684bcca4b50db3fde04b971/dr-srv/platform-utils/jaeger"
	"codeup.aliyun.com/6684bcca4b50db3fde04b971/dr-srv/platform-utils/log"
	"codeup.aliyun.com/6684bcca4b50db3fde04b971/dr-srv/platform-utils/tenant"
	"codeup.aliyun.com/6684bcca4b50db3fde04b971/dr-srv/platform-utils/tools"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func NewClientTransport(options ...Option) (runtime.ClientTransport, error) {
	opts := defaultOptions()

	for _, opt := range options {
		if opt != nil {
			opt(&opts)
		}
	}

	if err := opts.check(); err != nil {
		return nil, err
	}

	return opts.build()
}

func (opts *options) check() error {
	if opts.SkipCheck {
		return nil
	}

	if len(opts.TenantID) == 0 {
		return fmt.Errorf("Tenant ID required. Check the env 'TENANT_ID' or use 'TenantID' option")
	}

	if len(opts.ClientID) == 0 {
		return fmt.Errorf("Client ID required. Check the env 'CLIENT_ID' or use 'ClientID' option")
	}

	if len(opts.ClientSecret) == 0 {
		return fmt.Errorf("Secret required. Check the env 'CLIENT_SECRET' or use 'ClientSecret' option")
	}

	if len(opts.BasePath) == 0 {
		return fmt.Errorf("Base path required. Use 'BasePath' option")
	}

	return nil
}

func (opts *options) build() (runtime.ClientTransport, error) {
	logrus.Infof("[OpenAPI] Client options: %s", tools.StringStruct(opts))

	transport := http.DefaultTransport
	ctx := context.WithValue(opts.Ctx, oauth2.HTTPClient, &http.Client{Transport: transport})

	clientConfig := &clientcredentials.Config{
		ClientID:     opts.ClientID,
		ClientSecret: opts.ClientSecret,
		TokenURL:     opts.TokenURL,
		Scopes:       opts.Scopes,
	}
	httpClient := clientConfig.Client(ctx)

	httpClient.Transport = tenant.NewRoundTripper(opts.TenantID, httpClient.Transport)
	if opts.LogEnabled {
		httpClient.Transport = log.NewRoundTripper(httpClient.Transport)
	}
	if opts.TraceEnabled {
		httpClient.Transport = jaeger.NewRoundTripper(httpClient.Transport)
	}

	newTransport := httptransport.NewWithClient(
		opts.Host,
		opts.BasePath,
		opts.Schemes,
		httpClient,
	)

	return newTransport, nil
}
