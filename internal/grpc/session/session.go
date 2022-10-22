package session

import (
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/pkg/z/ancillary/crypto"
	"sync"
)

const (
	New         z.SessionState = iota + 1
	Active                     // session context was set.
	Invalidated                // something went wrong, a subject to GC.
	Closed                     // if closed, no more calls will be handled by the given session.
)

var sessionMap = sync.Map{}

func FindById(key z.SessionId) *session {
	if v, ok := sessionMap.Load(key); ok {
		return v.(*session)
	} else {
		return nil
	}
}

type session struct {
	id          z.SessionId
	state       z.SessionState
	startedAt   int64
	expireAt    int64
	blockCipher crypto.AEAD_CipherInterface
	nonce       []byte
	ctx         interface{}
}

func (s *session) Id() z.SessionId {
	return s.id
}

func (s *session) State() z.SessionState {
	return s.state
}

func (s *session) IsValid() bool {
	if s.state == New || s.state == Active {
		return true
	} else {
		return false
	}
}

func (s *session) BlockCipher() crypto.AEAD_CipherInterface {
	return s.blockCipher
}

func (s *session) WithBlockCipher(cipher crypto.AEAD_CipherInterface) {
	s.blockCipher = cipher
}

func (s *session) Nonce() []byte {
	return s.nonce
}

func (s *session) WithNonce(nonce []byte) {
	s.nonce = nonce
}

func (s *session) Context() interface{} {
	return s.ctx
}

func (s *session) WithContext(ctx interface{}) {
	s.ctx = ctx
	s.state = Active
}

func (s *session) Close() {
	s.state = Closed
}
