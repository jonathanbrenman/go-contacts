package repository

import (
	"Agenda/models"
	"github.com/jinzhu/gorm"
)

type contactRepository struct {
	dbClient *gorm.DB
}

type ContactRepository interface {
	GetAll()(contacts []models.Contact, err error)
	GetContact(contactId int)(contact models.Contact, err error)
	Add(contact models.Contact) error
	Delete(contactId int) error
}

func NewContactRepository(db *gorm.DB) *contactRepository {
	return &contactRepository{
		dbClient: db,
	}
}

func (repo *contactRepository) getAll() (contacts []models.Contact, err error){
	err = repo.dbClient.Find(&contacts).Error
	return contacts, err
}

func (repo *contactRepository) getContact(contactId int) (contact models.Contact, err error){
	err = repo.dbClient.Where("id = ?", contactId).Find(&contact).Error
	return contact, err
}

func (repo *contactRepository) Add(contact models.Contact) error{
	err := repo.dbClient.Save(&contact).Error
	return err
}

func (repo *contactRepository) Delete(contactId int) error{
	err := repo.dbClient.Where("id = ?", contactId).Delete(models.Contact{}).Error
	return err
}

// Interface implementation here

func (repo *contactRepository) GetAll() (contacts []models.Contact, err error) {
	return repo.getAll()
}

func (repo *contactRepository) GetContact(contactId int) (contact models.Contact, err error) {
	return repo.getContact(contactId)
}