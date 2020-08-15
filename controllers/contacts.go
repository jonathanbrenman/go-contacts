package controllers

import (
	"github.com/gin-gonic/gin"
)

type contactsController struct {}

type ContactsController interface {
	GetAll(c *gin.Context)
	GetContact(c *gin.Context)
	Add(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

func NewContactController() *contactsController{
	return &contactsController{}
}

func (ctrl contactsController) getAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":         200,
		"message": "pong",
	})
}

func (ctrl contactsController) getContact(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":         200,
		"message": "pong",
	})
}

func (ctrl contactsController) add(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":         200,
		"message": "pong",
	})
}

func (ctrl contactsController) delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":         200,
		"message": "pong",
	})
}

func (ctrl contactsController) update(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":         200,
		"message": "pong",
	})
}

// Interface implementation here

func (ctrl contactsController) GetAll(c *gin.Context) {
	ctrl.getAll(c)
}

func (ctrl contactsController) GetContact(c *gin.Context) {
	ctrl.getContact(c)
}

func (ctrl contactsController) Add(c *gin.Context) {
	ctrl.add(c)
}

func (ctrl contactsController) Delete(c *gin.Context) {
	ctrl.delete(c)
}

func (ctrl contactsController) Update(c *gin.Context) {
	ctrl.update(c)
}