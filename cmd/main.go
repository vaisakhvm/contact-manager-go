package main

import (
	"contact-manager-go/internal/contact/delivery/rest"
	"contact-manager-go/internal/contact/repository/inmemory"
	"contact-manager-go/internal/contact/usecase"
	"log"
	"net/http"
)

func main() {
	repo := inmemory.NewInMemoryRepository()
	svc := usecase.NewContactUsecase(repo)
	handler := rest.NewContactHandler(svc)

	http.HandleFunc("/add", handler.AddContact)
	http.HandleFunc("/list", handler.ListContacts)
	http.HandleFunc("/get", handler.GetContact)
	http.HandleFunc("/delete", handler.DeleteContact)

	log.Println("Server is listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
