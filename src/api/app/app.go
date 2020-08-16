package app

import (
	"Agenda/mocks"
	"Agenda/repository"
	"Agenda/router"
	"fmt"
)

func Start() {
	fmt.Println("Starting go-contacts api")
	// Setting up DB repository
	err := repository.InitDatabase("sqlite3","file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}
	
	// Mock contacts
	mocks.LoadMockContacts(repository.GetDbConn())

	// Configure router
	router := router.NewRouter(":8080")
	router.Setup()
}