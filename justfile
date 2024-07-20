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
    atlas migrate apply \
        --env local

test:
    atlas migrate apply \
        --url "sqlite://test.db" \
        --to "file://db/schema.hcl" \
        --dev-url "docker://postgres/16/dev?search_path=public" \
    && go test
