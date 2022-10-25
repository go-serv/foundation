package z

import "net/http"

type ApiKeyAwareInterface interface {
	ApiKey() []byte
	WithApiKey([]byte)
}

type ApiKeyAuthenticatorInterface interface {
	// VerifyApiKey verifies provided API key and returns a list of roles associated with it.
	VerifyApiKey(key []byte) (bool, RoleName)
	// CreateApiKey creates a new api key with associated roles.
	CreateApiKey(RoleName) []byte
	// RevokeApiKey revokes the given key.
	RevokeApiKey(key []byte)
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
