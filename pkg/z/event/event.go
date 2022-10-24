package event

type keyType int

const (
	NetClientNewContext keyType = iota + 1
	NetClientBeforeDial
)
