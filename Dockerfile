FROM alpine

ENV GOPATH=/go

RUN mkdir /go

RUN apk update && apk add go

COPY . /go/src/github.com/Xe/macarena

RUN go build -tags netgo github.com/Xe/macarena/...

ONBUILD COPY config.json /macarena/config.json
WORKDIR /macarena
CMD "/go/bin/macarena"
