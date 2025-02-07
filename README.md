# Book API

A RESTful API for managing books built with Go.

## Features
- Add and retrieve books
- Docker containerization
- Comprehensive test coverage
- Clear API documentation

## Prerequisites
- Go 1.23.4+
- Docker
- Make

## Installation
```bash
git clone <repository-url>
cd bookapi
go mod download
```

## Running the API

With Go:
```bash
make build
make run
```

With Docker:
```bash
docker-compose up --build
```

## API Usage

Add a book:
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

Get a book:
```bash
curl http://localhost:8000/books/1
```

## Testing
```bash
make test
```

## Project Structure
```
bookapi/
├── cmd/server/        # Entry point
├── internal/          # Private packages
│   ├── api/          # API handlers
│   └── config/       # Configuration
└── docs/             # Documentation
```

## License
MIT

## Author
Stella Oiro