// internal/api/handler_test.go
package api

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gorilla/mux"
)

func TestHandleAddBook(t *testing.T) {
    store := NewBookStore()
    handler := NewHandler(store)

    tests := []struct {
        name       string
        book      Book
        wantStatus int
    }{
        {
            name: "valid book",
            book: Book{ID: "1", Title: "Test Book", Author: "Author", Year: 2024},
            wantStatus: http.StatusCreated,
        },
        {
            name: "duplicate book",
            book: Book{ID: "1", Title: "Test Book", Author: "Author", Year: 2024},
            wantStatus: http.StatusConflict,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            body, _ := json.Marshal(tt.book)
            req := httptest.NewRequest("POST", "/books", bytes.NewBuffer(body))
            w := httptest.NewRecorder()

            handler.HandleAddBook(w, req)

            if w.Code != tt.wantStatus {
                t.Errorf("HandleAddBook() status = %v, want %v", w.Code, tt.wantStatus)
            }
        })
    }
}

func TestHandleGetBook(t *testing.T) {
    store := NewBookStore()
    handler := NewHandler(store)
    
    // Add test book
    book := Book{ID: "1", Title: "Test Book", Author: "Author", Year: 2024}
    store.AddBook(book)

    tests := []struct {
        name       string
        bookID     string
        wantStatus int
        wantBook   *Book
    }{
        {
            name:       "existing book",
            bookID:     "1",
            wantStatus: http.StatusOK,
            wantBook:   &book,
        },
        {
            name:       "non-existent book",
            bookID:     "999",
            wantStatus: http.StatusNotFound,
            wantBook:   nil,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := httptest.NewRequest("GET", "/books/"+tt.bookID, nil)
            w := httptest.NewRecorder()

            // Setup router context
            router := mux.NewRouter()
            router.HandleFunc("/books/{id}", handler.HandleGetBook)
            router.ServeHTTP(w, req)

            if w.Code != tt.wantStatus {
                t.Errorf("HandleGetBook() status = %v, want %v", w.Code, tt.wantStatus)
            }

            if tt.wantBook != nil {
                var gotBook Book
                json.NewDecoder(w.Body).Decode(&gotBook)
                if gotBook != *tt.wantBook {
                    t.Errorf("HandleGetBook() book = %v, want %v", gotBook, tt.wantBook)
                }
            }
        })
    }
}

func TestHandleInvalidJSON(t *testing.T) {
    store := NewBookStore()
    handler := NewHandler(store)

    // Send invalid JSON
    req := httptest.NewRequest("POST", "/books", bytes.NewBufferString("{invalid json}"))
    w := httptest.NewRecorder()

    handler.HandleAddBook(w, req)

    if w.Code != http.StatusBadRequest {
        t.Errorf("HandleAddBook() with invalid JSON status = %v, want %v", w.Code, http.StatusBadRequest)
    }
}