# Use official Golang image for building
FROM golang:1.22.5 AS builder
WORKDIR /app

# Copy go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all source files and build the Go app
COPY . .
RUN go build -o server server.go && chmod +x server

# Use a lightweight image for running the application
FROM alpine:latest
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

# Ensure execution permissions
RUN chmod +x server

# Expose the port the server runs on
EXPOSE 8080

# Run the compiled Go application
CMD ["./server"]

