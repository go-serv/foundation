package dictionary

type DictionaryInterface interface {
	RegisterTypeHandler(typ any, handler TypeHandler)
	Hydrate(target interface{}) error
	Dehydrate(target interface{}) error
	Context() interface{}
	WithContext(interface{})
}
