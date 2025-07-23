package rest

import (
	"net/http"
	"strings"
)

func RegisterRoutes(mux *http.ServeMux, handler *ContactHandler) {
	mux.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.ListContacts(w, r)
		case http.MethodPost:
			handler.AddContact(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/contacts/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/contacts/")
		if id == "" {
			http.Error(w, "Missing contact ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			handler.GetContact(w, r, id)
		case http.MethodDelete:
			handler.DeleteContact(w, r, id)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
