# Preparing the build environment
FROM golang:1.20-alpine AS builder

ENV DOCKER_BUILDKIT=1

# Working Directory
WORKDIR /workspace

# Install git, make, gcc, musl-de, because go alpine version doesn't include these
RUN apk add --no-cache git make gcc musl-dev

# Copy source code to workspace
COPY . .

# Build the binary and reduce the binary size
RUN go build -ldflags="-w -s" -o service.app

# Using latest alpine image
FROM alpine:latest

# Change working directory to /app
WORKDIR /app/

# Copy the binary from the builder image to the final image
COPY --from=builder /workspace/service.app /app/service.app

# Export port
EXPOSE 5001

# Set the permission and create a mountpoint
RUN mkdir /datastore && \
    chown -R 1001:0 /datastore /app/service.app && \
	chmod -R g+rwX /datastore

# Fix the permission issue
USER 1001

# Run the binary
ENTRYPOINT ["/app/service.app"]
