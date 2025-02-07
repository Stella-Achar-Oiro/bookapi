# API Usage Examples

## Adding Books

Add a new book to the system:
```bash
curl -X POST http://localhost:8000/books \
  -H "Content-Type: application/json" \
  -d '{
    "id": "1",
    "title": "Go Programming",
    "author": "Stella Oiro",
    "year": 2024
  }'
```

Expected successful response:
```
Status: 201 Created
```

## Retrieving Books

Get a specific book by ID:
```bash
curl http://localhost:8000/books/1
```

Successful response:
```json
{
    "id": "1",
    "title": "Go Programming",
    "author": "Stella Oiro",
    "year": 2024
}
```

## Error Scenarios

### 1. Book Not Found
```bash
curl http://localhost:8000/books/999
```
Response:
```
Status: 404 Not Found
Body: Book not found
```

### 2. Invalid JSON Format
```bash
curl -X POST http://localhost:8000/books \
  -H "Content-Type: application/json" \
  -d '{invalid json}'
```
Response:
```
Status: 400 Bad Request
Body: Invalid request body
```

### 3. Duplicate Book ID
```bash
curl -X POST http://localhost:8000/books \
  -H "Content-Type: application/json" \
  -d '{
    "id": "1",
    "title": "Another Book",
    "author": "Stella Oiro",
    "year": 2024
  }'
```
Response:
```
Status: 409 Conflict
Body: Book already exists
```

## Testing Script
```bash
#!/bin/bash
# Test all endpoints
./test_api.sh
```