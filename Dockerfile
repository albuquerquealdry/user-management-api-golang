# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download -x

# Copy the application source code
COPY src ./src

# Build the Go application
RUN ls /app
RUN go build -o user-management-api /app/src/cmd/main.go


# Stage 2: Create a lightweight runtime image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/user-management-api ./

# Expose the port your application listens on (adjust if needed)
EXPOSE 8080

# Command to run the application
CMD ["./user-management-api"]