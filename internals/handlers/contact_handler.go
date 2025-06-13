package handlers

import (
	"contact-manager-go/internals/models"
	"contact-manager-go/internals/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type ContactHandler struct {
	svc *services.ContactService
}

func NewContactHandler(service *services.ContactService) *ContactHandler {
	return &ContactHandler{svc: service}
}

func (h *ContactHandler) AddContact(w http.ResponseWriter, r *http.Request) {
	var newContact models.Contact
	err := json.NewDecoder(r.Body).Decode(&newContact)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	added, err := h.svc.AddContact(newContact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(added)
}

func (h *ContactHandler) ListContacts(w http.ResponseWriter, r *http.Request) {
	contactList := h.svc.ListContacts()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contactList)
}

func (h *ContactHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	contact, err := h.svc.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact)
}

func (h *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	err = h.svc.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	http.Error(w, "Contact not found", http.StatusNoContent)
}
