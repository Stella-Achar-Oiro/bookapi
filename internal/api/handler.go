// internal/api/handler.go
package api

import (
    "encoding/json"
    "net/http"
    
    "github.com/gorilla/mux"
)

type Handler struct {
    store *BookStore
}

func NewHandler(store *BookStore) *Handler {
    return &Handler{store: store}
}

// @Summary Add a new book
// @Description Add a new book to the collection
// @Tags books
// @Accept json
// @Produce json
// @Param book body Book true "Book object"
// @Success 201 "Created"
// @Failure 400 "Invalid request"
// @Failure 409 "Book already exists"
// @Router /books [post]
func (h *Handler) HandleAddBook(w http.ResponseWriter, r *http.Request) {
    var book Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := h.store.AddBook(book); err != nil {
        if err == ErrBookExists {
            http.Error(w, "Book already exists", http.StatusConflict)
            return
        }
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

// @Summary Get a book by ID
// @Description Get book details by ID
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} Book
// @Failure 404 "Book not found"
// @Router /books/{id} [get]
func (h *Handler) HandleGetBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    book, err := h.store.GetBook(id)
    if err == ErrBookNotFound {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(book)
}