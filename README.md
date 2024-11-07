# GoManage

<div>
  <img align="right" src="./gopher.png" height="100" alt="Gopher">
  <p align="justify">
    GoManage is a robust Go-based REST API designed to facilitate project management with a comprehensive and intuitive interface. This API allows users to create and manage projects, organize tasks into lists, and track progress efficiently. Leveraging an SQLite database for persistent storage, GoManage encompasses all essential CRUD (Create, Read, Update, Delete) operations for projects, lists, and tasks. The project adheres to best practices in Go development, ensuring clean, maintainable code, comprehensive integration tests, and a streamlined build process facilitated by a Makefile.
  </p>
</div>

Developed with Go 1.22, utilising the updated ["net/http"](https://go.dev/blog/routing-enhancements) package routing enhancements.

## Features

- **Project Management**
  - Create, update, delete, and retrieve projects
  - API versioning for scalable and maintainable endpoints
- **List Management**
  - Organize tasks within lists specific to projects
  - Comprehensive CRUD operations for lists
- **Task Management**
  - Create, update, delete, and retrieve tasks within lists and projects
  - Mark tasks as done or undone
- **Routing**
  - Utilizes the latest "net/http" package enhancements for Go 1.22
- **Database Support**
  - SQLite for lightweight, file-based storage
  - Future support planned for PostgreSQL
- **Testing**
  - Integration tests covering all endpoints to ensure reliability
- **Validation**
  - Request validation using [go-playground/validator](https://github.com/go-playground/validator)
- **Documentation**
  - Comprehensive API documentation with Swagger
  - GoDoc generated documentation for code references
- **Build Automation**
  - Makefile for easy building, testing, and deployment
- **ORM**
  - [GORM](https://gorm.io/) for streamlined database interactions

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

## API Documentation

GoManage provides comprehensive API documentation via Swagger. To access the Swagger UI, ensure the application is running and navigate to:

```
http://localhost:8080/swagger
```

Alternatively, refer to the [Swagger YAML file](./docs/swagger.yaml) for detailed endpoint specifications.

### Available Endpoints

#### Projects

- **Get all projects**
  - `GET /api/v1/projects`
- **Create a new project**
  - `POST /api/v1/projects`
- **Update a project**
  - `PUT /api/v1/projects/{id}`
- **Delete a project**
  - `DELETE /api/v1/projects/{id}`

#### Lists

- **Get all lists for a project**
  - `GET /api/v1/projects/{projectID}/lists`
- **Create a new list within a project**
  - `POST /api/v1/projects/{projectID}/lists`
- **Get a specific list within a project**
  - `GET /api/v1/projects/{projectID}/lists/{id}`
- **Update a list within a project**
  - `PUT /api/v1/projects/{projectID}/lists/{id}`
- **Delete a list within a project**
  - `DELETE /api/v1/projects/{projectID}/lists/{id}`

#### Tasks

- **Get all tasks for a list within a project**
  - `GET /api/v1/projects/{projectID}/lists/{listID}/tasks`
- **Create a new task for a list within a project**
  - `POST /api/v1/projects/{projectID}/lists/{listID}/tasks`
- **Update a task within a list and project**
  - `PUT /api/v1/projects/{projectID}/lists/{listID}/tasks/{taskID}`
- **Delete a task within a list and project**
  - `DELETE /api/v1/projects/{projectID}/lists/{listID}/tasks/{taskID}`
- **Mark a task as done**
  - `PATCH /api/v1/projects/{projectID}/lists/{listID}/tasks/{taskID}/done`
- **Mark a task as undone**
  - `PATCH /api/v1/projects/{projectID}/lists/{listID}/tasks/{taskID}/undone`

For detailed information on request and response schemas, refer to the [Swagger YAML](./docs/swagger.yaml).

## Used Tools

This project utilises the following technologies and libraries:

- [Golang](https://golang.org/) for backend development
- [net/http](https://go.dev/blog/routing-enhancements) for route management
- [GORM](https://gorm.io/) for ORM and database interactions
- [SQLite](https://www.sqlite.org/index.html) as the primary database
- [Docker](https://www.docker.com/) for containerized database management
- [go-playground/validator](https://github.com/go-playground/validator) for input validation
- [Swagger](https://swagger.io/) for API documentation

## To-dos

- [x] Utilise "net/http" package routing enhancements
- [x] API versioning
- [x] GORM integration
- [x] Set up for Test Driven Development
- [x] Projects
  - [x] CRUD endpoints
  - [x] Integration Tests
- [x] Lists
  - [x] CRUD endpoints
  - [x] Integration Tests
- [x] Tasks
  - [x] CRUD endpoints
  - [x] Mark as done/undone
  - [x] Integration Tests
- [x] Database support
  - [x] SQLite
  - [ ] PostgreSQL
- [x] Add go-playground/validator for request validation
- [x] Add Swagger for API documentation
- [x] Add GoDoc for code documentation
- [ ] Add a custom Logger package
