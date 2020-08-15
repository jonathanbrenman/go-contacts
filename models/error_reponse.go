package models

import (
	"github.com/gin-gonic/gin"
)

type httpErrorResponse struct {
	Code int
	Error string
}

func ReturnHttpError(c *gin.Context, httpError *httpErrorResponse) {
	c.JSON(httpError.Code, httpError.Error)
}

func NewHttpError(code int, error string) *httpErrorResponse {
	return &httpErrorResponse{
		Code: code,
		Error: error,
	}
}