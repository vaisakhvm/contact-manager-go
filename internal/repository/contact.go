package repository

import "contact-manager-go/internal/domain"

type ContactRepository interface {
	Add(contact domain.Contact) domain.Contact
	List() []domain.Contact
	GetByID(id int) (domain.Contact, bool)
	Delete(id int) bool
}
