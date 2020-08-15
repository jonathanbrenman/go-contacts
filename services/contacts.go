package services

import (
	"Agenda/models"
	"Agenda/repository"
	"github.com/jinzhu/gorm"
)

type ContactService struct {
	contactRepository repository.ContactRepository
}

type ContactServiceInterface interface {
	GetAll()(contacts []models.Contact, err error)
	GetContact(contactId int)(contact models.Contact, err error)
	Add(contact models.Contact) (err error)
	Delete(contactId int) (err error)
	Update(contact models.Contact) (err error)
}

func (service *ContactService) getAll() (contacts []models.Contact, err error) {
	return service.contactRepository.GetAll()
}

func (service *ContactService) getContact(contactId int) (contact models.Contact, err error) {
	return service.contactRepository.GetContact(contactId)
}

func (service *ContactService) add(contact models.Contact) (err error) {
	return service.contactRepository.Add(contact)
}

func (service *ContactService) delete(contactId int) (err error) {
	return service.contactRepository.Delete(contactId)
}

func (service *ContactService) update(contact models.Contact) (err error) {
	return service.contactRepository.Update(contact)
}

// Interface implementation here

func (service *ContactService) GetAll() (contacts []models.Contact, err error) {
	return service.getAll()
}

func (service *ContactService) GetContact(contactId int) (contact models.Contact, err error) {
	return service.getContact(contactId)
}

func (service *ContactService) Add(contact models.Contact) (err error) {
	return service.add(contact)
}

func (service *ContactService) Delete(contactId int) (err error) {
	return service.delete(contactId)
}

func (service *ContactService) Update(contact models.Contact) (err error) {
	return service.update(contact)
}

// Factory Functions

func NewContactService(db *gorm.DB) *ContactService{
	contactRepository := repository.NewContactRepository(db)
	return &ContactService{
		contactRepository: *contactRepository,
	}
}