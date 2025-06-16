package rest

import "net/http"

func RegisterRoutes(mux *http.ServeMux, handler *ContactHandler) {
	mux.HandleFunc("/add", handler.AddContact)
	mux.HandleFunc("/list", handler.ListContacts)
	mux.HandleFunc("/get", handler.GetContact)
	mux.HandleFunc("/delete", handler.DeleteContact)
}
