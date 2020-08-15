package services

import (
	"Agenda/models"
	"Agenda/repository"
	"github.com/jinzhu/gorm"
)

type contactService struct {
	Repository repository.ContactRepository
}

type ContactService interface {
	GetAll()(contacts []models.Contact, err error)
	GetContact(contactId int)(contact models.Contact, err error)
	Add(contact models.Contact) error
	Delete(contactId int) error
	Update(contactId int, contact models.Contact) error
}

func NewContactService(db *gorm.DB) *contactService{
	repository.NewContactRepository(db)
	return &contactService{}
}

func (service *contactService) getAll() ([]models.Contact, error){
	return service.Repository.GetAll()
}

func (service *contactService) getContact(contactId int) (models.Contact, error){
	return service.Repository.GetContact(contactId)
}

func (service *contactService) add(contact models.Contact) error{
	return service.Repository.Add(contact)
}

func (service *contactService) delete(contactId int) error{
	return service.Repository.Delete(contactId)
}

func (service *contactService) update(contactId int, contact models.Contact) error{
	return service.Repository.Update(contactId, contact)
}

// Interface implementation here

func (service *contactService) GetAll() ([]models.Contact, error) {
	return service.getAll()
}

func (service *contactService) GetContact(contactId int) (models.Contact, error) {
	return service.getContact(contactId)
}

func (service *contactService) Delete(contactId int) error {
	return service.delete(contactId)
}

func (service *contactService) Update(contactId int, contact models.Contact) error {
	return service.update(contactId, contact)
}