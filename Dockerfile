FROM golang as build

ADD . /usr/local/go/src/dbks

WORKDIR /usr/local/go/src/dbks

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dbks_server

