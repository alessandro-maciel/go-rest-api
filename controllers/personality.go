package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/alessandro-maciel/go-rest-api/models"
)

func PersonalityIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}
