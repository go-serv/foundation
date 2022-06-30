package codec

func init() {
	dataFramePtrPool = make(dataFramePtrPoolTyp, ptrPoolSize)
}
