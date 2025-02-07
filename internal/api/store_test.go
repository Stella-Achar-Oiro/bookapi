// internal/api/store_test.go
package api

import (
    "testing"
)

func TestAddBook(t *testing.T) {
    store := NewBookStore()
    book := Book{
        ID:     "1",
        Title:  "Test Book",
        Author: "Test Author",
        Year:   2024,
    }

    err := store.AddBook(book)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    // Test duplicate book
    err = store.AddBook(book)
    if err != ErrBookExists {
        t.Errorf("Expected ErrBookExists, got %v", err)
    }
}

func TestGetBook(t *testing.T) {
    store := NewBookStore()
    book := Book{
        ID:     "1",
        Title:  "Test Book",
        Author: "Test Author",
        Year:   2024,
    }

    store.AddBook(book)

    // Test existing book
    got, err := store.GetBook("1")
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if got != book {
        t.Errorf("Expected %v, got %v", book, got)
    }

    // Test non-existent book
    _, err = store.GetBook("2")
    if err != ErrBookNotFound {
        t.Errorf("Expected ErrBookNotFound, got %v", err)
    }
}

func TestListBooks(t *testing.T) {
    store := NewBookStore()
    books := []Book{
        {ID: "1", Title: "Book 1", Author: "Author 1", Year: 2024},
        {ID: "2", Title: "Book 2", Author: "Author 2", Year: 2024},
    }

    for _, book := range books {
        store.AddBook(book)
    }

    got := store.ListBooks()
    if len(got) != len(books) {
        t.Errorf("Expected %d books, got %d", len(books), len(got))
    }
}