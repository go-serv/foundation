package meta

import "google.golang.org/grpc/metadata"

type meta struct {
	data        metadata.MD
	cryptoAlgo  string
	cryptoNonce string
}

type serverMeta struct {
	meta
}

type clientMeta struct {
	meta
}
