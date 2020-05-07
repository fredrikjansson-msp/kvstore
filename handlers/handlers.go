package handlers

import (
	"KVStore/store"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type StoreHandler struct {
	Store store.KVStore
}

func (h *StoreHandler) Post(w http.ResponseWriter, r *http.Request) {
	routeVars := mux.Vars(r)

	if key, ok := routeVars["key"]; ok {
		entry, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := h.Store.Set(key, entry); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("value with key \"%s\" stored in cache.", key)
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *StoreHandler) Get(w http.ResponseWriter, r *http.Request) {
	routeVars := mux.Vars(r)

	if val, ok := routeVars["key"]; ok {
		if entry, err := h.Store.Get(val); err == nil && entry != nil {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/text")
			w.Write(entry)
		} else if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *StoreHandler) NotImplemented(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusNotImplemented)
}
