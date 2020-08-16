package repository

import (
	"errors"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitDatabase(t *testing.T) {
	err := InitDatabase("sqlite3", "file::memory:?cache=shared")
	assert := assert.New(t)
	assert.Equal(nil, err)
}

func TestInitDatabaseFail(t *testing.T) {
	err := InitDatabase("sqlite3", "./")
	assert := assert.New(t)
	assert.Equal(errors.New("[Error] Connection to database sqlite3 has failed"), err)
}

func TestGetDbConn(t *testing.T) {
	dbClient := GetDbConn()
	assert := assert.New(t)
	assert.NotEqual(nil, dbClient)
}