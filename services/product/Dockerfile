# Build Stage
FROM golang:1.22.3 AS builder
WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN go build -o /app/go-microservices-products ./

# Default command to run the binary
EXPOSE 8081
CMD ["./go-microservices-products"]