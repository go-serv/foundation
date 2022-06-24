package net

import (
	"github.com/go-serv/service/pkg/z"
	"github.com/go-serv/service/pkg/z/ancillary/crypto"
)

type netClient struct {
	z.ClientInterface
	svc         z.NetworkServiceInterface
	insecure    bool
	blockCipher crypto.AEAD_CypherInterface
}

func (c *netClient) NetService() z.NetworkServiceInterface {
	return c.svc
}

func (c *netClient) BlockCipher() crypto.AEAD_CypherInterface {
	return c.blockCipher
}

func (c *netClient) WithBlockCipher(cipher crypto.AEAD_CypherInterface) {
	c.blockCipher = cipher
}
