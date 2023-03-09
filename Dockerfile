FROM golang:alpine as builder

WORKDIR /go/src/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="wuda@wuda.com"

WORKDIR /go/src/server

COPY --from=0 /go/src/server ./

EXPOSE 8888
ENTRYPOINT ./server