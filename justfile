# build main
build:
    go run cmd/main.go

tidy:
    go mod tidy

make-migrations MIGRATION_NAME:
    atlas migrate diff {{MIGRATION_NAME}} \
        --dir "file://db/migrations" \
        --to "file://db/schema.hcl" \
        --dev-url "sqlite://dev?mode=memory"

migrate:
    atlas schema apply \
        --env turso \
        --to "file://db/schema.hcl"

test:
    atlas schema apply \
        --url "sqlite://test.db" \
        --to "file://db/schema.hcl" \
        --dev-url "sqlite://file?mode=memory" \
    && go test
