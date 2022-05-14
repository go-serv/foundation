package meta

import "google.golang.org/grpc/metadata"

type meta struct {
	header      metadata.MD
	cryptoAlgo  string
	cryptoNonce string
}

type serverMeta struct {
	meta
	trailer metadata.MD
}

type clientMeta struct {
	meta
}
