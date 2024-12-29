set dotenv-filename := "back/.env.local"

build:
    docker build --pull -t kracker-backend:local -f deploy/back/Dockerfile ./back \
    && docker build --pull -t kracker-frontend:local -f deploy/front/Dockerfile ./front \
    && DOMAIN=localhost KRACKER_VERSION=local docker compose -f compose.local.yml up

tag-and-push VERSION:
    git tag v{{VERSION}} && git push origin v{{VERSION}}
