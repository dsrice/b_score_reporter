FROM golang:1.20-alpine

RUN apk update && apk add git && \
    tzdata

WORKDIR /go
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.52.2

RUN mkdir ./login
WORKDIR /go/src/login

RUN  go install golang.org/x/tools/cmd/godoc@latest

