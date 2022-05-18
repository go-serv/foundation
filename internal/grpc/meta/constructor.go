package meta

import "google.golang.org/grpc/metadata"

func NewMeta(md metadata.MD) MetaInterface {
	m := new(meta)
	m.data = md
	return m
}
