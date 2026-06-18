# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git for go mod download (if private modules are needed later).
RUN apk add --no-cache git

# Copy dependency files first for better layer caching.
COPY go.mod ./
RUN go mod download

# Copy source code and build the binary.
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/server

# Runtime stage
FROM alpine:3.21

# Install curl for the HEALTHCHECK.
RUN apk add --no-cache curl

# Create a non-root user.
RUN addgroup -g 1000 appgroup && adduser -u 1000 -G appgroup -s /bin/sh -D appuser

WORKDIR /app

# Copy the compiled binary and application assets.
COPY --from=builder /app/server .
COPY --from=builder /app/web ./web

# Set ownership to the non-root user.
RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

# Health check used by Coolify for zero-downtime deployments.
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
    CMD curl -f http://localhost:8080/healthz || exit 1

CMD ["./server"]
