set dotenv-filename := ".env.local"

build:
    docker compose -f compose.local.yaml up --build

tidy:
    go mod tidy

make-migrations MIGRATION_NAME:
    atlas migrate diff {{MIGRATION_NAME}} \
        --dir "file://db/migrations" \
        --to "file://db/schema.sql" \
        --dev-url "docker://postgres/17/dev?search_path=public" \
        && sqlc generate

migrate:
    atlas migrate apply \
        --dir "file://db/migrations" \
        --url "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?search_path=public&sslmode=disable"

test:
    go test -v ./...

lint:
    golangci-lint run
