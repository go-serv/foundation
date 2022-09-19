package dashboard

import (
	"embed"
	"net/http"
)

type dashboard struct {
	urlPath   string
	contentFs *embed.FS
}

func (d *dashboard) IsFeatureOn() bool {
	return d.contentFs != nil
}

func (d *dashboard) UrlPath() string {
	return d.urlPath
}

func (d *dashboard) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello, World!"))
}
