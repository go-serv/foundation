package net

import "github.com/go-serv/service/pkg/z"

type netService struct {
	z.ServiceInterface
	encKey []byte
}

func (b *netService) Service_OnNewSession(req z.RequestInterface) error {
	//b.MethodMustBeImplemented.Panic()
	return nil
}

func (b *netService) EncriptionKey() []byte {
	//TODO implement me
	return []byte("secret")
}

func (b *netService) WithEncriptionKey(key []byte) {
	b.encKey = key
}
