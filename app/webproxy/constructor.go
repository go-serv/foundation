package webproxy

import "github.com/go-serv/foundation/pkg/z"

func NewWebProxy(cfg *WebProxyConfig) *webproxy {
	wp := new(webproxy)
	wp.cfg = cfg
	return wp
}

func DefaultConfig(adminDb z.DashboardInterface) *WebProxyConfig {
	cfg := new(WebProxyConfig)
	cfg.adminDashboard = adminDb
	// For development only.
	cfg.AllowedOriginsFn = func(origin string) bool {
		return true
	}
	cfg.ApiUrlPath = "/api"
	return cfg
}
