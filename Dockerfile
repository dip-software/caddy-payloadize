# The Context must be the ./Source folder, otherwise monitizer won't be found
FROM golang:1.25 AS builder
WORKDIR /caddy

# Build
COPY . /caddy

# Pre-download dependencies with our pinned versions to populate the module cache
RUN go mod download

RUN go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest

# Build Caddy with explicit version and replace directives to pin incompatible dependencies
# This fixes the incompatibility between certificates v0.28.4 and nebula v1.10.0
RUN /go/bin/xcaddy build v2.10.2 \
    --with github.com/loafoe/caddy-token \
    --with payloadsize=/caddy \
    --replace github.com/slackhq/nebula=github.com/slackhq/nebula@v1.6.1 \
    --replace github.com/smallstep/certificates=github.com/smallstep/certificates@v0.26.1

FROM alpine:latest
USER root
COPY --from=builder /caddy/caddy /usr/bin
