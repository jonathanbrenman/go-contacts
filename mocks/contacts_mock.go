package mocks

import (
	"Agenda/models"
	"github.com/jinzhu/gorm"
)

func LoadMockContacts(db *gorm.DB) {
	// Save contacts to memory database
	for _, contact := range getContacts() {
		err := db.Save(&contact).Error
		if err != nil {
			panic(err)
		}
	}
}

func getContacts() []models.Contact {
	// Create some mock contacts
	var contacts []models.Contact

	homero := models.Contact{
		Name: "Homero",
		LastName: "Simpsons",
		Company: "Chebronlet",
		Phone: "4934-2345",
		Mail: "homerothomson@simpsons.com",
		Address: "av siemprevivas",
		Thumb: "https://vignette.wikia.nocookie.net/lossimpson/images/b/bd/Homer_Simpson.png/revision/latest?cb=20100522180809&path-prefix=es",
	}
	contacts = append(contacts, homero)

	marsh := models.Contact{
		Name: "Marsh",
		LastName: "Simpsons",
		Company: "Toyita",
		Phone: "4934-2346",
		Mail: "marsh@simpsons.com",
		Address: "av siemprevivas",
		Thumb: "https://vignette.wikia.nocookie.net/lossimpson/images/0/0b/Marge_Simpson.png/revision/latest?cb=20090415001251&path-prefix=es",
	}
	contacts = append(contacts, marsh)

	bart := models.Contact{
		Name: "Bartolomeo",
		LastName: "Simpsons",
		Company: "Citroeniete",
		Phone: "4934-2347",
		Mail: " bart@simpsons.com",
		Address: "av siemprevivas",
		Thumb: "https://upload.wikimedia.org/wikipedia/en/a/aa/Bart_Simpson_200px.png",
	}
	contacts = append(contacts,  bart)

	lisa := models.Contact{
		Name: "Lisa",
		LastName: "Simpsons",
		Company: "Volkswagenator",
		Phone: "4934-2348",
		Mail: " lisa@simpsons.com",
		Address: "av siemprevivas",
		Thumb: "https://upload.wikimedia.org/wikipedia/en/thumb/e/ec/Lisa_Simpson.png/220px-Lisa_Simpson.png",
	}
	contacts = append(contacts,  lisa)

	magie := models.Contact{
		Name: "Magie",
		LastName: "Simpsons",
		Company: "Peugito",
		Phone: "4934-2348",
		Mail: " lisa@simpsons.com",
		Address: "av siemprevivas",
		Thumb: "https://upload.wikimedia.org/wikipedia/en/9/9d/Maggie_Simpson.png",
	}
	contacts = append(contacts,  magie)

	return contacts
}