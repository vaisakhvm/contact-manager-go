package domain

type ContactRepository interface {
	Add(contact Contact) Contact
	List() []Contact
	GetByID(id int) (Contact, bool)
	Delete(id int) bool
}
