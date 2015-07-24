#!/bin/bash

set -e
set -x

source ./env.sh

cd ./src/cmd/macarena
go get ./...

cd ../../..
./bin/macarena -h
