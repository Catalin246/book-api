package data

import (
	"book-api/models"
	"sync"
)

var (
	books = []models.Book{
		{ID: nextID, Title: "The Alchemist", Author: "Paulo Coelho", Year: 1988},
	}
	nextID = 2
	mu     sync.Mutex
)

// GetAllBooks returns all books in memory
func GetAllBooks() []models.Book {
	return books
}

// GetBookByID returns a book by its ID
func GetBookByID(id int) *models.Book {
	for _, book := range books {
		if book.ID == id {
			return &book
		}
	}
	return nil
}

// AddBook adds a new book to the list
func AddBook(book models.Book) {
	mu.Lock()
	book.ID = nextID
	nextID++
	books = append(books, book)
	mu.Unlock()
}

// UpdateBook updates an existing book
func UpdateBook(id int, updatedBook models.Book) *models.Book {
	mu.Lock()
	defer mu.Unlock()

	for i, book := range books {
		if book.ID == id {
			updatedBook.ID = id
			books[i] = updatedBook
			return &updatedBook
		}
	}
	return nil
}

// DeleteBook deletes a book by ID
func DeleteBook(id int) bool {
	mu.Lock()
	defer mu.Unlock()

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return true
		}
	}
	return false
}
