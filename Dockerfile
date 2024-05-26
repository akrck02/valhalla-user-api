
# Download base image ubuntu 22.04
FROM golang:1.22.3-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && apk add --no-cache bash git openssh

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download
run go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o app .

# This container exposes ports to the outside world
EXPOSE 3333

# Run the executable
CMD ["./app"]