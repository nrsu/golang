package models

type AllActors struct {
	Actors []Actor `json:"actors"`
}

type Actor struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}