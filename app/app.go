package app

import (
	"Agenda/repository"
	"Agenda/router"
)

func Start() {
	// Configure router
	router := router.NewRouter(":8080")
	router.Setup()

	// Setting up DB repository
	err := repository.InitDatabase("sqlite3","file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}
}