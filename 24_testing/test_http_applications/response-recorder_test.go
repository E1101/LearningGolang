package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
)

// TestGetUsers test HTTP Get to "/users" using ResponseRecorder
func TestGetUsers(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}
