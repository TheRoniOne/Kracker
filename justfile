build:
    docker build --pull -f deploy/Dockerfile -t kracker . \
    && docker run -it --rm --name kracker-instance kracker

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
    go test -v ./...
