FROM golang:1.22-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/ordersystem

COPY ./cmd/ordersystem/.env .

EXPOSE 8080 8000 50051

CMD ["./wait-for-it.sh", "rabbitmq:5672", "--timeout=30", "--", "./wait-for-it.sh", "mysql:3306", "--timeout=30", "--", "./myapp"]
