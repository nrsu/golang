package controllers
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tsis1/pkg/models"
	"tsis1/pkg/dataperson"
	"github.com/gorilla/mux"
	"strconv"
)
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func Actors(w http.ResponseWriter, r *http.Request) {
	log.Println("entering persons end point")
	var response models.AllActors
	persons := dataperson.GetDataPerson()

	response.Actors = persons

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}
func ActorDetail(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	actorId := vars["actorId"]
	actors := dataperson.GetDataPerson()
	var foundActor models.Actor
	checking := foundActor
	for _, actor := range actors {
		if strconv.Itoa(actor.Id) == actorId {
			foundActor = actor
			break
		}
	}
	if checking==foundActor{
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(foundActor)
	if err != nil {
		return
	}

	w.Write(jsonResponse)

}