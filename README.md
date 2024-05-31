
# Golang gRPC User Service

## Introduction

This project implements a gRPC user service in Go for managing user details and includes search functionality. The project includes a Docker setup for easy deployment.

## Prerequisites

- Go 1.16 or later
- Protocol Buffers compiler (`protoc`)
- Docker
- Postman (for testing)

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
- `AddUser`

## Configuration Details

No specific configuration required.

## Testing with Postman

You can use Postman to interact with the gRPC service. Follow these steps:

1. **Open Postman and Create a New gRPC Request:**
   - Click on the `New` button.
   - Select `Request`.
   - Name your request (e.g., `AddUser`) and save it to a collection.

2. **Select gRPC as the Protocol:**
   - In the new request tab, select `gRPC` from the dropdown menu.

3. **Enter the Server URL:**
   - Enter the server URL, which is `localhost:50051` if your server is running locally.

4. **Load the `.proto` File:**
   - Click on the `+ Add Proto File` button.
   - Select and upload your `user.proto` file.
   - Postman will parse the `.proto` file and show the available services and methods.

### Example Requests

1. **AddUser Request:**
   - Select the `UserService` service.
   - Choose the `AddUser` method.
   - Fill in the request data. For example:

   ```json
   {
     "fname": "Himanshu",
     "city": "Gurgaon",
     "phone": 9058613131,
     "height": 5.5,
     "married": false
   }
   ```

   - Click on the `Invoke` button to send the request.
   - Postman will display the response from the gRPC server, including the assigned user ID.

2. **GetUserDetails Request:**
   - Select the `UserService` service.
   - Choose the `GetUserDetails` method.
   - Fill in the request data. For example:

   ```json
   {
     "id": 1
   }
   ```

   - Click on the `Invoke` button to send the request.
   - Postman will display the response with the user details.

3. **GetUsersDetails Request:**
   - Select the `UserService` service.
   - Choose the `GetUsersDetails` method.
   - Fill in the request data. For example:

   ```json
   {
     "ids": [1, 2]
   }
   ```

   - Click on the `Invoke` button to send the request.
   - Postman will display the response with the details of the specified users.

4. **SearchUsers Request:**
   - Select the `UserService` service.
   - Choose the `SearchUsers` method.
   - Fill in the request data. For example:

   ```json
   {
     "city": "Gurgaon",
     "phone": 9058613131,
     "married": false
   }
   ```

   - Click on the `Invoke` button to send the request.
   - Postman will display the response with the users matching the search criteria.

### Example `grpcurl` Commands

If you prefer to use `grpcurl`:

1. **Get User Details:**

   ```sh
   grpcurl -plaintext -d '{"id": 1}' localhost:50051 user.UserService/GetUserDetails
   ```

2. **Search Users:**

   ```sh
   grpcurl -plaintext -d '{"city": "Gurgaon"}' localhost:50051 user.UserService/SearchUsers
   ```

3. **Add User:**

   ```sh
   grpcurl -plaintext -d '{
     "fname": "Himanshu",
     "city": "Gurgaon",
     "phone": 9058613131,
     "height": 5.5,
     "married": false
   }' localhost:50051 user.UserService/AddUser
   ```

## License

This project is licensed under the MIT License. 