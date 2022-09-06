package token

import (
	"github.com/go-serv/foundation/pkg/z"
)

type tokenTyp int

const (
	RbacToken tokenTyp = iota + 1
)

func (t tokenTyp) String() string {
	return [...]string{"rbac"}[t]
}

type token struct {
	typ       tokenTyp   `json:"type"`
	issuedAt  int64      `json:"issued_at"`
	notBefore int64      `json:"not_before"`
	expireAt  int64      `json:"expire_at"`
	issuer    string     `json:"issuer"` // Host address of an authorization server issued the given token
	tenantId  z.TenantId `json:"tenant_id"`
}

type rbacToken struct {
	token
	roles []z.RoleName `json:"roles"`
}
