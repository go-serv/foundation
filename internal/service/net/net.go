package net

import (
	i "github.com/go-serv/service/internal"
)

type netService struct {
	i.ServiceInterface
	encKey []byte
}

func (b *netService) Service_OnNewSession(req i.RequestInterface) error {
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

func (b *netService) Service_InfoNewSession(methodName string) int32 {
	return 0
	//mDesc := b.service.sd.FindMethodDescriptorByName(methodName)
	//if mDesc == nil {
	//	return 0
	//} else {
	//	v, has := mDesc.Get(go_serv.E_NetNewSession)
	//	if !has {
	//		return 0
	//	} else {
	//		return v.(int32)
	//	}
	//}
}

func (b *netService) Service_InfoMsgEncryption(methodName string) bool {
	return true
	//mDesc := b.service.sd.FindMethodDescriptorByName(methodName)
	//if v, has := mDesc.Get(go_serv.E_MNetMsgEnc); has {
	//	return v.(bool)
	//} else {
	//	if v, has := b.service.sd.Get(go_serv.E_NetMsgEnc); has {
	//		return v.(bool)
	//	} else {
	//		return false
	//	}
	//}
}
