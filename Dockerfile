# Use an official Go runtime as the base image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local source files to the container's working directory
COPY . .

# Build the Go program
RUN go build -o wait-for-port

# Use a lightweight Alpine Linux image as the base
FROM alpine:3.14

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the previous build stage
COPY --from=0 /app/wait-for-port .

# Run the Go program when the container starts
ENTRYPOINT  ["/app/wait-for-port"]