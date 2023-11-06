package main

import (
	"github.com/alessandro-maciel/go-rest-api/models"
	"github.com/alessandro-maciel/go-rest-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Nome 1", History: "História 1"},
		{Id: 2, Name: "Nome 2", History: "História 2"},
	}

	routes.HandleRequest()
}
