package security

import "net/http"

type AccessTokenInterface interface {
}

type RoleName string

const (
	OwnerRole RoleName = "owner"
	AdminRole          = "admin"
)

type ManagerInterface interface {
	AssertRolesByRequest(r *http.Request, roles ...RoleName)
}
