FROM golang:1.25-alpine AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o todo ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/todo .
EXPOSE 8081
CMD ["./todo"]
