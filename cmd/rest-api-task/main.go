package main

import (
	"github.com/mojcacolnik/go-rest-api-task/internal/api"
	"github.com/mojcacolnik/go-rest-api-task/internal/db"
)

func main() {
	db.InitialMigration()
	api.InitializeRouter()
}
