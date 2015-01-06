FROM golang:latest
RUN mkdir -p /go/src/github.com/dcoxall
ADD . /go/src/github.com/dcoxall/financier
WORKDIR /go/src/github.com/dcoxall/financier
