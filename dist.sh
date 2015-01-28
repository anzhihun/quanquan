#!/bin/sh
export GOPATH=$GOPATH:`pwd`/server

killall quanquan
rm -f client/quanquan

go build -a -v -o client/quanquan server/src/main.go
if [ $? -ne 0 ]; then
	echo "Failed to build"
	exit 1
fi

cd client
./quanquan