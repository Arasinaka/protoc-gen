package configs

import (
	"sync"

	"codeup.aliyun.com/6684bcca4b50db3fde04b971/dr-srv/platform-utils/common"
	"github.com/spf13/viper"
)

const (
	TenantIDKey     = "HORIZON_TENANT_ID"
	ClientIDKey     = "HORIZON_CLIENT_ID"
	ClientSecretKey = "HORIZON_CLIENT_SECRET"
	TokenUrlKey     = "HORIZON_TOKEN_URL"
	HostKey         = "HORIZON_HOST"
	TLSEnabledKey   = "HORIZON_TLS_ENABLED"
	LogEnabledKey   = "HORIZON_LOG_ENABLED"
	TraceEnabledKey = "HORIZON_TRACE_ENABLED"
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
