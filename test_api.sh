#!/bin/bash
# test_api.sh

echo "Testing Book API..."

echo -e "\n1. Adding a new book..."
curl -X POST http://localhost:8000/books \
  -H "Content-Type: application/json" \
  -d '{
    "id": "1",
    "title": "Go Programming",
    "author": "Stella Oiro",
    "year": 2024
  }'

echo -e "\n\n2. Getting the book..."
curl http://localhost:8000/books/1

echo -e "\n\n3. Testing error handling..."
echo "3.1 Book not found:"
curl http://localhost:8000/books/999

echo -e "\n\n3.2 Invalid JSON:"
curl -X POST http://localhost:8000/books \
  -H "Content-Type: application/json" \
  -d '{invalid}'

echo -e "\n\n3.3 Duplicate book:"
curl -X POST http://localhost:8000/books \
  -H "Content-Type: application/json" \
  -d '{
    "id": "1",
    "title": "Go Programming",
    "author": "Stella Oiro",
    "year": 2024
  }'