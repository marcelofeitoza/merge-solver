# Use the official Golang image to create a build artifact.
FROM golang:1.22 as builder

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /app

# Copy the go.mod and go.sum files first to leverage Docker cache.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code.
COPY . .

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o merge-solver-api cmd/app/main.go

# Use the alpine image for a small, final production image.
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /app/merge-solver-api .

# Expose port 8080 to the outside world.
EXPOSE 8080

# Command to run the executable.
CMD ["./merge-solver-api"]
