package repository

import (
	"Agenda/models"
	"github.com/jinzhu/gorm"
)

type ContactRepository struct {
	dbClient *gorm.DB
}

type ContactRepositoryInterface interface {
	GetAll()(contacts []models.Contact, err error)
	GetContact(contactId int)(contact models.Contact, err error)
	Add(contact models.Contact) (err error)
	Delete(contactId int) (err error)
	Update(contact models.Contact) (err error)
	Search(query string)(contacts []models.Contact, err error)
}

func (repo *ContactRepository) getAll() (contacts []models.Contact, err error){
	err = repo.dbClient.Find(&contacts).Error
	return contacts, err
}

func (repo *ContactRepository) getContact(contactId int) (contact models.Contact, err error){
	err = repo.dbClient.Where("id = ?", contactId).Find(&contact).Error
	return contact, err
}

func (repo *ContactRepository) add(contact models.Contact) (err error){
	err = repo.dbClient.Save(&contact).Error
	return err
}

func (repo *ContactRepository) delete(contactId int) (err error){
	err = repo.dbClient.Where("id = ?", contactId).Delete(models.Contact{}).Error
	return err
}

func (repo *ContactRepository) update(contact models.Contact) (err error){
	err = repo.dbClient.Table("contacts").Where("id = ?", contact.ID).Update(&contact).Error
	return err
}

func (repo *ContactRepository) search(query string) (contacts []models.Contact, err error){
	query = "%"+query+"%"
	err = repo.dbClient.Where(`
		name like ? or 
		last_name like ? or
		company like ? or
		mail like ? or
		phone like ? or
		address like ?
	`,query,query,query,query,query,query).Find(&contacts).Error
	return contacts, err
}

// Interface implementation here

func (repo *ContactRepository) GetAll() (contacts []models.Contact, err error) {
	return repo.getAll()
}

func (repo *ContactRepository) GetContact(contactId int) (contact models.Contact, err error) {
	return repo.getContact(contactId)
}

func (repo *ContactRepository) Add(contact models.Contact) (err error) {
	return repo.add(contact)
}

func (repo *ContactRepository) Delete(contactId int) (err error) {
	return repo.delete(contactId)
}

func (repo *ContactRepository) Update(contact models.Contact) (err error) {
	return repo.update(contact)
}

func (repo *ContactRepository) Search(query string) (contacts []models.Contact, err error) {
	return repo.search(query)
}

// Factory Functions
func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{
		dbClient: db,
	}
}