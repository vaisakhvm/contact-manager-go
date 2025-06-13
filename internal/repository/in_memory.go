package repository

import (
	"contact-manager-go/internal/models"
	"slices"
	"sync"
)

type InMemoryRepo struct {
	mu       sync.Mutex
	nextID   int
	contacts []models.Contact
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		contacts: make([]models.Contact, 0),
		nextID:   1,
	}
}

func (r *InMemoryRepo) Add(contact models.Contact) models.Contact {
	r.mu.Lock()
	defer r.mu.Unlock()

	contact.ID = r.nextID
	r.nextID++
	r.contacts = append(r.contacts, contact)
	return contact
}

func (r *InMemoryRepo) List() []models.Contact {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.contacts
}

func (r *InMemoryRepo) GetByID(id int) (models.Contact, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, contact := range r.contacts {
		if contact.ID == id {
			return contact, true
		}
	}
	return models.Contact{}, false
}

func (r *InMemoryRepo) Delete(id int) bool {
	for i, contact := range r.contacts {
		if contact.ID == id {
			r.contacts = slices.Delete(r.contacts, i, i+1)
			return true
		}
	}

	return false
}
