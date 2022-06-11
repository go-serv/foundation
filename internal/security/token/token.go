package token

import "github.com/go-serv/service/pkg/z"

type token struct {
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
