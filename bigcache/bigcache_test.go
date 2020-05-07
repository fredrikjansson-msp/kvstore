package bigcache_test

import (
	"KVStore/bigcache"
	"math/rand"
	"testing"
	"time"
)

func TestStoreOk(t *testing.T) {
	store := createKVStore(t)
	key := randomString(10)

	err := store.Set(key, []byte("test-val"))
	defer store.Destroy()

	if err != nil {
		t.Error(err)
	}
}

func TestGetOk(t *testing.T) {
	store := createKVStore(t)
	key := randomString(10)

	errStore := store.Set(key, []byte("test-val"))
	value, errGet := store.Get(key)

	defer store.Destroy()

	if errStore != nil {
		t.Error(errStore)
	}

	if errGet != nil {
		t.Error(errGet)
	}

	if string(value) != "test-val" {
		t.Error("value does not match")
	}
}

func TestStore_emptyKey_shouldFail(t *testing.T) {
	store := createKVStore(t)

	err := store.Set("", []byte("test-val"))
	defer store.Destroy()

	if err == nil {
		t.Error("Error expected")
	}
}

func TestGet_emptyKey_shouldFail(t *testing.T) {
	store := createKVStore(t)
	_, err := store.Get("")
	defer store.Destroy()

	if err == nil {
		t.Error("Error expected")
	}
}

func TestGet_NonExistingKey_shouldFail(t *testing.T) {
	store := createKVStore(t)
	key := randomString(10)
	value, err := store.Get(key)
	defer store.Destroy()

	if value != nil || err != nil {
		t.Error("Error expected")
	}
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func createKVStore(t *testing.T) bigcache.KVStore {
	config := bigcache.Config{
		Ttl: 30 * time.Second,
	}
	store, err := bigcache.New(config)
	if err != nil {
		t.Fatal(err)
	}
	return store
}
