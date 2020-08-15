package app

import "Agenda/router"

func Start() {
	router := router.NewRouter(":8080")
	router.Setup()
}