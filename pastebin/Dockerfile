FROM golang:alpine AS builder

WORKDIR /builder

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN go build -o ./cmd/main ./cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /builder/cmd/main ./cmd/main
COPY --from=builder /builder/templates ./templates
COPY --from=builder /builder/.env ./.env

EXPOSE 8080

CMD ["./cmd/main"]
