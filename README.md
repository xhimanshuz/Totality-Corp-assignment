
# Golang gRPC User Service

## Introduction

This project implements a gRPC user service in Go for managing user details and includes search functionality. The project includes a Docker setup for easy deployment.

## Prerequisites

- Go 1.16 or later
- Protocol Buffers compiler (`protoc`)
- Docker

## Setup

### Step 1: Clone the Repository

```sh
git clone https://github.com/xhimanshuz/Totality-Corp-assignment.git
cd Totality-Corp-assignment
```

### Step 2: Initialize the Go Module

```sh
go mod init github.com/xhimanshuz/Totality-Corp-assignment
go mod tidy
```

### Step 3: Generate gRPC Code

Make sure `protoc` and the Go plugins are installed and available in your `PATH`.

```sh
protoc --go_out=. --go-grpc_out=. user.proto
```

### Step 4: Build and Run the Application

#### Using Go

```sh
go run server.go
```

#### Using Docker

1. **Build the Docker Image:**

   ```sh
   docker build -t golang-grpc-user-service .
   ```

2. **Run the Docker Container:**

   ```sh
   docker run -p 50051:50051 golang-grpc-user-service
   ```

## Access the gRPC Service Endpoints

- `GetUserDetails`
- `GetUsersDetails`
- `SearchUsers`

## Configuration Details

No specific configuration required.

## Testing

You can test the gRPC service using `grpcurl` or any other gRPC client.

### Example `grpcurl` Commands

1. **Get User Details:**

   ```sh
   grpcurl -plaintext -d '{"id": 1}' localhost:50051 user.UserService/GetUserDetails
   ```

2. **Search Users:**

   ```sh
   grpcurl -plaintext -d '{"city": "LA"}' localhost:50051 user.UserService/SearchUsers
   ```

## License

This project is licensed under the MIT License.
