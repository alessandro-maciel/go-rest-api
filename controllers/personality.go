package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/alessandro-maciel/go-rest-api/database"
	"github.com/alessandro-maciel/go-rest-api/models"
	"github.com/gorilla/mux"
)

func PersonalityIndex(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personality
	database.Db.Find(&personalities)

	json.NewEncoder(w).Encode(personalities)
}

func PersonalityShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality
	database.Db.First(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func PersonalityStore(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality

	json.NewDecoder(r.Body).Decode(&personality)
	database.Db.Create(&personality)

	json.NewEncoder(w).Encode(personality)
}
