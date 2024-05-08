# Todolist

<p align="right">
  <img src="./gopher.png" alt="Gopher">
</p>

This project is a Go-based web application that provides a simple and intuitive interface for managing to-do lists. It features a robust backend with a SQLite database, and includes functionalities for dealing with to-do lists. The project also includes a comprehensive set of unit tests and a Makefile for easy building and testing. It's designed with a focus on clean, maintainable code and follows best practices for Go project structure.

## To-dos

- [x] GORM
- [x] Lists
  - [x] Integration Tests
- [ ] Tasks

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
