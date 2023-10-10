package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-cache/cache/database"
	"github.com/go-cache/cache/models"
)

func TestGetUserHandler(t *testing.T) {
	s := database.NewServer()
	ts := httptest.NewServer(http.HandlerFunc(s.HandleGetUserById))
	nReq := 1000

	for i := 0; i < nReq; i++ {
		id := i%100 + 1
		url := fmt.Sprintf("%s/?id=%d", ts.URL, id)
		resp, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}
		user := &models.User{}
		if err := json.NewDecoder(resp.Body).Decode(user); err != nil {
			t.Error(err)
		}
		fmt.Println(user)
	}
	fmt.Println("total DB hit: ", s.DBHit)
}
