package router

import "Agenda/controllers"

func (r routerCtx) routes() {
	// Define Controllers
	pingController := controllers.NewPingController()

	// Define Routes Here
	r.router.GET("/ping", pingController.Ping)
}
