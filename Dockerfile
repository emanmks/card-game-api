FROM golang:1.19-alpine AS base

WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=1

# System dependencies
RUN apk update \
    && apk add --no-cache \
        ca-certificates \
        git \
    && apk upgrade \
    && update-ca-certificates

### Development with hot reload and debugger
FROM base AS dev
WORKDIR /app

ENV GOFLAGS="-buildvcs=false"

# Hot reloading mod
RUN go install github.com/cosmtrek/air@latest
RUN apk add --no-cache gcc g++ sqlite

EXPOSE 8080
EXPOSE 2345

ENTRYPOINT ["air"]

### CI
FROM base AS ci
WORKDIR /app

ENV GOFLAGS="-buildvcs=false"

RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN apk add --no-cache gcc g++ sqlite

CMD ["go"]

### Executable builder
FROM base AS builder
WORKDIR /app

ENV GOFLAGS="-buildvcs=false"
RUN apk add --no-cache gcc g++ sqlite

# Application dependencies
COPY . /app
RUN go mod download \
    && go mod verify

RUN go build -o card-game-api -a .

### Production
FROM alpine:latest

ENV GIN_MODE=release

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    && update-ca-certificates

# Copy executable
COPY --from=builder /app/card-game-api /usr/local/bin/card-game-api
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/card-game-api"]