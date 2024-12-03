# Stage 1: Build the application
FROM golang:1.20 as builder

# Set the working directory
WORKDIR /app

# Copy the Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the application
RUN go build -o main .

# Stage 2: Create a lightweight production image
FROM alpine:latest

# Install CA certificates for HTTPS support
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the built application from the builder stage
COPY --from=builder /app/main .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
