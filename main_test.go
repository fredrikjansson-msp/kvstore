package main_test

import (
	"KVStore/app"
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a app.App

func TestMain(m *testing.M) {
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestGetValueByExistingKey(t *testing.T) {
	testBody := "testdata"
	reqPost, _ := http.NewRequest("POST", "/my-key", bytes.NewBuffer([]byte(testBody)))
	reqPost.Header.Set("Content-Type", "text/plain;charset=utf-8")
	responsePost := executeRequest(reqPost)
	checkResponseCode(t, http.StatusCreated, responsePost.Code)

	reqGet, _ := http.NewRequest("GET", "/my-key", nil)
	responseGet := executeRequest(reqGet)
	checkResponseCode(t, http.StatusOK, responseGet.Code)

	if body := responseGet.Body.String(); body != testBody {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestPostWithInvalidEncoding(t *testing.T) {
	testBody := "testdata"
	reqPost, _ := http.NewRequest("POST", "/my-key", bytes.NewBuffer([]byte(testBody)))
	reqPost.Header.Set("Content-Type", "application/json")
	responsePost := executeRequest(reqPost)
	checkResponseCode(t, http.StatusMethodNotAllowed, responsePost.Code)
}

func TestGetValueByNonExistingKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "/invalid-key", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
