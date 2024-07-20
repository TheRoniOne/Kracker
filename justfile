# build main
build:
    go run cmd/main.go

tidy:
    go mod tidy

make-migrations MIGRATION_NAME:
    atlas migrate diff {{MIGRATION_NAME}} \
        --dir "file://db/migrations" \
        --to "file://db/schema.hcl" \
        --dev-url "docker://postgres/16/dev?search_path=public"

migrate:
    atlas schema apply \
        --env turso \
        --to "file://db/schema.hcl"

test:
    atlas schema apply \
        --url "sqlite://test.db" \
        --to "file://db/schema.hcl" \
        --dev-url "docker://postgres/16/dev?search_path=public" \
    && go test
