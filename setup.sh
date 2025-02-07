#!/bin/bash

# Create project directory and subdirectories
mkdir -p bookapi/{cmd/server,internal/{api,config},docs}

# Create go.mod
cd bookapi
go mod init github.com/Stella-Achar-Oiro/bookapi
go get github.com/gorilla/mux

# Create empty files
touch cmd/server/main.go \
      internal/api/{handler,handler_test,model,store}.go \
      internal/config/config.go \
      docs/{api.md,swagger.yaml} \
      README.md

# Make script executable
chmod +x setup.sh

echo "Project structure created successfully"