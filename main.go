package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
)

func initializeRouter() {
	router := chi.NewRouter()

	router.Get("/api/users", handleGetUsers)
	router.Get("/api/users/{id}", handleGetUser)
	router.Post("/api/users", handleCreateUser)
	router.Put("/api/users/{id}", handleUpdateUser)
	router.Delete("/api/users/{id}", handleDeleteUser)

	log.Fatal(http.ListenAndServe(":9000", router))
}

func main() {
	initialMigration()
	initializeRouter()
}
