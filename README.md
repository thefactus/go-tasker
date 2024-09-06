# GoTasker

<div>
  <img align="right" src="./gopher.png" height="100" alt="Gopher">
  <p align="justify">
    GoTasker is a Go-based REST API that helps users manage tasks with a simple, intuitive interface. It features an SQLite database for storing tasks and includes all the essential functions for creating, updating, and deleting tasks. The project follows best practices in Go development, with clean, easy-to-maintain code. It also includes integration tests and a Makefile to simplify building and testing.
  </p>
</div>

Developed with Go 1.22, utilising the updated ["net/http"](https://go.dev/blog/routing-enhancements) package routing enhancements.

## Features

- Routing with the latest "net/http" package enhancements for Go 1.22
- Endpoints for managing lists
- Endpoints for managing tasks
- Integration tests for all endpoints
- SQLite database support
- Makefile for easy building and testing
- GORM for database management
- Go-playground/validator for request validation

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests

```bash
make all build
```

build the application

```bash
make build
```

run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB container

```bash
make docker-down
```

live reload the application

```bash
make watch
```

run the test suite

```bash
make test
```

clean up binary from the last build

```bash
make clean
```

## Used Tools

This project uses the following tools:

- [Golang](https://golang.org/) for backend development
- [net/http](https://go.dev/blog/routing-enhancements) for route management
- [GORM](https://gorm.io/) for database communication
- [SQLite](https://www.sqlite.org/index.html) as the database

## To-dos

- [x] utilise "net/http" package routing enhancements
- [x] API versioning
- [x] GORM
- [x] Set up for Test Driven Development
- [x] Lists
  - [x] Integration Tests
- [x] Tasks
  - [x] Integration Tests
- [ ] Databases support
  - [x] SQLite
  - [ ] PostgreSQL
- [x] Add go-playground/validator
- [ ] Add a custom Logger package
- [ ] Add Swagger
- [x] Add Go Docs
