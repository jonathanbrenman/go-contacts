package models

import "github.com/gin-gonic/gin"

type HttpErrorResponse struct {
	Code int
	Error string
}

func ReturnHttpError(c *gin.Context, err HttpErrorResponse) {
	c.JSON(err.Code, err)
}