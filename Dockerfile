FROM golang:alpine as builder

ENV CGO_ENABLED=1

WORKDIR /app
COPY . /app
RUN apk update && apk add gcc librdkafka-dev openssl-libs-static zlib-static zstd-libs libsasl librdkafka-static lz4-dev lz4-static zstd-static libc-dev musl-dev

RUN go mod download
RUN go build -tags musl -ldflags '-w -extldflags "-static"' -o /usr/local/bin/split-entrypoints ./cmd/main.go

FROM scratch

COPY --from=0 /usr/local/bin/split-entrypoints .

ENTRYPOINT ["/split-entrypoints"]