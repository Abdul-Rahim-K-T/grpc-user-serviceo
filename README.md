# gRPC User Service
This project is a gRPC User built with Go, following Clean Architecture principles.It includes functionalities for fetching user details based on user ID, retrieving a list of user details based on a list of user IDs, and searching user details based on specific criteria.
## Table of Contents
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Setup Instructions](#setup-instructions)
- [Building and Running the Service](#building-and-running-the-service)
- [Accessing the gRPC Service](#accessing-the-grpc-service)
- [gRPC Endpoints](#grpc-endpoints)
- [Running Tests](#running-tests)
- [Extra Features](#extra-features)
- [License](#license)
- [Acknowledgments](#acknowledgments)
## Project Structure
grpc-user-serviceo
├───cmd
│   └───server
│       └───main.go
├───configs
│   └───config.go
├───internal
│   ├───model
│   │   └───user.go
│   └───user
│       ├───delivery
│       │   └───grpc
│       │       └───user_handler.go
│       ├───repository
│       │   └───memory
│       │       └───user_repository.go
│       └───usecase
│           └───user_usecase.go
├───pkg
│   └───grpc
│       └───user
│           └───server.go
└───test
    └───user_service_test.go

## Prerequisites
. Go 1.19+
. Docker
. Docker Compose

# Setup Instructions

## Clone the Repository
git clone <your-repo-url>
cd grpc-user-serviceo

## Build and Run with Docker Compose
docker-compose up --build


# gRPC Endpoints
The gRPC server runs on port 50051

#   g r p c - u s e r - s e r v i c e o  
 