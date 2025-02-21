# Build stage: use the official Golang image
FROM golang:1.19 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application (the output binary is "main")
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Run stage: use a minimal image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Start the application
CMD ["./main"]
