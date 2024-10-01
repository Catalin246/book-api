package main

import (
	"log"
	"net/http"

	"book-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	log.Println("Starting server on :8000...")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
