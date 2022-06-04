package dictionary

type DictionaryInterface interface {
	RegisterTypeHandler(typ any, handler TypeHandler)
}
