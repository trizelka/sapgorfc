# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.20-bullseye AS builder

# Set the Current Working Directory inside the container
WORKDIR /build

# Copy go mod and sum files
COPY go.mod go.sum ./
COPY gorfc/lib/* /usr/lib

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=1 GOOS=linux go build -o sapgorfc ./cmd/sapgorfc/main.go

FROM debian:bullseye-slim

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apt-get -y update && apt-get -y upgrade && \
    apt-get -y install bash git ssh

# Add Maintainer Info
LABEL maintainer="Trisia Juniarto <trizelka@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

COPY --from=builder /build/sapgorfc   ./
COPY gorfc/lib/* /usr/lib
COPY config.json ./
COPY .env ./

# Expose port 8088 to the outside world
EXPOSE 8089

# Run the executable
CMD ["./sapgorfc"]
