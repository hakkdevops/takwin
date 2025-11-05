# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o takwin main.go

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk --no-cache add ca-certificates gcc g++ make

# Create non-root user
RUN addgroup -g 1001 -S takwin && \
    adduser -u 1001 -S takwin -G takwin

# Set working directory
WORKDIR /workspace

# Copy binary from builder stage
COPY --from=builder /app/takwin /usr/local/bin/takwin

# Make binary executable
RUN chmod +x /usr/local/bin/takwin

# Change ownership
RUN chown takwin:takwin /usr/local/bin/takwin

# Switch to non-root user
USER takwin

# Set entrypoint
ENTRYPOINT ["takwin"]

# Default command
CMD ["--help"]

# Labels
LABEL org.opencontainers.image.title="Takwin"
LABEL org.opencontainers.image.description="Modern build tool for C/C++ projects"
LABEL org.opencontainers.image.vendor="Hakkim"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/hakkim/takwin-go"
LABEL org.opencontainers.image.documentation="https://github.com/hakkim/takwin-go/docs"