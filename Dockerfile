# Use an official Golang runtime as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY go.mod ./

# install them
RUN go mod download

# 
COPY ./ ./

# Build the Go application
RUN go build -o api

# Expose a port for the application to listen on (adjust as needed)
EXPOSE 8080

# Define the command to run your application
CMD ["./api"]
