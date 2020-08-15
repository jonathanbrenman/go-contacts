package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func NewRepositoryDB(dialect string, host string) (*repository, error) {
	db, err := gorm.Open(dialect, host)
	if err != nil {
		return &repository{}, errors.New("[Error] Connection to database " + dialect +" has failed")
	}
	defer db.Close()

	return &repository{
		dbClient: db,
	}, nil
}
