# syntax=docker/dockerfile:1
FROM golang:1.19 AS builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend-app .

FROM alpine:latest

ENV mode prod

# Set the timezone to Asia/Bangkok
RUN apk update && apk add --no-cache tzdata && cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && echo "Asia/Bangkok" > /etc/timezone

# Copy the built binary from the previous stage
COPY --from=builder /app/backend-app /usr/local/bin/backend-app
COPY --from=builder /app/config /config

# Expose the port your application listens on
EXPOSE 8000

# Set the command to start your application
CMD ["/usr/local/bin/backend-app"]