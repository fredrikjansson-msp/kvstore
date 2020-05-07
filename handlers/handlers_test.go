package handlers_test

import (
	"KVStore/handlers"
	"bytes"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyKVStore struct{}

var setMock func(key string, value []byte) error
var getMock func(key string, val []byte) (value []byte, err error)
var destroyMock func() error

func (s *SpyKVStore) Set(key string, value []byte) error {
	return setMock(key, value)
}

func (s *SpyKVStore) Get(key string) (value []byte, err error) {
	return getMock(key, value)
}

func (s *SpyKVStore) Destroy() error {
	return destroyMock()
}

func TestPost(t *testing.T) {
	type test struct {
		err        error
		httpStatus int
	}

	tests := []test{
		{err: nil, httpStatus: 201},
		{err: errors.New("error in cache"), httpStatus: 500},
	}

	spyKVStore := &SpyKVStore{}
	handler := handlers.StoreHandler{
		Store: spyKVStore,
	}

	for _, tc := range tests {
		setMock = func(key string, value []byte) error {
			return tc.err
		}

		var str = []byte("test")
		req, err := http.NewRequest("POST", "/key", bytes.NewBuffer(str))
		req.Header.Set("Content-Type", "application/text")

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/{key}", handler.Post)
		router.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != tc.httpStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tc.httpStatus)
		}
	}
}

func TestGet(t *testing.T) {
	type test struct {
		entry      []byte
		err        error
		httpStatus int
	}

	tests := []test{
		{entry: []byte("test"), err: nil, httpStatus: 200},
		{entry: nil, err: nil, httpStatus: 404},
		{entry: nil, err: errors.New("error in cache"), httpStatus: 500},
	}

	spyKVStore := &SpyKVStore{}
	handler := handlers.StoreHandler{
		Store: spyKVStore,
	}

	for _, tc := range tests {
		getMock = func(key string, value []byte) ([]byte, error) {
			return tc.entry, tc.err
		}

		req, err := http.NewRequest("GET", "/key", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/{key}", handler.Get)
		router.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != tc.httpStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tc.httpStatus)
		}
	}
}

func TestNotImplemented(t *testing.T) {
	spyKVStore := &SpyKVStore{}
	req, err := http.NewRequest("DELETE", "/key", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := handlers.StoreHandler{
		Store: spyKVStore,
	}

	handler.NotImplemented(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotImplemented)
	}
}
