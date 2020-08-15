package repository

import (
	"Agenda/models"
	"github.com/jinzhu/gorm"
)

<<<<<<< HEAD
type contactRepository struct {
=======
type repository struct {
>>>>>>> 76086c78eaa33a197616b3ffc02667d8c3bf55b5
	dbClient *gorm.DB
}

type ContactRepository interface {
	GetAll()(contacts []models.Contact, err error)
	GetContact(contactId int)(contact models.Contact, err error)
	Add(contact models.Contact) error
	Delete(contactId int) error
<<<<<<< HEAD
}

func NewContactRepository(db *gorm.DB) *contactRepository {
	return &contactRepository{
		dbClient: db,
	}
}

func (repo *contactRepository) getAll() (contacts []models.Contact, err error){
=======
	Update(contact models.Contact) error
}

func NewContactRepository(db *gorm.DB) {
	dbClient = db
}

func (repo *repository) getAll() (contacts []models.Contact, err error){
>>>>>>> 76086c78eaa33a197616b3ffc02667d8c3bf55b5
	err = repo.dbClient.Find(&contacts).Error
	return contacts, err
}

<<<<<<< HEAD
func (repo *contactRepository) getContact(contactId int) (contact models.Contact, err error){
=======
func (repo *repository) getContact(contactId int) (contact models.Contact, err error){
>>>>>>> 76086c78eaa33a197616b3ffc02667d8c3bf55b5
	err = repo.dbClient.Where("id = ?", contactId).Find(&contact).Error
	return contact, err
}

<<<<<<< HEAD
func (repo *contactRepository) Add(contact models.Contact) error{
=======
func (repo *repository) add(contact models.Contact) error{
>>>>>>> 76086c78eaa33a197616b3ffc02667d8c3bf55b5
	err := repo.dbClient.Save(&contact).Error
	return err
}

<<<<<<< HEAD
func (repo *contactRepository) Delete(contactId int) error{
=======
func (repo *repository) delete(contactId int) error{
>>>>>>> 76086c78eaa33a197616b3ffc02667d8c3bf55b5
	err := repo.dbClient.Where("id = ?", contactId).Delete(models.Contact{}).Error
	return err
}

<<<<<<< HEAD
// Interface implementation here

func (repo *contactRepository) GetAll() (contacts []models.Contact, err error) {
	return repo.getAll()
}

func (repo *contactRepository) GetContact(contactId int) (contact models.Contact, err error) {
	return repo.getContact(contactId)
=======
func (repo *repository) update(contact models.Contact) error{
	err := repo.dbClient.Where("id = ?", contact.ID).Update(&contact).Error
	return err
}

// Interface implementation here

func (repo *repository) GetAll() ([]models.Contact, error) {
	return repo.getAll()
}

func (repo *repository) GetContact(contactId int) (models.Contact, error) {
	return repo.getContact(contactId)
}

func (repo *repository) Delete(contactId int) error {
	return repo.delete(contactId)
}

func (repo *repository) Update(contact models.Contact) error {
	return repo.update(contact)
>>>>>>> 76086c78eaa33a197616b3ffc02667d8c3bf55b5
}