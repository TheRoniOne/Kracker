FROM denoland/deno:alpine-2.1.4 AS base

COPY . /app

WORKDIR /app

RUN deno run build

FROM denoland/deno:alpine-2.1.4 AS prod

COPY --from=base /app/dist /app

WORKDIR /app

CMD ["deno", "--allow-env", "--allow-read", "--allow-net", "--allow-sys", "main.ts"]
