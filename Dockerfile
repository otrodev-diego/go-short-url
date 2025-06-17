FROM golang:1.24-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o shortener ./cmd/shortener

# Imagen
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/shortener .

# Expone el puerto usado por el servidor en este caso 8081
EXPOSE 8081

# Comando por defecto al ejecutar el contenedor
CMD ["./shortener"]