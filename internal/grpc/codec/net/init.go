package codec

import "google.golang.org/grpc/encoding"

func init() {
	encoding.RegisterCodec(codec{})
}
