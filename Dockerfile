FROM golang:1.18-alpine
WORKDIR /app

# Copy the current directory contents into the container at /app

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy
COPY . .

# Build the Golang application
RUN go build ./cmd/server

# Expose port 8080
EXPOSE 8080

# Set the entrypoint to run the Golang application
CMD ["go", "run", "./cmd/server"]

FROM ubuntu:16.04

# Install prerequisites
RUN apt-get update && apt-get install -y \
curl
CMD /bin/bash