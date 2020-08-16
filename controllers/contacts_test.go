package controllers

import (
	"Agenda/mocks"
	"Agenda/models"
	"Agenda/repository"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

var (
	contactCtrl *contactsController
)

func init() {
	repository.InitDatabase("sqlite3", "file::memory:?cache=shared")
	mocks.LoadMockContacts(repository.GetDbConn())
	contactCtrl = NewContactController()
}

func TestGetALL(t *testing.T) {
	assert := assert.New(t)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	contactCtrl.GetAll(c)

	var response []models.Contact
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(len(response), 5)
}
