package controllers

import "github.com/gin-gonic/gin"

type pingController struct {}

type PingControler interface {
	Ping(c *gin.Context)
}

func NewPingController() *pingController{
	return &pingController{}
}

func (ctrl pingController) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":         200,
		"message": "pong",
	})
}

func (ctrl pingController) Ping(c *gin.Context) {
	ctrl.ping(c)
}