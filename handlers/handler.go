package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-cache/cache/models"
)

type Server models.Server

func (srv *Server) TryCache(id int) (*models.User, bool) {
	user, ok := srv.Cache[id]

	return user, ok
}

func (srv *Server) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	fmt.Println(idStr)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	user, ok := srv.TryCache(id)

	if ok {
		json.NewEncoder(w).Encode(user)
		return
	}

	user, ok = srv.Users[id]

	if !ok {
		json.NewEncoder(w).Encode("user not found!")
		return
	}
	srv.Cache[id] = user
	srv.DBHit++
	json.NewEncoder(w).Encode(user)
}
