# Building RESTful API with Hexagonal Architecture in Go

[Building RESTful API with Hexagonal Architecture in Go](https://dev.to/bagashiz/building-restful-api-with-hexagonal-architecture-in-go-1mij)

## Getting Started

1. Ensure you have [Go](https://go.dev/dl/) 1.21 or higher and [Task](https://taskfile.dev/installation/) installed on your machine:

    ```bash
    go version && task --version
    ```

2. Install all required tools for the project:

    ```bash
    task install
    ```

3. Create a copy of the `.env.example` file and rename it to `.env`:

    ```bash
    cp .env.example .env
    ```

    Update configuration values as needed.

4. Run the service containers:

    ```bash
    task service:up
    task migrate:up
    ```

    **NOTE**: the command use `podman` and `podman-compose` by default. If you want to use `docker` or `docker compose`, manually run `docker` commands instead or replace `podman` and `podman-compose` with `docker` or `docker compose` in the tasks inside `Taskfile.yml` file.

5. Run the project in development mode:

    ```bash
    task dev
    ```
