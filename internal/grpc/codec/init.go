package codec

import (
	"google.golang.org/grpc/encoding"
)

func init() {
	encoding.RegisterMessageWrapper("proto", MessageWrapperHandler())
}
