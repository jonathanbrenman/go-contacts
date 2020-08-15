package repository

import "github.com/jinzhu/gorm"

type repository struct {
	dbClient *gorm.DB
}

type Repository interface {

}

