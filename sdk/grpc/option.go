package grpc

import (
	"context"

	"github.com/Arasinaka/protoc-gen/sdk/configs"

	"google.golang.org/grpc"
)

type options struct {
	Ctx          context.Context `json:"-"`
	TenantID     string          `json:"tenant_id"`
	ClientID     string          `json:"client_id"`
	ClientSecret string          `json:"-"`

	TokenURL string   `json:"token_url"`
	Host     string   `json:"host"`
	Scopes   []string `json:"scopes"`

	TLSEnabled   bool `json:"tls_enabled"`
	LogEnabled   bool `json:"log_enabled"`
	TraceEnabled bool `json:"trace_enabled"`

	UnaryInterceptors  []grpc.UnaryClientInterceptor  `json:"-"`
	StreamInterceptors []grpc.StreamClientInterceptor `json:"-"`

	SkipCheck bool `json:"-"`
}

func defaultOptions() options {
	conf := configs.GetConfig()

	return options{
		Ctx:                context.Background(),
		TenantID:           conf.TenantID,
		ClientID:           conf.ClientID,
		ClientSecret:       conf.ClientSecret,
		TokenURL:           conf.TokenURL,
		Host:               conf.Host,
		Scopes:             nil,
		TLSEnabled:         conf.TLSEnabled,
		LogEnabled:         conf.LogEnabled,
		TraceEnabled:       conf.TraceEnabled,
		UnaryInterceptors:  nil,
		StreamInterceptors: nil,
	}
}

type Option func(*options)

func Context(ctx context.Context) Option {
	return func(o *options) {
		o.Ctx = ctx
	}
}

func TenantID(tenantID string) Option {
	return func(o *options) {
		o.TenantID = tenantID
	}
}

func ClientID(clientID string) Option {
	return func(o *options) {
		o.ClientID = clientID
	}
}

func ClientSecret(clientSecret string) Option {
	return func(o *options) {
		o.ClientSecret = clientSecret
	}
}

func TokenURL(tokenURL string) Option {
	return func(o *options) {
		o.TokenURL = tokenURL
	}
}

func Host(host string) Option {
	return func(o *options) {
		o.Host = host
	}
}

func Scopes(scopes ...string) Option {
	return func(o *options) {
		o.Scopes = scopes
	}
}

func TLSEnabled(tlsEnabled bool) Option {
	return func(o *options) {
		o.TLSEnabled = tlsEnabled
	}
}

func TraceEnabled(traceEnabled bool) Option {
	return func(o *options) {
		o.TraceEnabled = traceEnabled
	}
}

func LogEnabled(logEnabled bool) Option {
	return func(o *options) {
		o.LogEnabled = logEnabled
	}
}

func UnaryInterceptors(interceptors ...grpc.UnaryClientInterceptor) Option {
	return func(o *options) {
		o.UnaryInterceptors = interceptors
	}
}

func StreamInterceptors(interceptors ...grpc.StreamClientInterceptor) Option {
	return func(o *options) {
		o.StreamInterceptors = interceptors
	}
}

func SkipCheck() Option {
	return func(o *options) {
		o.SkipCheck = true
	}
}
