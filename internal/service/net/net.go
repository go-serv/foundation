package net

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
	"github.com/go-serv/service/internal/service"
)

type networkService struct {
	service.BaseService
	encKey []byte
}

func (b *networkService) Service_OnNewSession(req i.RequestInterface) error {
	b.MethodMustBeImplemented.Panic()
	return nil
}

func (b *networkService) Service_EncriptionKey() []byte {
	//TODO implement me
	return []byte("secret")
}

func (b *networkService) Service_WithEncriptionKey(key []byte) {
	b.encKey = key
}

func (b *networkService) Service_InfoNewSession(methodName string) int32 {
	mDesc := b.BaseService.Sd.FindMethodDescriptorByName(methodName)
	if mDesc == nil {
		return 0
	} else {
		v, has := mDesc.Get(go_serv.E_NetNewSession)
		if !has {
			return 0
		} else {
			return v.(int32)
		}
	}
}

func (b *networkService) Service_InfoMsgEncryption(methodName string) bool {
	mDesc := b.BaseService.Sd.FindMethodDescriptorByName(methodName)
	if v, has := mDesc.Get(go_serv.E_MNetMsgEnc); has {
		return v.(bool)
	} else {
		if v, has := b.BaseService.Sd.Get(go_serv.E_NetMsgEnc); has {
			return v.(bool)
		} else {
			return false
		}
	}
}
