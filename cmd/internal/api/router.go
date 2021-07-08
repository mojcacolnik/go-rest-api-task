package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/mojcacolnik/go-rest-api-task/cmd/internal/api/user"
	"log"
	"net/http"
)

func InitializeRouter() {
	router := chi.NewRouter()

	router.Get("/api/users", user.HandleListUsers)
	router.Get("/api/users/{id}", user.HandleFindUser)
	router.Post("/api/users", user.HandleCreateUser)
	router.Put("/api/users/{id}", user.HandleUpdateUser)
	router.Delete("/api/users/{id}", user.HandleDeleteUser)

	log.Fatal(http.ListenAndServe(":9000", router))
}
