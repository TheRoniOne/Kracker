FROM golang:1.23

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY back/go.mod back/go.sum ./
RUN go mod download && go mod verify

COPY back/. .
RUN mkdir -p /app/bin && go build -v -o /app/bin ./...

EXPOSE 1323

CMD ["bin/cmd"]
