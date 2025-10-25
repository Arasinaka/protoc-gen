package grpc

import (
	"context"
	"fmt"

	"codeup.aliyun.com/6684bcca4b50db3fde04b971/dr-srv/platform-utils/jaeger"
	"codeup.aliyun.com/6684bcca4b50db3fde04b971/dr-srv/platform-utils/log"
	"codeup.aliyun.com/6684bcca4b50db3fde04b971/dr-srv/platform-utils/tenant"
	"codeup.aliyun.com/6684bcca4b50db3fde04b971/dr-srv/platform-utils/tools"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/clientcredentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/keepalive"
)

func NewClientConn(options ...Option) (*grpc.ClientConn, error) {
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
		return fmt.Errorf("Tenant ID required. Check the env 'HORIZON_TENANT_ID' or use 'TenantID' option")
	}

	if len(opts.ClientID) == 0 {
		return fmt.Errorf("Client ID required. Check the env 'HORIZON_CLIENT_ID' or use 'ClientID' option")
	}

	if len(opts.ClientSecret) == 0 {
		return fmt.Errorf("Secret required. Check the env 'HORIZON_CLIENT_SECRET' or use 'ClientSecret' option")
	}

	return nil
}

func (opts *options) build() (*grpc.ClientConn, error) {
	logrus.Infof("[gRPC] Client options: %s", tools.StringStruct(opts))

	clientConfig := &clientcredentials.Config{
		ClientID:     opts.ClientID,
		ClientSecret: opts.ClientSecret,
		TokenURL:     opts.TokenURL,
		Scopes:       opts.Scopes,
	}
	tokenSource := oauth.TokenSource{TokenSource: clientConfig.TokenSource(opts.Ctx)}

	dialOpts := make([]grpc.DialOption, 0)
	dialOpts = append(dialOpts, grpc.WithPerRPCCredentials(&perRPCCredentials{tokenSource}))
	dialOpts = append(dialOpts, grpc.WithUnaryInterceptor(tenant.UnaryClientInterceptor(opts.TenantID)))
	dialOpts = append(dialOpts, grpc.WithStreamInterceptor(tenant.StreamClientInterceptor(opts.TenantID)))
	dialOpts = append(dialOpts, grpc.WithKeepaliveParams(keepalive.ClientParameters{PermitWithoutStream: true}))
	dialOpts = append(dialOpts, grpc.WithChainUnaryInterceptor(opts.UnaryInterceptors...))
	dialOpts = append(dialOpts, grpc.WithChainStreamInterceptor(opts.StreamInterceptors...))

	if opts.TLSEnabled {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	} else {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	if opts.LogEnabled {
		dialOpts = append(dialOpts, grpc.WithChainUnaryInterceptor(log.UnaryClientInterceptor()))
		dialOpts = append(dialOpts, grpc.WithStreamInterceptor(log.StreamClientInterceptor()))
	}

	if opts.TraceEnabled {
		dialOpts = append(dialOpts, grpc.WithChainUnaryInterceptor(jaeger.UnaryClientInterceptor()))
		dialOpts = append(dialOpts, grpc.WithStreamInterceptor(jaeger.StreamClientInterceptor()))
	}

	return grpc.DialContext(opts.Ctx, opts.Host, dialOpts...)
}

type perRPCCredentials struct {
	oauth.TokenSource
}

func (ts perRPCCredentials) RequireTransportSecurity() bool {
	return false
}

func (ts perRPCCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	token, err := ts.Token()
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"authorization": token.Type() + " " + token.AccessToken,
	}, nil
}
