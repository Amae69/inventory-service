# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod ./
# COPY go.sum ./ # No dependencies yet, so go.sum might not exist
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o inventory-service .

# Final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/inventory-service .
EXPOSE 8080
CMD ["./inventory-service"]
