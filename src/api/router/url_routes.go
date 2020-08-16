package router

import (
	"Agenda/controllers"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/files" // swagger embed files
	_ "Agenda/docs"
)

func (r routerCtx) routes() {
	// Define Controllers
	pingController := controllers.NewPingController()
	contactController := controllers.NewContactController()
	
	// Define Routes Here
	r.router.GET("/ping", pingController.Ping)
	r.router.GET("/contacts/search", contactController.Search)
	r.router.GET("/contacts/id/:id", contactController.GetContact)
	r.router.GET("/contacts", contactController.GetAll)
	r.router.POST("/contacts/update", contactController.Update)
	r.router.POST("/contacts", contactController.Add)
	r.router.DELETE("/contacts/id/:id", contactController.Delete)
	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
