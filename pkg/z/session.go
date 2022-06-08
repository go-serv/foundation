package z

type (
	OnGCFn func()
)
type SessionInterface interface {
	Id() SessionId
	Nonce() []byte
	WithNonce([]byte)
	EncKey() []byte
	WithEncKey([]byte)
	Context() interface{}
	WithContext(interface{})
}
