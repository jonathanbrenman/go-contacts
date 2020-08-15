package models

import "github.com/jinzhu/gorm"

type Contact struct {
	gorm.Model
	Name string `db:"name" json:"name"`
	LastName string `db:"last_name" json:"last_name"`
	Company string `db:"company" json:"company"`
	Phone string `db:"phone" json:"phone"`
	Mail string `db:"mail" json:"mail"`
	Address string `db:"address" json:"address"`
	ContactImage
}

type ContactImage struct {
	Thumb string `db:"thumb" json:"thumb"`
}