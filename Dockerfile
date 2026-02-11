# Compilacion
FROM golang:alpine AS builder

# Certificados (necesarios para conectar con AWS SES y Google Cloud)
RUN apk add --no-cache ca-certificates git

WORKDIR /app

# Copiamos archivos de dependencias primero para aprovechar la caché de Docker
COPY go.mod go.sum ./
RUN go mod download

# Copiamos el resto
COPY . .

# Compilamos el binario (estático para que corra en alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# Imagen de ejecución ligera
FROM alpine:latest

RUN apk add --no-cache ca-certificates
WORKDIR /root/

# Copiamos el binario desde la etapa anterior
COPY --from=builder /app/main .
# Copiamos la carpeta de logs por si acaso, aunque se creará sola
RUN mkdir logs

# No exponemos puertos porque es un Worker que consume de una cola, no una API
CMD ["./main"]