package routes

import (
	"book-api/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
}
