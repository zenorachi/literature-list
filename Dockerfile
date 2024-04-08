# Stage 1: Build the Go application with dlv for debugging
FROM golang:alpine AS build

WORKDIR /app

COPY ./ ./

# Build application
RUN go build ./cmd/literature-list/

# Stage 2: Create the final Docker image
FROM alpine:latest

WORKDIR /app

# Copy binary from Stage 1
COPY --from=build /app/literature-list /app/

# Copy configs from Stage 1
COPY --from=build /app/configs/main.yml /app/configs/main.yml

CMD ["./literature-list"]