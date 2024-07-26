set dotenv-filename := ".env.local"

build:
    docker compose -f compose.local.yaml up

tidy:
    go mod tidy

make-migrations MIGRATION_NAME:
    atlas migrate diff {{MIGRATION_NAME}} \
        --dir "file://db/migrations" \
        --to "file://db/schema.hcl" \
        --dev-url "docker://postgres/16/dev?search_path=public"

migrate:
    atlas migrate apply \
        --dir "file://db/migrations" \
        --url "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?search_path=public&sslmode=disable"

test:
    go test -v ./...
