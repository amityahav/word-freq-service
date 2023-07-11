# builder
FROM golang:1.20-alpine AS builder

RUN set -ex &&\
    apk add --no-progress --no-cache \
      gcc \
      musl-dev \
      git
ENV GO111MODULE=on

WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -buildvcs=false -o freq ./cmd

# image
FROM alpine:edge
WORKDIR /
COPY --from=builder /app .
ENTRYPOINT ["/freq", "--config", "./configs/config.yaml"]
