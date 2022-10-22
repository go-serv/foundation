package z

import "github.com/go-serv/foundation/pkg/z/ancillary/crypto"

type (
	OnGCFn       func()
	SessionState int
)

type SessionInterface interface {
	Id() SessionId
	State() SessionState
	IsValid() bool
	Nonce() []byte
	WithNonce([]byte)
	BlockCipher() crypto.AEAD_CipherInterface
	WithBlockCipher(crypto.AEAD_CipherInterface)
	Context() interface{}
	WithContext(interface{})
	Close()
}
