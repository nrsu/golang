package dataperson
import("tsis1/pkg/models")

func GetDataPerson() []models.Actor {
	var actors []models.Actor

	var actor models.Actor
	actor.Id = 1
	actor.FirstName = "Pedro"
	actor.LastName = "Pascal"
	actors = append(actors, actor)

	actor.Id = 2
	actor.FirstName = "Margot"
	actor.LastName = "Robbie"
	actors = append(actors, actor)

	actor.Id = 3
	actor.FirstName = "Cillian"
	actor.LastName = "Murphy"
	actors = append(actors, actor)

	actor.Id = 4
	actor.FirstName = "Emily"
	actor.LastName = "Blunt"
	actors = append(actors, actor)

	actor.Id = 5
	actor.FirstName = "Rami"
	actor.LastName = "Malek"
	actors = append(actors, actor)

	actor.Id = 6
	actor.FirstName = "Jessica"
	actor.LastName = "Chastain"
	actors = append(actors, actor)

	actor.Id = 7
	actor.FirstName = "Matt"
	actor.LastName = "Damon"
	actors = append(actors, actor)

	actor.Id = 8
	actor.FirstName = "Florence"
	actor.LastName = "Pugh"
	actors = append(actors, actor)

	actor.Id = 9
	actor.FirstName = "Timothee"
	actor.LastName = "Chalamet"
	actors = append(actors, actor)

	actor.Id = 10
	actor.FirstName = "Zooey"
	actor.LastName = "Deschanel"
	actors = append(actors, actor)
	return actors
}