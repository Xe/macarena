FROM yikaus/alpine-base

ENV GOPATH=/go

RUN mkdir /go

RUN apk update && apk add go git

COPY . /go/src/github.com/Xe/macarena

RUN go get github.com/Xe/macarena/...

ONBUILD COPY config.json /macarena/config.json
WORKDIR /macarena
CMD "/go/bin/macarena"
