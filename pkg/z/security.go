package z

import "net/http"

type ApiKeyAwareInterface interface {
	ApiKey() []byte
	WithApiKey([]byte)
}

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
