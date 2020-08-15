package services

import (
	"Agenda/models"
	"Agenda/repository"
	"github.com/jinzhu/gorm"
)

type ContactService struct {
	Repository repository.ContactRepository
}

type ContactServiceInterface interface {
	GetAll()(contacts []models.Contact, err error)
	GetContact(contactId int)(contact models.Contact, err error)
	Add(contact models.Contact) error
	Delete(contactId int) error
	Update(contact models.Contact) error
}

func NewContactService(db *gorm.DB) *ContactService{
	repository.NewContactRepository(db)
	return &ContactService{}
}

func (service *ContactService) getAll() ([]models.Contact, error){
	return service.Repository.GetAll()
}

func (service *ContactService) getContact(contactId int) (models.Contact, error){
	return service.Repository.GetContact(contactId)
}

func (service *ContactService) add(contact models.Contact) error{
	return service.Repository.Add(contact)
}

func (service *ContactService) delete(contactId int) error{
	return service.Repository.Delete(contactId)
}

func (service *ContactService) update(contact models.Contact) error{
	return service.Repository.Update(contact)
}

// Interface implementation here

func (service *ContactService) GetAll() ([]models.Contact, error) {
	return service.getAll()
}

func (service *ContactService) GetContact(contactId int) (models.Contact, error) {
	return service.getContact(contactId)
}

func (service *ContactService) Add(contact models.Contact) error {
	return service.add(contact)
}

func (service *ContactService) Delete(contactId int) error {
	return service.delete(contactId)
}

func (service *ContactService) Update(contact models.Contact) error {
	return service.update(contact)
}