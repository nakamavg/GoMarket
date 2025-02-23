# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copiar módulos primero para caché
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código
COPY . .

# Build de la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copiar binario desde builder
COPY --from=builder /app/main .
COPY --from=builder /app/internal ./internal

EXPOSE 8080

# Comando de ejecución
CMD ["./main"]