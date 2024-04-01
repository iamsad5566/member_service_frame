# Application Overview
This backend application provides a comprehensive member service system with both RESTful APIs and a gRPC service. It supports functionalities like user registration, login, password update, and user existence verification. The application also automates the creation of the necessary PostgreSQL database and includes a gRPC service for token-based authorization.
- [Features](#features)
- [gRPC Service](#grpc-service)
- [Getting Started](#getting-started)
- [Usage](#usage)

## Features
- **Automatic Database Creation:** Automatically creates the specified PostgreSQL database if it doesn't exist.
- **RESTful Member Service APIs:** Supports user registration, login, password updates, and user existence checks.
- **gRPC Authorization Service:** A gRPC service that validates user tokens for authorization, providing an additional layer of security.

## gRPC Service
The gRPC service within the application offers token-based authorization. It validates incoming tokens to determine whether a user is authorized to perform certain operations.

gRPC Authorization Service Details:
- **Service Name:** AuthorizerServer
- **Method:** AuthorizeByToken
- **Functionality:** Validates user tokens sent in metadata. It checks for token presence, validates the token, and responds with the authorization status. The service can also refresh tokens if necessary.

## Getting Started
Running the Application
1. Ensure that PostgreSQL and Redis are running.
2. Configure the application to connect to your database and Redis server.
3. Start the application. It will initiate the RESTful server and the gRPC service.
```
go run .
```
4. The RESTful APIs will be available on the configured port, and the gRPC service will be listening on its designated port.

## Usage
- **RESTful API:** Access the member service APIs for user-related operations.
- **gRPC Service:** Utilize gRPC clients to interact with the authorization service, sending tokens for validation and receiving authorization responses.
