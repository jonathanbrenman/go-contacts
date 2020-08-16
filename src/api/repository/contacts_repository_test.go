package repository

import (
	"Agenda/mocks"
	"Agenda/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	contactRepository *ContactRepository
)

func init() {
	InitDatabase("sqlite3", "file::memory:?cache=shared")
	mocks.LoadMockContacts(dbClient)
	contactRepository = NewContactRepository(dbClient)
}

func TestGetAll(t *testing.T) {
	assert := assert.New(t)
	contacts, err := contactRepository.getAll()
	assert.Equal(len(contacts), 5)
	assert.Equal(nil, err)
}

func TestGetContact(t *testing.T) {
	assert := assert.New(t)
	contact, err := contactRepository.getContact(1)
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
	err := contactRepository.Add(contact)
	assert.Equal(nil, err)
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)
	contact := models.Contact{
		Name: "Contact",
		LastName: "Test",
		Company: "go-contacts",
	}
	err := contactRepository.Update(contact)
	assert.Equal(nil, err)
}

func TestSearch(t *testing.T) {
	assert := assert.New(t)
	contacts, err := contactRepository.Search("Homero")
	assert.Equal(len(contacts), 1)
	assert.Equal(contacts[0].Name, "Homero")
	assert.Equal(nil, err)
}

func TestSearchNotFound(t *testing.T) {
	assert := assert.New(t)
	contacts, err := contactRepository.Search("wakanda!")
	assert.Equal(len(contacts), 0)
	assert.Equal(nil, err)
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	err := contactRepository.Delete(1)
	assert.Equal(nil, err)
}