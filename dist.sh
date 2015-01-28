#!/bin/sh
export GOPATH=$GOPATH:`pwd`

killall quanquan
rm -f quanquan

go build -a -v -o quanquan src/main.go
./quanquan