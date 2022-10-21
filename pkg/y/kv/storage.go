package kv

type resolverKey struct{}

type KeyValueStorageInterface interface {
	Get(key any) (v any, has bool)
	Set(key any, value any)
}

var (
	KeyValueStorageKey resolverKey
)
