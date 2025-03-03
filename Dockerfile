FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o curlie

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/curlie .
COPY --from=builder /app/internal/infrastructure/ui ./internal/infrastructure/ui

EXPOSE 8080
CMD ["./curlie"] 