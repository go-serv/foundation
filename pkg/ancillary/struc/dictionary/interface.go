package dictionary

type AwareInterface interface {
	Dictionary() DictionaryInterface
	WithDictionary(DictionaryInterface)
}

type DictionaryInterface interface {
	Import(target DictionaryInterface) error
	Export(target DictionaryInterface) error
}
