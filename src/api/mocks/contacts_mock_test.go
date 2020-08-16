package mocks

import (
	"Agenda/repository"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadMockContacts(t *testing.T) {
	err := repository.InitDatabase("sqlite3", "file::memory:?cache=shared")
	assert := assert.New(t)
	assert.Equal(nil, err)

	dbClient := repository.GetDbConn()
	contacts := LoadMockContacts(dbClient)
	assert.Equal(len(contacts), 5)
	for _, contact := range contacts {
		assert.Equal(fmt.Sprintf("%T\n",contact.ID), string("uint\n"))
		assert.Equal(fmt.Sprintf("%T\n",contact.Name), string("string\n"))
		assert.Equal(fmt.Sprintf("%T\n",contact.LastName), string("string\n"))
		assert.Equal(fmt.Sprintf("%T\n",contact.Address), string("string\n"))
		assert.Equal(fmt.Sprintf("%T\n",contact.Thumb), string("string\n"))
		assert.Equal(fmt.Sprintf("%T\n",contact.Phone), string("string\n"))
		assert.Equal(fmt.Sprintf("%T\n",contact.Mail), string("string\n"))
		assert.Equal(fmt.Sprintf("%T\n",contact.Company), string("string\n"))
	}
}

func TestGetContacts(t *testing.T) {
	assert := assert.New(t)
	contacts := getContacts()
	assert.Equal(len(contacts), 5)
	assert.Equal("Homero", contacts[0].Name)
	assert.Equal("Marsh", contacts[1].Name)
	assert.Equal("Bartolomeo", contacts[2].Name)
	assert.Equal("Lisa", contacts[3].Name)
	assert.Equal("Magie", contacts[4].Name)
}