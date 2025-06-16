package usecase

import (
	"contact-manager-go/domain"
	"fmt"
)

type ContactUsecase struct {
	repo domain.ContactRepository
}

func NewContactUsecase(repo domain.ContactRepository) *ContactUsecase {
	return &ContactUsecase{repo: repo}
}

func (s *ContactUsecase) AddContact(c domain.Contact) (domain.Contact, error) {
	return s.repo.Add(c), nil
}

func (s *ContactUsecase) ListContacts() []domain.Contact {
	return s.repo.List()
}

func (s *ContactUsecase) GetByID(id int) (domain.Contact, error) {
	contact, found := s.repo.GetByID(id)
	if !found {
		return domain.Contact{}, fmt.Errorf("contact with ID %d not found", id)
	}

	return contact, nil
}

func (s *ContactUsecase) Delete(id int) error {
	if !s.repo.Delete(id) {
		return fmt.Errorf("contact with ID %d not found", id)
	}

	return nil
}
