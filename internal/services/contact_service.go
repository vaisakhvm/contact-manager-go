package services

import (
	"contact-manager-go/internal/models"
	"contact-manager-go/internal/repository"
	"fmt"
)

type ContactService struct {
	repo repository.ContactRepository
}

func NewContactService(repo repository.ContactRepository) *ContactService {
	return &ContactService{repo: repo}
}

func (s *ContactService) AddContact(c models.Contact) (models.Contact, error) {
	return s.repo.Add(c), nil
}

func (s *ContactService) ListContacts() []models.Contact {
	return s.repo.List()
}

func (s *ContactService) GetByID(id int) (models.Contact, error) {
	contact, found := s.repo.GetByID(id)
	if !found {
		return models.Contact{}, fmt.Errorf("contact with ID %d not found", id)
	}

	return contact, nil
}

func (s *ContactService) Delete(id int) error {
	if !s.repo.Delete(id) {
		return fmt.Errorf("contact with ID %d not found", id)
	}

	return nil
}
