package x

type DictionaryAwareInterface interface {
	Dictionary() DictionaryInterface
	WithDictionary(DictionaryInterface)
}

type DictionaryInterface interface {
	Import(target DictionaryInterface) error
	Export(target DictionaryInterface) error
}
