package dashboard

import (
	"embed"
	http_helper "github.com/go-serv/foundation/internal/ancillary/http"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type dashboard struct {
	urlPath   string
	contentFs *embed.FS
}

func (d *dashboard) PathPrefix() string {
	return d.urlPath
}

func (d *dashboard) WithPathPrefix(url string) {
	d.urlPath = strings.TrimRight(url, "/") + "/"
}

func (d *dashboard) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var (
		mimeType string
		filename string
	)
	path := strings.TrimLeft(req.URL.Path, d.urlPath)
	if len(path) > 0 {
		ext := filepath.Ext(path)
		if len(ext) > 0 {
			mimeType = mime.TypeByExtension(ext)
			http_helper.WithContentType(res, mimeType)
		}
		filename = path
	} else {
		//todo ManagerService, token check
		filename = "dashboard.html"
	}
	if file, err := d.contentFs.ReadFile("build" + string(os.PathSeparator) + filename); err == nil {
		res.Write(file)
	} else {
		res.WriteHeader(404)
	}
}
