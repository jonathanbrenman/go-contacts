package router

import "Agenda/controllers"

func (r routerCtx) routes() {
	// Define Controllers
	pingController := controllers.NewPingController()
	contactController := controllers.NewContactController()

	// Define Routes Here
	r.router.GET("/ping", pingController.Ping)
	r.router.GET("/contacts/:id", contactController.GetContact)
	r.router.GET("/contacts", contactController.GetAll)
	r.router.PUT("/contacts/:id", contactController.Update)
	r.router.POST("/contacts", contactController.Add)
	r.router.DELETE("/contacts/:id", contactController.Delete)
}
