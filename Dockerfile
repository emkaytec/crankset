# Build stage
FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /go/bin/app ./src

# Deployment stage
FROM scratch

# Copy the Go binary from the builder stage
COPY --from=builder /go/bin/app /app

# Command to run the binary
ENTRYPOINT ["/app"]
