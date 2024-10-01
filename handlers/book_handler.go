package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"book-api/data"
	"book-api/models"

	"github.com/gorilla/mux"
)

// GetBooks returns all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := data.GetAllBooks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook returns a single book by ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	book := data.GetBookByID(id)

	if book == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// CreateBook adds a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data.AddBook(book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// UpdateBook modifies an existing book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedBook := data.UpdateBook(id, book)
	if updatedBook == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

// DeleteBook removes a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if data.DeleteBook(id) {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	http.NotFound(w, r)
}
