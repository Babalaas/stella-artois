# source: https://codefresh.io/docs/docs/learn-by-example/golang/golang-hello-world/
FROM golang:1.19-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/stella-artois

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/stella-artois .

# Start fresh from a smaller image
FROM alpine:latest 
RUN apk add ca-certificates

COPY --from=build_base /tmp/stella-artois/out/stella-artois /app/stella-artois

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Environment Variables

# Run the binary program produced by `go install`
CMD ["/app/stella-artois"]