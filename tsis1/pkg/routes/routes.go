package routes

import (
	"tsis1/pkg/controllers"
	"github.com/gorilla/mux"
)


var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/health-check", controllers.HealthCheck).Methods("GET")
	router.HandleFunc("/actors", controllers.Actors).Methods("GET")
	router.HandleFunc("/actors/{actorId}", controllers.ActorDetail).Methods("GET")
}