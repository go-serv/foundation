package dictionary

func NewNetRequestDictionary() *NetRequestDictionary {
	netReq := new(NetRequestDictionary)
	netReq.BaseDictionary = &BaseDictionary{}
	return netReq
}

func NewNetResponseDictionary() *NetResponseDictionary {
	netRes := new(NetResponseDictionary)
	netRes.BaseDictionary = &BaseDictionary{}
	return netRes
}
