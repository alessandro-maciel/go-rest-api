package main

import (
	"github.com/alessandro-maciel/go-rest-api/database"
	"github.com/alessandro-maciel/go-rest-api/routes"
)

func main() {
	database.ConnectDatabase()

	routes.HandleRequest()
}
