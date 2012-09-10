#!/bin/sh
APP="pinterest"
ADDR=":8083"
PWD=`pwd`/..
GO_PWD=${PWD}/../..
echo "Golanger Web Framework"
echo "Golanger is a lightweight framework for writing web applications in Golang."
export GOPATH=${GO_PWD}/framework:${GO_PWD}/samples/add-on:${PWD}
cd ${PWD}/src

if [ -f ${APP} ]; then
    rm ${APP}
fi

echo "Building ${APP}"
go build .

if [ -f src ]; then
    mv ./src ${APP}
    echo "Runing ${APP}"
    ./$APP -addr=${ADDR}
fi
