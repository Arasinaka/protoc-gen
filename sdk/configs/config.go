package configs

import (
	"sync"

	"codeup.aliyun.com/6684bcca4b50db3fde04b971/dr-srv/platform-utils/common"
	"github.com/spf13/viper"
)

const (
	TenantIDKey     = "TENANT_ID"
	ClientIDKey     = "CLIENT_ID"
	ClientSecretKey = "CLIENT_SECRET"
	TokenUrlKey     = "TOKEN_URL"
	HostKey         = "HOST"
	TLSEnabledKey   = "TLS_ENABLED"
	LogEnabledKey   = "LOG_ENABLED"
	TraceEnabledKey = "TRACE_ENABLED"
)

type Config struct {
	TenantID     string
	ClientID     string
	ClientSecret string
	TokenURL     string
	Host         string
	TLSEnabled   bool
	LogEnabled   bool
	TraceEnabled bool
}

func NewConfig() *Config {
	v := viper.New()
	v.AutomaticEnv()

	v.SetDefault(ClientIDKey, "")
	v.SetDefault(TenantIDKey, "")
	v.SetDefault(ClientSecretKey, "")
	v.SetDefault(TokenUrlKey, "https://oauth.hpbp.io/oauth/v1/token")
	v.SetDefault(HostKey, "api.hpbp.io:443")
	v.SetDefault(TLSEnabledKey, true)
	v.SetDefault(LogEnabledKey, true)
	v.SetDefault(TraceEnabledKey, false)

	common.LoadFromFile(v)

	return &Config{
		TenantID:     v.GetString(TenantIDKey),
		ClientID:     v.GetString(ClientIDKey),
		ClientSecret: v.GetString(ClientSecretKey),
		TokenURL:     v.GetString(TokenUrlKey),
		Host:         v.GetString(HostKey),
		TLSEnabled:   v.GetBool(TLSEnabledKey),
		LogEnabled:   v.GetBool(LogEnabledKey),
		TraceEnabled: v.GetBool(TraceEnabledKey),
	}
}

var conf *Config
var once sync.Once

func GetConfig() *Config {
	if conf == nil {
		once.Do(func() {
			conf = NewConfig()
		})
	}
	return conf
}
