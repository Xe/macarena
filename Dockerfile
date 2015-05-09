FROM golang:1.4.2

RUN go get github.com/constabulary/gb/...

COPY . /usr/src/macarena
COPY ./build/run.sh /macarena/run.sh

RUN cd /usr/src/macarena &&\
    gb build all &&\
    cp bin/macarena /macarena/macarena

ONBUILD COPY config.json /macarena/config.json
WORKDIR /macarena
CMD "/macarena/run.sh"
