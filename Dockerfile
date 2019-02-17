FROM golang:1.11-alpine
MAINTAINER  wj

RUN apk add --no-cache git && go get github.com/jason-wj/bitesla-strategy-exec