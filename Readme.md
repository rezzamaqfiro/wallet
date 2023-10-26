# Backend for Wallet

## Pre-requisites
1. Install [sqlc](https://github.com/sqlc-dev/sqlc/blob/main/docs/overview/install.md)

    Quick install command:
    ```
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    ```

## Development

1. Start database server: `docker-compose up -d`
2. Copy `config.example.yaml` to `config.yaml`
3. Fetch Dependency: `go mod tidy`
4. Vendoring Dependency: `go mod vendor`
5. Generate sqlc: `make generate-queries`
6. Update database: `go run main.go migrate`
7. Start server: `go run main.go start`

Whenever there is changes in queries, re-run step number 5.

Whenever there is changes in db_schema, re-run step number 5 and the re-run step number 6.

## Docker Tools

- [Adminer (Database Client)](http://localhost:8080)