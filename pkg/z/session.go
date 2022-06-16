package z

type (
	OnGCFn       func()
	SessionState int
)

type SessionInterface interface {
	Id() SessionId
	State() SessionState
	Nonce() []byte
	WithNonce([]byte)
	EncKey() []byte
	WithEncKey([]byte)
	Context() interface{}
	WithContext(interface{})
}
