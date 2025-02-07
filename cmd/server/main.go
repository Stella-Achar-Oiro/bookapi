package main

import (
    "log"
    "net/http"

    "github.com/Stella-Achar-Oiro/bookapi/internal/api"
    "github.com/Stella-Achar-Oiro/bookapi/internal/config"
    "github.com/gorilla/mux"
)

func main() {
    cfg := config.New()
    store := api.NewBookStore()
    handler := api.NewHandler(store)

    r := mux.NewRouter()
    r.HandleFunc("/books", handler.HandleAddBook).Methods("POST")
    r.HandleFunc("/books/{id}", handler.HandleGetBook).Methods("GET")

    log.Printf("Server starting on %s", cfg.Address)
    log.Fatal(http.ListenAndServe(cfg.Address, r))
}