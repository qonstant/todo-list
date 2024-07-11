make migrateup and migratedown for migration

sqlc for Database
brew install sqlc
or
go get github.com/kyleconroy/sqlc/cmd/sqlc

with "make test-html" u can see coverage of tests

for swagger: swag init -g internal/handlers/http/task.go

For running app, for docker compose up -d
make up

and make down
for docker compose down

# Todo List App

## Goal

Creating a RESTful API for a Todo List microservice. The service is deployed on Render, containerized using Docker, and run with Docker Compose and Makefile. Input data validated, and the README.md file is comprehensive.

## Features
- **Database was created by using dbdiagram.**

  *Photo*

- **CRUD operations for managing tasks.**
  - CRUD was created using SQLC. For regeneration of CRUD:
    ```bash
    make sqlc
    ```

- **Swagger UI for API documentation.**
  - For regeneration of Swagger documentation:
    ```bash
    swag init -g internal/handlers/http/task.go
    ```

- **Docker support for containerization.**
  - For running up:
    ```bash
    make up
    ```
  - For shutting down:
    ```bash
    make down
    ```
  - For restart:
    ```bash
    make restart
    ```

- **Unit tests for validating the CRUD.**
  - For running unit tests with visualization:
    ```bash
    make test-html
    ```

## Publishing and Deployment
- Project published on GitHub.
- Project deployed on Render.
- Project containerized using Docker.
- Project runs using Docker Compose and Makefile.

## Prerequisites

- Go 1.22 or later
- Docker (for containerization)
- SQlC (for CRUD generation)
```bash
brew install sqlc
```
or
```bash
go get github.com/kyleconroy/sqlc/cmd/sqlc
```

## Getting Started

### Clone the Repository

```bash
https://github.com/qonstant/go-proxy.git
cd go-proxy
```
## Build and Run Locally

### Build the application:

```bash
make build
```

### Run the application:

```bash
make run
```
After running make run, the server will start and be accessible at http://localhost:8080.

Link to the deployment: https://go-proxy-1fo6.onrender.com/proxy

### Generate Swagger Documentation

```bash
make swagger
```
### Run Tests
```bash
make test
```
## Docker
### For docker compose up
```bash
make up
```
### For docker compose down
```bash
make down
```
### For restarting container
```bash
make restart
```

## API Endpoints

### Proxy Request

- URL: /proxy
- Method: POST
- Request Body:
```json
{
    "method": "GET",
    "url": "http://example.com",
    "headers": {
        "Content-Type": "application/json"
    }
}
```

For instance:
```json
{
    "method": "GET",
    "url": "http://jsonplaceholder.typicode.com/posts/1",
    "headers": {
        "Content-Type": "application/json"
    }
}
```
- Response:
```json
{
    "id": "1627563890765102000",
    "status": 200,
    "headers": {
        "Content-Type": "application/json"
    },
    "length": 1270
}
```

From the example above:
```json
{
    "id": "d7f4fb3d-6def-4378-857e-91711eb018c6",
    "status": 200,
    "headers": {
        "Access-Control-Allow-Credentials": "true",
        "Alt-Svc": "h3=\":443\"; ma=86400",
        "Cache-Control": "max-age=43200",
        "Cf-Cache-Status": "REVALIDATED",
        "Cf-Ray": "89db0c3a8d6d5d50-FRA",
        "Connection": "keep-alive",
        "Content-Type": "application/json; charset=utf-8",
        "Date": "Thu, 04 Jul 2024 00:37:37 GMT",
        "Etag": "W/\"124-yiKdLzqO5gfBrJFrcdJ8Yq0LGnU\"",
        "Expires": "-1",
        "Nel": "{\"report_to\":\"heroku-nel\",\"max_age\":3600,\"success_fraction\":0.005,\"failure_fraction\":0.05,\"response_headers\":[\"Via\"]}",
        "Pragma": "no-cache",
        "Report-To": "{\"group\":\"heroku-nel\",\"max_age\":3600,\"endpoints\":[{\"url\":\"https://nel.heroku.com/reports?ts=1719290587&sid=e11707d5-02a7-43ef-b45e-2cf4d2036f7d&s=eAlPGj2psKwqFTi3aRIeAycEDJsdhwHLI%2F0cXgblPNM%3D\"}]}",
        "Reporting-Endpoints": "heroku-nel=https://nel.heroku.com/reports?ts=1719290587&sid=e11707d5-02a7-43ef-b45e-2cf4d2036f7d&s=eAlPGj2psKwqFTi3aRIeAycEDJsdhwHLI%2F0cXgblPNM%3D",
        "Server": "cloudflare",
        "Vary": "Origin, Accept-Encoding",
        "Via": "1.1 vegur",
        "X-Content-Type-Options": "nosniff",
        "X-Powered-By": "Express",
        "X-Ratelimit-Limit": "1000",
        "X-Ratelimit-Remaining": "999",
        "X-Ratelimit-Reset": "1719290646"
    },
    "length": 292,
    "body": "{\n  \"userId\": 1,\n  \"id\": 1,\n  \"title\": \"sunt aut facere repellat provident occaecati excepturi optio reprehenderit\",\n  \"body\": \"quia et suscipit\\nsuscipit recusandae consequuntur expedita et cum\\nreprehenderit molestiae ut ut quas totam\\nnostrum rerum est autem sunt rem eveniet architecto\"\n}"
}
```

### Swagger Documentation

- URL: /swagger/

- Link: https://go-proxy-1fo6.onrender.com/swagger/index.html

# Swagger: HTTP tutorial for beginners

1. Add comments to your API source code, See [Declarative Comments Format](#declarative-comments-format).

2. Download swag by using:
```sh
go install github.com/swaggo/swag/cmd/swag@latest
```
To build from source you need [Go](https://golang.org/dl/) (1.17 or newer).

Or download a pre-compiled binary from the [release page](https://github.com/swaggo/swag/releases).

3. Run `swag init` in the project's root folder which contains the `main.go` file. This will parse your comments and generate the required files (`docs` folder and `docs/docs.go`).
```sh
swag init
```

  Make sure to import the generated `docs/docs.go` so that your specific configuration gets `init`'ed. If your General API annotations do not live in `main.go`, you can let swag know with `-g` flag.
  ```sh
  swag init -g internal/handler/handler.go
  ```

4. (optional) Use `swag fmt` format the SWAG comment. (Please upgrade to the latest version)

  ```sh
  swag fmt
  ```

## Project Structure

- main.go: The main server implementation.
- main_test.go: Unit tests for the proxy handler.
- Makefile: Makefile for building, running, testing, and Docker tasks.
- Dockerfile: Dockerfile for containerizing the application.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any improvements.