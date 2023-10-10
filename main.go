package main

import (
	"net/http"

	"github.com/go-cache/cache/database"
)

func main() {
	srv := database.NewServer()

	http.HandleFunc("/", srv.HandleGetUserById)

	http.ListenAndServe("localhost:8080", nil)
}
