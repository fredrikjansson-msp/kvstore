package store

type KVStore interface {
	Set(key string, value []byte) error
	Get(key string) (value []byte, err error)
	Destroy() error
}
