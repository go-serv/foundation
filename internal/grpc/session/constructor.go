package session

import (
	"github.com/go-serv/service/pkg/z"
	"time"
)

// NewSession creates a new session with the given lifetime in seconds and stores it in the session map by its ID
// for future use.
func NewSession(lifetime uint16) *session {
	s := new(session)
	s.id = z.SessionId(z.UniqueId(0).Generate())
	s.state = New
	s.startedAt = time.Now().Unix()
	s.expireAt = s.startedAt + int64(lifetime)
	sessionMap.Store(s.id, s)
	return s
}

func NewSecureSession(lifetime uint16, nonce []byte, encKey []byte) *session {
	s := NewSession(lifetime)
	s.nonce = nonce
	s.encKey = encKey
	return s
}
