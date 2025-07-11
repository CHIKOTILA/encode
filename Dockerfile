# Этап сборки
FROM golang:1.24 AS builder

WORKDIR /app

# Установим зависимости для cgo + sqlite3
RUN apt-get update && apt-get install -y build-essential libsqlite3-dev

# Копируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем проект
COPY . .

# Включаем cgo
ENV CGO_ENABLED=1

# Собираем бинарник
RUN go build -o person-crud main.go

# Финальный образ
FROM debian:bookworm-slim

WORKDIR /app

# Устанавливаем runtime-зависимости
RUN apt-get update && apt-get install -y libsqlite3-0 ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/person-crud .

EXPOSE 8080

CMD ["./person-crud"]