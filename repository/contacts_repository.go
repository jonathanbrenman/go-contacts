package repository

import (
	"Agenda/models"
	"github.com/jinzhu/gorm"
)

type repository struct {
	dbClient *gorm.DB
}

type ContactRepository interface {
	GetAll()(contacts []models.Contact, err error)
	GetContact(contactId int)(contact models.Contact, err error)
	Add(contact models.Contact) error
	Delete(contactId int) error
}

func NewContactRepository(db *gorm.DB) *repository {
	return &repository{
		dbClient: db,
	}
}

func (repo *repository) getAll() (contacts []models.Contact, err error){
	err = repo.dbClient.Find(&contacts).Error
	return contacts, err
}

func (repo *repository) getContact(contactId int) (contact models.Contact, err error){
	err = repo.dbClient.Where("id = ?", contactId).Find(&contact).Error
	return contact, err
}

func (repo *repository) Add(contact models.Contact) error{
	err := repo.dbClient.Save(&contact).Error
	return err
}

func (repo *repository) Delete(contactId int) error{
	err := repo.dbClient.Where("id = ?", contactId).Delete(models.Contact{}).Error
	return err
}

// Interface implementation here

func (repo *repository) GetAll() (contacts []models.Contact, err error) {
	return repo.getAll()
}

func (repo *repository) GetContact(contactId int) (contact models.Contact, err error) {
	return repo.getContact(contactId)
}