package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strconv"
	"sync"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone int    `json:"phone"`
	Age   int    `json:"age"`
}

type ContactManager struct {
	mu       sync.Mutex
	contacts []Contact
	nextID   int
}

func NewContactManager() *ContactManager {
	return &ContactManager{
		contacts: make([]Contact, 0),
		nextID:   1,
	}
}

//Add new contact

func (cm *ContactManager) AddContact(w http.ResponseWriter, r *http.Request) {
	var newContact Contact
	err := json.NewDecoder(r.Body).Decode(&newContact)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	cm.mu.Lock()
	defer cm.mu.Unlock()

	newContact.ID = cm.nextID
	cm.nextID++
	cm.contacts = append(cm.contacts, newContact)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newContact)
}

// List contacts

func (cm *ContactManager) ListContacts(w http.ResponseWriter, r *http.Request) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cm.contacts)
}

//Get contact by ID

func (cm *ContactManager) GetContact(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	cm.mu.Lock()
	defer cm.mu.Unlock()

	for _, contact := range cm.contacts {
		if contact.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(contact)
			return
		}
	}

	http.Error(w, "Contact not found", http.StatusNotFound)
}

//Delete contact by ID 

func (cm *ContactManager) DeleteContact(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	cm.mu.Lock()
	defer cm.mu.Unlock()

	for i, contact := range cm.contacts {
		if contact.ID == id {
			// remove contact by slicing
			// cm.contacts = append(cm.contacts[:i], cm.contacts[i+1:]...)

			cm.contacts = slices.Delete(cm.contacts, i, i+1)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(contact)
			return
		}
	}

	http.Error(w, "Contact not found", http.StatusNotFound)
}

func main() {
	contactManager := NewContactManager()

	http.HandleFunc("/add", contactManager.AddContact)
	http.HandleFunc("/list", contactManager.ListContacts)
	http.HandleFunc("/get", contactManager.GetContact)
	http.HandleFunc("/delete", contactManager.DeleteContact)

	http.ListenAndServe(":8080", nil)
}
