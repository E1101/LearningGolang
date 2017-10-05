package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
)

// TestGetUsersWithServer test HTTP Get to "/users" using Server
func TestGetUsersWithServer(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")

	server := httptest.NewServer(r)
	defer server.Close()

	usersURL     := fmt.Sprintf("%s/users", server.URL)
	request, err := http.NewRequest("GET", usersURL, nil)
	res, err     := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", res.StatusCode)
	}
}
