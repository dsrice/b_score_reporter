FROM golang:1.20-alpine

RUN apk update && apk add git && \
    tzdata

RUN mkdir ./login
WORKDIR /go/src/login