FROM golang:alpine AS builder

WORKDIR /builder

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./cmd/main ./cmd/main.go

FROM alpine

WORKDIR /app

COPY ./redis.env ./redis.env

COPY --from=builder /builder/cmd/main ./cmd/main

EXPOSE 8081

CMD ["./cmd/main"]
