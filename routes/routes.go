package routes

import (
	"log"
	"net/http"

	"github.com/alessandro-maciel/go-rest-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.PersonalityIndex)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
