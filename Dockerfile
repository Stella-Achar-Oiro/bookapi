FROM golang:1.23.4-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server cmd/server/main.go

FROM alpine:latest
COPY --from=builder /app/server /server
EXPOSE 8000
CMD ["/server"]

