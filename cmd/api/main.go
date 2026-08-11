package main

import (
	"fmt"
	
	"github.com/go-chi/chi/v5"
	
	"drive-chi/internal/users"
	"drive-chi/pkg/database"
)

func main () {

	r := chi.NewRouter()

	db, err := database.NewConnection()
	if err != nil {
		fmt.Print("Error connecting to the database")
	}

	users.SetRoutes(r, db)
}