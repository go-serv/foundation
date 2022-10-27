package dictionary

import "github.com/go-serv/foundation/pkg/z"

type BaseInterface interface {
	GetSessionId() z.SessionId
	SetSessionId(z.SessionId)
}

type NetRequestInterface interface {
	GetApiKey() []byte
	SetApiKey([]byte)
}
