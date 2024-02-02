package main

import (
	"net/http"
	"tsis1/pkg/routes"
	"log"
	"github.com/gorilla/mux"
)


func main() {
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	log.Println("creating routes")
	//specify endpoints

	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}



