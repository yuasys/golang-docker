FROM golang:1.16.5

RUN mkdir /go/src/app

WORKDIR /go/src/app

ADD . /go/src/app/
