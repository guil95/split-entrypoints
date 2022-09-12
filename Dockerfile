FROM golang:buster as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE='on'

WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o /usr/local/bin/split-entrypoints ./cmd/main.go

FROM scratch

COPY --from=0 /usr/local/bin/split-entrypoints .

ENTRYPOINT ["/split-entrypoints"]