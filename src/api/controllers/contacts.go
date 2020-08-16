package controllers

import (
	"Agenda/models"
	"Agenda/repository"
	"Agenda/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type contactsController struct {
	dbConn *gorm.DB
	contactService *services.ContactService
}

type ContactsController interface {
	GetAll(c *gin.Context)
	GetContact(c *gin.Context)
	Add(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Search(c *gin.Context)
}

// getAll godoc
// @Summary Get all contacts
// @Description Get all contacts
// @Accept  json
// @Produce  json
// @Success 200
// @Router /contactss [get]
func (ctrl contactsController) getAll(c *gin.Context) {
	contacts, err := ctrl.contactService.GetAll()
	if err != nil {
		httpError := models.NewHttpError(http.StatusInternalServerError,"internal error")
		models.ReturnHttpError(c, httpError)
		return
	}
	c.JSON(http.StatusOK, contacts)
}

// getContact godoc
// @Summary Get contact by id
// @Description Retrieve a contact using his ID
// @Accept  json
// @Produce  json
// @Param id query string true "contact search by id"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /contacts/id/:id [get]
func (ctrl contactsController) getContact(c *gin.Context) {
	contactIdStr := c.Param("id")
	if contactIdStr == "" {
		httpError := models.NewHttpError(http.StatusBadRequest,"bad request.")
		models.ReturnHttpError(c, httpError)
		return
	}
	contactId, _ := strconv.Atoi(contactIdStr)
	contact, err := ctrl.contactService.GetContact(contactId)
	if err != nil {
		code := 500
		if err == gorm.ErrRecordNotFound {
			code = 404
		}
		httpError := models.NewHttpError(code, err.Error())
		models.ReturnHttpError(c, httpError)
		return
	}
	c.JSON(http.StatusOK, contact)
}

// add godoc
// @Summary Add new contact
// @Description Create new contact
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 500
// @Router /contacts [post]
func (ctrl contactsController) add(c *gin.Context) {
	var contact models.Contact
	c.BindJSON(&contact)

	if contact.Name == "" || contact.Mail == "" {
		httpError := models.NewHttpError(http.StatusBadRequest,"bad request.")
		models.ReturnHttpError(c, httpError)
		return
	}
	err := ctrl.contactService.Add(contact)
	if err != nil {
		httpError := models.NewHttpError(http.StatusInternalServerError,"internal error")
		models.ReturnHttpError(c, httpError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "the user "+contact.Name+" has been added!",
	})
}

// delete godoc
// @Summary Delete contact by id
// @Description delete a contact using his ID
// @Accept  json
// @Produce  json
// @Param id query string true "contact delete by id"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /contacts/id/:id [delete]
func (ctrl contactsController) delete(c *gin.Context) {
	contactIdStr := c.Param("id")
	if contactIdStr == "" {
		httpError := models.NewHttpError(http.StatusBadRequest,"bad request.")
		models.ReturnHttpError(c, httpError)
		return
	}

	contactId, _ := strconv.Atoi(contactIdStr)
	err := ctrl.contactService.Delete(contactId)
	if err != nil {
		httpError := models.NewHttpError(http.StatusInternalServerError,"internal error")
		models.ReturnHttpError(c, httpError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "the user id "+contactIdStr+" has been deleted!",
	})
}

// update godoc
// @Summary Update contact by id
// @Description update a contact using his ID
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 404
// @Failure 500
// @Router /contacts/id/:id [put]
func (ctrl contactsController) update(c *gin.Context) {
	var contact models.Contact
	c.BindJSON(&contact)

	if contact.ID == 0 {
		httpError := models.NewHttpError(http.StatusBadRequest,"bad request.")
		models.ReturnHttpError(c, httpError)
		return
	}
	err := ctrl.contactService.Update(contact)
	if err != nil {
		httpError := models.NewHttpError(http.StatusInternalServerError,"internal error")
		models.ReturnHttpError(c, httpError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "the user "+contact.Name+" has been updated!",
	})
}

// search godoc
// @Summary Search contact by id
// @Description search a contact using his ID
// @Accept  json
// @Produce  json
// @Param query query string true "contact search by query"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /contacts/search [get]
func (ctrl contactsController) search(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		httpError := models.NewHttpError(http.StatusBadRequest,"bad request.")
		models.ReturnHttpError(c, httpError)
		return
	}
	contacts, err := ctrl.contactService.Search(query)
	if err != nil {
		code := 500
		if err == gorm.ErrRecordNotFound {
			code = 404
		}
		httpError := models.NewHttpError(code, err.Error())
		models.ReturnHttpError(c, httpError)
		return
	}
	c.JSON(http.StatusOK, contacts)
}

// Interface implementation here

func (ctrl contactsController) GetAll(c *gin.Context) {
	ctrl.getAll(c)
}

func (ctrl contactsController) GetContact(c *gin.Context) {
	ctrl.getContact(c)
}

func (ctrl contactsController) Add(c *gin.Context) {
	ctrl.add(c)
}

func (ctrl contactsController) Delete(c *gin.Context) {
	ctrl.delete(c)
}

func (ctrl contactsController) Update(c *gin.Context) {
	ctrl.update(c)
}

func (ctrl contactsController) Search(c *gin.Context) {
	ctrl.search(c)
}

// Factory Functions

func NewContactController() *contactsController{
	db := repository.GetDbConn()
	contactService := services.NewContactService(db)
	return &contactsController{
		dbConn: db,
		contactService: contactService,
	}
}