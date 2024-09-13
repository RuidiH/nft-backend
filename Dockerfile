# Use the official Go image as the build environment
FROM golang:1.21.6-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download dependencies (if any)
RUN go mod download

# Copy the source code
COPY cmd/ ./cmd/

# Build the Go application
RUN go build -o main ./cmd/app

# Use a minimal base image for the final build
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Command to run when starting the container
CMD ["./main"]

EXPOSE 8080