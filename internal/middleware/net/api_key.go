package net

import (
	"github.com/go-serv/foundation/internal/autogen/foundation"
	"github.com/go-serv/foundation/internal/grpc/meta/net"
	"github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/pkg/z"
)

func verifyApiKey(ctx z.NetServerContextInterface, req z.RequestInterface) bool {
	var (
		verified       bool
		err            error
		keyRole        z.RoleName
		authenticator  z.ApiKeyAuthenticatorInterface
		authType       foundation.AuthType
		rawKey, apiKey []byte
	)
	if v, has := req.ServiceReflection().Get(foundation.E_AuthType); !has {
		return true
	} else {
		authType = v.(foundation.AuthType)
		if authType != foundation.AuthType_ApiKey {
			return true
		}
	}

	rawKey = req.Meta().Dictionary().(*net.HttpDictionary).ApiKey
	if ctx.NetworkService().IsTlsEnabled() {
		apiKey = rawKey
	} else {
		cipher := ctx.Session().BlockCipher()
		if apiKey, err = cipher.Decrypt(rawKey, nil); err != nil {
			return false
		}
	}

	authResolverKey := (*z.ApiKeyAuthenticatorInterface)(nil)
	authResolverValue, _ := runtime.Runtime().Resolve(authResolverKey)
	authenticator = authResolverValue.(z.ApiKeyAuthenticatorInterface)
	if verified, keyRole = authenticator.VerifyApiKey(apiKey); !verified {
		return false
	}

	if v, has := req.ServiceReflection().Get(foundation.E_RequiresRole); has {
		reqRoles := v.([]string)
		for _, reqRole := range reqRoles {
			if string(keyRole) == reqRole {
				return true
			}
		}
		return false // the given api key does not have required role.
	}
	return true
}
