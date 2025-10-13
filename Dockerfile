# Stage 1: Build Go binary
FROM golang:1.25-alpine AS builder
WORKDIR /app

# Install git (needed for Go modules sometimes)
RUN apk add --no-cache git

# Copy source code
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o myapp ./...

# Stage 2: Minimal runtime image
FROM alpine:latest
WORKDIR /app

# Copy binary
COPY --from=builder /app/myapp .

# Optional: expose port if your app listens
# EXPOSE 8080

# Run the binary
CMD ["./myapp"]