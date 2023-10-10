package database

import (
	"fmt"

	"github.com/go-cache/cache/handlers"
	"github.com/go-cache/cache/models"
)

func NewServer() handlers.Server {
	db := &handlers.Server{
		Users: make(map[int]*models.User), // Initialize the Users slice with a length of 100
		Cache: make(map[int]*models.User),
		DBHit: 0,
	}
	for i := 0; i < 100; i++ { // Start from 0 and go up to 99
		db.Users[i+1] = &models.User{
			Id:       i + 1,
			UserName: fmt.Sprintf("user_%d", i+1),
		}
	}
	return *db
}
