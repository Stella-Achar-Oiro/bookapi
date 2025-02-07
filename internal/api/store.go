// internal/api/store.go
package api

import (
    "sync"
)

type BookStore struct {
    mu    sync.RWMutex
    books map[string]Book
}

func NewBookStore() *BookStore {
    return &BookStore{
        books: make(map[string]Book),
    }
}

func (s *BookStore) AddBook(book Book) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    if _, exists := s.books[book.ID]; exists {
        return ErrBookExists
    }

    s.books[book.ID] = book
    return nil
}

func (s *BookStore) GetBook(id string) (Book, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    book, exists := s.books[id]
    if !exists {
        return Book{}, ErrBookNotFound
    }

    return book, nil
}

func (s *BookStore) ListBooks() []Book {
    s.mu.RLock()
    defer s.mu.RUnlock()

    books := make([]Book, 0, len(s.books))
    for _, book := range s.books {
        books = append(books, book)
    }
    return books
}