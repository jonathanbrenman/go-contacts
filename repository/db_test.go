package repository

import (
	"errors"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRepositoryDB(t *testing.T) {
	db, _ := NewRepositoryDB("sqlite3", "file::memory:?cache=shared")
	assert := assert.New(t)

	assert.NotEqual(nil, db)
}

func TestNewRepositoryDBFail(t *testing.T) {
	_, err := NewRepositoryDB("sqlite3", "./")
	assert := assert.New(t)

	assert.Equal(errors.New("[Error] Connection to database sqlite3 has failed"), err)
}
