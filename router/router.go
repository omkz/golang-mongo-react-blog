package router

import (
	"github.com/gorilla/mux"
	"github.com/omkz/golang-mongo-rest/controllers"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/posts", controllers.GetAllPosts).Methods("GET", "OPTIONS")

	return router
}