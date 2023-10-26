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
