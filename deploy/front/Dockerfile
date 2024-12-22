FROM denoland/deno:alpine-2.1.4 AS builder

COPY ./front /app

WORKDIR /app

RUN deno run build

FROM denoland/deno:alpine-2.1.4 AS prod

COPY --from=builder /app/dist /app

WORKDIR /app

EXPOSE 4321

CMD ["deno", "--allow-env", "--allow-read", "--allow-net", "--allow-sys", "main.ts"]
