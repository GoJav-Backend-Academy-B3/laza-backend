FROM golang:1.20-alpine AS builder
RUN apk update && apk add --no-cache git
ARG VERSION=dev

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main -ldflags=-X=main.version=${VERSION} cmd/server/main.go

FROM golang:1.20-alpine
WORKDIR /server
COPY --from=builder /app/main /server/main

EXPOSE 8080
ENTRYPOINT ["/server/main"]