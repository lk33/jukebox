#!/bin/bash

# Create directories
mkdir -p ./cmd ./internal/handlers ./internal/models ./internal/repositories ./internal/services ./pkg/utils ./config ./migrations ./scripts ./static ./templates ./tests

# Create files
touch ./cmd/main.go ./internal/handlers/handlers.go ./internal/models/models.go ./internal/repositories/repository.go ./internal/services/service.go ./pkg/utils/utils.go ./config/config.go ./.gitignore ./go.mod ./README.md

# Output message
echo "Go project template created successfully!"
