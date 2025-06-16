package inmemory

import (
	"contact-manager-go/domain"
	"slices"
	"sync"
)

type InMemoryRepository struct {
	mu       sync.Mutex
	nextID   int
	contacts []domain.Contact
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		contacts: make([]domain.Contact, 0),
		nextID:   1,
	}
}

func (r *InMemoryRepository) Add(contact domain.Contact) domain.Contact {
	r.mu.Lock()
	defer r.mu.Unlock()

	contact.ID = r.nextID
	r.nextID++
	r.contacts = append(r.contacts, contact)
	return contact
}

func (r *InMemoryRepository) List() []domain.Contact {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.contacts
}

func (r *InMemoryRepository) GetByID(id int) (domain.Contact, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, contact := range r.contacts {
		if contact.ID == id {
			return contact, true
		}
	}
	return domain.Contact{}, false
}

func (r *InMemoryRepository) Delete(id int) bool {
	for i, contact := range r.contacts {
		if contact.ID == id {
			r.contacts = slices.Delete(r.contacts, i, i+1)
			return true
		}
	}

	return false
}
