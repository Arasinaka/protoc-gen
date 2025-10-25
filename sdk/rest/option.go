package rest

import (
	"context"

	"github.com/Arasinaka/protoc-gen/sdk/configs"
)

type options struct {
	Ctx          context.Context `json:"-"`
	TenantID     string          `json:"tenant_id"`
	ClientID     string          `json:"client_id"`
	ClientSecret string          `json:"-"`

	TokenURL string   `json:"token_url"`
	Host     string   `json:"host"`
	BasePath string   `json:"base_path"`
	Schemes  []string `json:"schemes"`
	Scopes   []string `json:"scopes"`

	LogEnabled   bool `json:"log_enabled"`
	TraceEnabled bool `json:"trace_enabled"`

	SkipCheck bool `json:"-"`
}

func defaultOptions() options {
	conf := configs.GetConfig()

	schemes := []string{"http"}
	if conf.TLSEnabled {
		schemes = []string{"https"}
	}

	return options{
		Ctx:          context.Background(),
		TenantID:     conf.TenantID,
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		TokenURL:     conf.TokenURL,
		Host:         conf.Host,
		BasePath:     "",
		Schemes:      schemes,
		Scopes:       nil,
		LogEnabled:   conf.LogEnabled,
		TraceEnabled: conf.TraceEnabled,
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

func BasePath(basePath string) Option {
	return func(o *options) {
		o.BasePath = basePath
	}
}

func Schemes(schemes ...string) Option {
	return func(o *options) {
		o.Schemes = schemes
	}
}

func Scopes(scopes ...string) Option {
	return func(o *options) {
		o.Scopes = scopes
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

func SkipCheck() Option {
	return func(o *options) {
		o.SkipCheck = true
	}
}
