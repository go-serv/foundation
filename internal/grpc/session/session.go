package session

import (
	"github.com/go-serv/service/pkg/z"
	"sync"
)

type state int

const (
	New          state = iota + 1
	Active             // session context was set
	Invalidated        // something went wrong, a subject to GC
	Completed          // no more calls will be handled by the given session
	ExpiredState       // marked as expired by the session manager goroutine
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
	id z.SessionId
	state
	startedAt int64
	expireAt  int64
	encKey    []byte
	nonce     []byte
	ctx       interface{}
}

func (s *session) Id() z.SessionId {
	return s.id
}

func (s *session) State() state {
	return s.state
}

func (s *session) EncKey() []byte {
	return s.encKey
}

func (s *session) WithEncKey(key []byte) {
	s.encKey = key
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
