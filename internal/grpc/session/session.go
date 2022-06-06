package session

import (
	"github.com/go-serv/service/pkg/z"
)

type session struct {
	id z.SessionId
}

func (s *session) Id() z.SessionId {
	return s.id
}
