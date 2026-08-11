package main

import (
	"log"
		"net/http"
	
	"github.com/go-chi/chi/v5"
	
	"drive-chi/internal/users"
	"drive-chi/pkg/database"
)

func main () {

	r := chi.NewRouter()

	db, err := database.NewConnection()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	users.SetRoutes(r, db)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}