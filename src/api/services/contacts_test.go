package services

import (
	"Agenda/mocks"
	"Agenda/models"
	"Agenda/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	contactService *ContactService
)

func init() {
	repository.InitDatabase("sqlite3", "file::memory:?cache=shared")
	db := repository.GetDbConn()
	mocks.LoadMockContacts(db)
	contactService = NewContactService(db)
}

func TestGetAll(t *testing.T) {
	assert := assert.New(t)
	contacts, err := contactService.getAll()
	assert.Equal(len(contacts), 5)
	assert.Equal(nil, err)
}

func TestGetContact(t *testing.T) {
	assert := assert.New(t)
	contact, err := contactService.getContact(1)
	assert.Equal(contact.Name, "Homero")
	assert.Equal(nil, err)
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	contact := models.Contact{
		Name: "Contact",
		LastName: "Test",
		Company: "go-contacts",
	}
	err := contactService.Add(contact)
	assert.Equal(nil, err)
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)
	contact := models.Contact{
		Name: "Contact",
		LastName: "Test",
		Company: "go-contacts",
	}
	err := contactService.Update(contact)
	assert.Equal(nil, err)
}

func TestSearch(t *testing.T) {
	assert := assert.New(t)
	contacts, err := contactService.Search("Homero")
	assert.Equal(len(contacts), 1)
	assert.Equal(contacts[0].Name, "Homero")
	assert.Equal(nil, err)
}

func TestSearchNotFound(t *testing.T) {
	assert := assert.New(t)
	contacts, err := contactService.Search("wakanda!")
	assert.Equal(len(contacts), 0)
	assert.Equal(nil, err)
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	err := contactService.Delete(1)
	assert.Equal(nil, err)
}
