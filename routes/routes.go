package routes

import (
	"log"
	"net/http"

	"github.com/alessandro-maciel/go-rest-api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.PersonalityIndex)

	log.Fatal(http.ListenAndServe(":8000", r))
}
