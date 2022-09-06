package session

import (
	"github.com/go-serv/foundation/pkg/z"
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
