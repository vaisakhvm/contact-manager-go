package inmemory

import (
	"contact-manager-go/domain"
	"slices"
	"sync"
)

type InMemoryContactRepository struct {
	mu       sync.Mutex
	nextID   int
	contacts []domain.Contact
}

func NewInMemoryContactRepository() *InMemoryContactRepository {
	return &InMemoryContactRepository{
		contacts: make([]domain.Contact, 0),
		nextID:   1,
	}
}

func (r *InMemoryContactRepository) Add(contact domain.Contact) domain.Contact {
	r.mu.Lock()
	defer r.mu.Unlock()

	contact.ID = r.nextID
	r.nextID++
	r.contacts = append(r.contacts, contact)
	return contact
}

func (r *InMemoryContactRepository) List() []domain.Contact {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.contacts
}

func (r *InMemoryContactRepository) GetByID(id int) (domain.Contact, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, contact := range r.contacts {
		if contact.ID == id {
			return contact, true
		}
	}
	return domain.Contact{}, false
}

func (r *InMemoryContactRepository) Delete(id int) bool {
	for i, contact := range r.contacts {
		if contact.ID == id {
			r.contacts = slices.Delete(r.contacts, i, i+1)
			return true
		}
	}

	return false
}
