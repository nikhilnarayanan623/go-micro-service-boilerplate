# Go Microservice Boilerplate

![Go Version](https://img.shields.io/badge/Go-1.17-blue)

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Instructions](#instructions)
    - [API Gateway](#api-gateway)
    - [Auth Service](#auth-service)
    - [Employee Service](#employee-service)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Technologies Used](#technologies-used)

## Introduction

This is a boilerplate for building microservices in Go. It provides a clean and structured codebase, along with commonly used libraries and tools for building scalable and maintainable microservices. It's designed with best practices in mind and includes features like gRPC, JWT authentication, configuration management with Viper, and more.

## Features

- Clean and modular code architecture
- gRPC support for efficient communication between microservices
- JWT-based authentication for secure API access
- Configuration management using Viper
- Web API using the Gin framework
- Dependency injection with Wire

## Getting Started

### Prerequisites

To use this boilerplate, you need to have the following installed on your system:

- Go (1.17 or later)
- PostgreSQL (as the database)

### Installation

  #### 1. Clone the repository:

  ```bash
    git clone https://github.com/nikhiloayaw/go-micro-service-boilerplate.git && \
    cd ./go-micro-service-boilerplate
  ```

### Instructions

  #### 1. API Gateway
  ##### 1. Install dependencies
  ```bash
    ## Assuming you are in root of the project
    cd ./api-gateway && \
    go mod tidy
  ```
  ##### 2. Setup Env
  create a .env file and add the below values
  ```.env
     ## api gateway
     API_PORT="port that you want to run api gateway service"
     ## auth service
     AUTH_SERVICE_HOST="auth service host"
     AUTH_SERVICE_PORT="auth service port"
     ## employee service
     EMPLOYEE_SERVICE_HOST="employee service host"
     EMPLOYEE_SERVICE_PORT="employee service port"
  ```
  ##### 3. Run Application
  ```bash
    go run ./cmd/api
  ```
  #### 2. Auth Service
  ##### 1. Install dependencies
  ```bash
    ## Assuming you are in root of the project
    cd ./auth-service && \
    go mod tidy
  ```
  ##### 2. Setup Env
  create a .env file and add the below values
  ```.env
    AUTH_SERVICE_HOST="auth service host"
    AUTH_SERVICE_PORT="auth service port"
    ## Database
    DB_HOST="database running host"
    DB_PORT="database running port"
    DB_NAME="database name"
    DB_USER="database user"
    DB_PASSWORD="database user password"
    ## JWT
    JWT_KEY="a key that you wan't to sign in for token"
  ```
  ##### 3. Run Application
  ```bash
    go run ./cmd/api
  ```
  #### 3. Employee Service
  ##### 1. Install dependencies
  ```bash
    cd ./employee-service && \
    go mod tidy
  ```
  ##### 2. Setup Env
  create a .env file and add the below values
  ```.env
    EMPLOYEE_SERVICE_HOST="employee service host"
    EMPLOYEE_SERVICE_PORT="employee service port"
  ```
  ##### 3. Run Application
  ```bash
    go run ./cmd/api
  ```

## Usage
### 1. Live API Documentation
If you are running the project then visit (http://localhost:{$API_PORT}/swagger/index.html)

### 2. Explanation of Services
##### 1. API Gateway:
The API Gateway is the entry point for all external requests to your microservices. It acts as a central hub for routing requests to the appropriate microservice.
It handles incoming HTTP requests and directs them to the relevant microservices based on the request path or other criteria.
The API Gateway is responsible for distributing requests to various services, making it a crucial component for scaling and managing your microservices.
It also provides a unified interface to clients, abstracting the underlying services and simplifying the API for consumers.
To run the API Gateway, you need to set up the environment variables in its .env file, including the ports and host information for the Auth and Employee services. Then, you can start the gateway using go run ./cmd/api for checkout the Instructions.

##### Auth Service:
The Auth Service is responsible for handling user authentication and authorization. It ensures secure access to your microservices.
It provides functionalities like user registration (sign up), user authentication (sign in), and generating JSON Web Tokens (JWT) for authorized users.
The service uses a PostgreSQL database to store user information, and you need to configure database connection details in its .env file.
The JWT_KEY in the .env file is used to sign and verify JWTs, which are crucial for secure communication and access control in your microservices.
To run the Auth Service, you need to install its dependencies, set up the environment variables, and then start the service using go run ./cmd/api for more checkout the Instructions.

##### Employee Service:
The Employee Service is responsible for generating random employee details. It can be used as a sample service for testing and demonstration purposes.
It can generate a specified number of random employee records on request, making it useful for simulating interactions with a database.
The Employee Service has its environment variables defined in its .env file, specifying the host and port where it will run.
To run the Employee Service, you need to install its dependencies, set up the environment variables, and then start the service using go run ./cmd/api for more checkout the Instructions.
