# Use the official Golang image as the base image
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory (if you have them)
COPY go.mod go.sum ./

# Download and cache Go dependencies (if you have them)
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Create a new lightweight image for the final executable
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the executable from the builder stage
COPY --from=builder /app/main .

# Set the command to run the executable
CMD ["./main"]
