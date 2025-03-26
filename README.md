# Go Application

This is a Go application that provides a web service. You can run it either using Docker Compose or directly with Go.

## Prerequisites

- Go 1.16+ (if running directly)
- Docker and Docker Compose (if using containers)

## Running with Docker Compose (Recommended)

1. Clone this repository:
   ```bash
   git clone https://github.com/rogdevil/flamingo-test.git
   cd flamingo-test
   ```
2. Build and start the containers:
    ```bash
    docker-compose up -d
    ```
3. Run the project
    ```bash
    go mod download
    go run main.go serve
    ```
4. Run frontend
    ```bash
    cd ago-fe && npm run dev
    ```