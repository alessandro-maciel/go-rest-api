package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alessandro-maciel/go-rest-api/models"
	"github.com/gorilla/mux"
)

func PersonalityIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}

func PersonalityShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personality := range models.Personalities {
		if strconv.Itoa(personality.Id) == id {
			json.NewEncoder(w).Encode(personality)
		}
	}
}
