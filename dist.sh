#!/bin/sh
export GOPATH=$GOPATH:`pwd`
killall quanquan
go build -a -v -o quanquan src/main.go
./quanquan