#!/bin/bash

# Build for Linux
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o target/delete-gitlab-project-linux main.go

# Build for Mac
echo "Building for Mac..."
GOOS=darwin GOARCH=amd64 go build -o target/delete-gitlab-project-mac main.go

# Build for Mac M1
echo "Building for Mac M1..."
GOOS=darwin GOARCH=arm64 go build -o target/delete-gitlab-project-mac-m1 main.go

echo "Build complete!"
