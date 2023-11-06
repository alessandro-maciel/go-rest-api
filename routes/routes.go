package routes

import (
	"log"
	"net/http"

	"github.com/alessandro-maciel/go-rest-api/controllers"
	"github.com/alessandro-maciel/go-rest-api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.PersonalityIndex).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.PersonalityStore).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.PersonalityShow).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.PersonalityUpdate).Methods("Put")
	r.HandleFunc("/api/personalities/{id}", controllers.PersonalityDelete).Methods("Delete")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
