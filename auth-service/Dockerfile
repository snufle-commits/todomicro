FROM golang:1.25-alpine AS build

WORKDIR /app

# Копируем go.mod и скачиваем зависимости
COPY go.mod ./
RUN go mod download

# Копируем остальной код
COPY . .

# Собираем бинарник из папки cmd
RUN go build -o auth ./cmd

# Финальный минимальный образ
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/auth .
EXPOSE 8080
CMD ["./auth"]
