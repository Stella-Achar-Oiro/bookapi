package api

import "errors"

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

var (
	ErrBookNotFound = errors.New("book not found")
	ErrBookExists   = errors.New("book already exists")
)
