package repository

import (
	"Agenda/models"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	dbClient *gorm.DB
)

func InitDatabase(dialect string, host string) (err error) {
	dbClient, err = gorm.Open(dialect, host)
	if err != nil {
		return errors.New("[Error] Connection to database " + dialect +" has failed")
	}
	// Initialize table contacts on database
	dbClient.Debug().LogMode(true)
	dbClient.AutoMigrate(&models.Contact{})
	return err
}

func GetDbConn() *gorm.DB {
	return dbClient
}