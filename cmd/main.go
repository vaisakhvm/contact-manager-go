package main

import (
	contactRest "contact-manager-go/internal/contact/delivery/rest"
	contactInMemoryRepository "contact-manager-go/internal/contact/repository/inmemory"
	contactUsecase "contact-manager-go/internal/contact/usecase"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//Contact wiring
	contactRepository := contactInMemoryRepository.NewInMemoryContactRepository()
	contactUsecase := contactUsecase.NewContactUsecase(contactRepository)
	contactHandler := contactRest.NewContactHandler(contactUsecase)
	contactRest.RegisterRoutes(mux, contactHandler)

	log.Println("Server is listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
