package http

import "net/http"

type HeaderName string

const (
	ContentTypeHeader HeaderName = "content-type"
)

func WithContentType(res http.ResponseWriter, mimeType string) {
	res.Header().Set(string(ContentTypeHeader), mimeType)
}
