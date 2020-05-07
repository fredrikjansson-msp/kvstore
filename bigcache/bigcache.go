package bigcache

import (
	"errors"
	"github.com/allegro/bigcache/v2"
	"log"
	"time"
)

type KVStore struct {
	cache *bigcache.BigCache
}

func (store KVStore) Set(key string, value []byte) error {
	if err := validateKey(key); err != nil {
		return err
	}

	if err := store.cache.Set(key, value); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (store KVStore) Get(key string) (value []byte, err error) {
	if err := validateKey(key); err != nil {
		return nil, err
	}

	entry, err := store.cache.Get(key)
	if err != nil {
		if err == bigcache.ErrEntryNotFound {
			return nil, nil
		}
		return nil, err
	}
	return entry, nil
}

func (store KVStore) Destroy() error {
	return store.cache.Close()
}

type Config struct {
	Ttl time.Duration
}

func New(config Config) (KVStore, error) {
	result := KVStore{}
	var err error

	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(config.Ttl))
	if err != nil {
		return result, err
	}

	result.cache = cache
	return result, nil
}

func validateKey(key string) error {
	if key == "" {
		return errors.New("invalid key")
	}
	return nil
}
