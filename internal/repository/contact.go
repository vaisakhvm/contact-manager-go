package repository

import "contact-manager-go/internal/models"

type ContactRepository interface {
	Add(contact models.Contact) models.Contact
	List() []models.Contact
	GetByID(id int) (models.Contact, bool)
	Delete(id int) bool
}
