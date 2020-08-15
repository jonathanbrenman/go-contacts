package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	pingController := NewPingController()
	assert := assert.New(t)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	pingController.Ping(c)

	var response struct{
		Code int `json:"code"`
		Message string `json:"message"`
	}
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(200, response.Code)
	assert.Equal("pong", response.Message)
}