package main

import (
	"contact-manager-go/internal/handlers"
	"contact-manager-go/internal/repository"
	"contact-manager-go/internal/usecase"
	"log"
	"net/http"
)

func main() {
	repo := repository.NewInMemoryRepo()
	svc := usecase.NewContactUsecase(repo)
	handler := handlers.NewContactHandler(svc)

	http.HandleFunc("/add", handler.AddContact)
	http.HandleFunc("/list", handler.ListContacts)
	http.HandleFunc("/get", handler.GetContact)
	http.HandleFunc("/delete", handler.DeleteContact)

	log.Println("Server is listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
