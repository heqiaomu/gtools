package gsvc

type Config struct {
	AppDebug bool       `mapstructure:"app_debug" json:"app_debug" yaml:"app_debug"`
	Server   HTTPServer `mapstructure:"server" json:"server" yaml:"server"`
}

type HTTPServer struct {
	API              API          `mapstructure:"api" yaml:"api" json:"api"`
	Web              Web          `mapstructure:"web" yaml:"web" json:"web"`
	AllowCrossDomain bool         `mapstructure:"allow_cross_domain" yaml:"allow_cross_domain" json:"allow_cross_domain"`
	TrustProxies     TrustProxies `mapstructure:"trust_proxies" yaml:"trust_proxies" json:"trust_proxies"`
	IPLimit          IPLimit      `mapstructure:"ip_limit" yaml:"ip_limit" json:"ip_limit"`
	TLS              TLS          `mapstructure:"tls" yaml:"tls" json:"tls"`
}

type API struct {
	Port string `mapstructure:"port" yaml:"port" json:"port"`
}
type Web struct {
	Port string `mapstructure:"port" yaml:"port" json:"port"`
}
type TrustProxies struct {
	IsOpen          int      `mapstructure:"is_open" yaml:"is_open" json:"is_open"`
	ProxyServerList []string `mapstructure:"proxy_server_list" yaml:"proxy_server_list" json:"proxy_server_list"`
}
type IPLimit struct {
	Count int `mapstructure:"count" yaml:"count" json:"count"`
	Time  int `mapstructure:"time" yaml:"time" json:"time"`
}
type TLS struct {
	Enabled bool   `mapstructure:"enabled" yaml:"enabled" json:"enabled"`
	Cert    string `mapstructure:"cert" yaml:"cert" json:"cert"`
	Key     string `mapstructure:"key" yaml:"key" json:"key"`
}
