#!/bin/sh
killall quanquan
go build -o webclient/quanquan src/github.com/anzhihun/quanquan/main.go
cd webclient
./quanquan