FROM golang:1.25-alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api-service ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /api-service /app/.env ./

EXPOSE 8080

CMD ["./api-service"]