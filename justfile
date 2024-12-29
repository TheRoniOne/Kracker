set dotenv-filename := "back/.env.local"

build:
    docker build --pull -t kracker-backend:local -f deploy/back/Dockerfile ./back \
    && docker build --pull -t kracker-frontend:local -f deploy/front/Dockerfile ./front \
    && DOMAIN=localhost docker compose -f compose.local.yaml up

tidy:
    go mod tidy

make-migrations MIGRATION_NAME:
    atlas migrate diff {{MIGRATION_NAME}} \
        --dir "file://back/db/migrations" \
        --to "file://back/db/schema.sql" \
        --dev-url "docker://postgres/17/dev?search_path=public" \
        && sqlc generate

migrate:
    atlas migrate apply \
        --dir "file://back/db/migrations" \
        --url "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?search_path=public&sslmode=disable"

test-back:
    go test -v ./...

lint-back:
    golangci-lint run

tag-and-push VERSION:
    git tag v{{VERSION}} && git push origin v{{VERSION}}
