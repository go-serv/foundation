package session

import z "github.com/go-serv/service/internal"

type session struct {
	id z.SessionId
}

func (s *session) Id() z.SessionId {
	return s.id
}
