FROM alpine

ENV GOPATH /go
ENV DOCKER yes

RUN mkdir /go

RUN apk update && apk add go ca-certificates

COPY . /go/src/github.com/Xe/macarena
COPY ./build/run.sh /macarena/run.sh

RUN go get github.com/Xe/macarena

ONBUILD COPY config.json /macarena/config.json
WORKDIR /macarena
CMD "/macarena/run.sh"
