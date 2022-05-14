package meta

type MetaInterface interface {
	SessionId()
	CryptoAlgo() string
	CryptoNonce() string
}
