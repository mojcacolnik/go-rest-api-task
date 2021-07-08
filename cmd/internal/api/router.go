package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/mojcacolnik/rest-api-task/cmd/internal/api/user"
	"log"
	"net/http"
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
