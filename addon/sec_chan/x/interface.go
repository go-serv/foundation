package x

import (
	"github.com/go-serv/foundation/pkg/z/ancillary/crypto"
)

type DataFrameInterface interface {
	Parse(wire []byte) error
	Compose() ([]byte, error)
	WithBlockCipher(cipher crypto.AEAD_CipherInterface)
	Decrypt() error
}
