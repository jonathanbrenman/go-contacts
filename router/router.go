package router

import (
	"Agenda/middlewares"
	"github.com/gin-gonic/gin"
)

type routerCtx struct {
	router *gin.Engine
	port string
}

type Router interface {
	Setup()
}

func NewRouter(port string) *routerCtx {
	return &routerCtx{
		router: gin.Default(),
		port: port,
	}
}

// Router Setup
func (r routerCtx) configure() {
	r.router.Use(middlewares.CORSMiddleware())
	r.routes()
	r.router.Run(r.port)
}

func (r routerCtx) Setup() {
	r.configure()
}

