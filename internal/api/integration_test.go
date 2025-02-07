// internal/api/integration_test.go
package api

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gorilla/mux"
)

func setupTestServer() (*mux.Router, *BookStore) {
    store := NewBookStore()
    handler := NewHandler(store)
    
    r := mux.NewRouter()
    r.HandleFunc("/books", handler.HandleAddBook).Methods("POST")
    r.HandleFunc("/books/{id}", handler.HandleGetBook).Methods("GET")
    
    return r, store
}

func TestBookLifecycle(t *testing.T) {
    router, _ := setupTestServer()
    server := httptest.NewServer(router)
    defer server.Close()

    // Test adding a book
    book := Book{
        ID:     "1",
        Title:  "Integration Test Book",
        Author: "Test Author",
        Year:   2024,
    }
    
    body, _ := json.Marshal(book)
    resp, err := http.Post(server.URL+"/books", "application/json", bytes.NewBuffer(body))
    if err != nil {
        t.Fatalf("Failed to create book: %v", err)
    }
    if resp.StatusCode != http.StatusCreated {
        t.Errorf("Expected status Created, got %v", resp.Status)
    }

    // Test retrieving the book
    resp, err = http.Get(server.URL + "/books/1")
    if err != nil {
        t.Fatalf("Failed to get book: %v", err)
    }
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status OK, got %v", resp.Status)
    }

    var retrievedBook Book
    json.NewDecoder(resp.Body).Decode(&retrievedBook)
    if retrievedBook != book {
        t.Errorf("Expected book %v, got %v", book, retrievedBook)
    }
}

func TestConcurrentAccess(t *testing.T) {
    router, store := setupTestServer()
    server := httptest.NewServer(router)
    defer server.Close()

    // Add initial book
    book := Book{ID: "1", Title: "Concurrent Test", Author: "Test", Year: 2024}
    store.AddBook(book)

    // Test concurrent reads
    done := make(chan bool)
    for i := 0; i < 10; i++ {
        go func() {
            resp, err := http.Get(server.URL + "/books/1")
            if err != nil || resp.StatusCode != http.StatusOK {
                t.Errorf("Concurrent read failed")
            }
            done <- true
        }()
    }

    // Wait for all goroutines
    for i := 0; i < 10; i++ {
        <-done
    }
}