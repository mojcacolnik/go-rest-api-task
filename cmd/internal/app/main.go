package main

import (
	"github.com/mojcacolnik/go-rest-api-task/cmd/internal/api"
	"github.com/mojcacolnik/go-rest-api-task/cmd/internal/db"
)

func main() {
	db.InitialMigration()
	api.InitializeRouter()
}
