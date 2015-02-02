#!/bin/sh
export GOPATH=$GOPATH:`pwd`/../server

killall quanquan
cd ../
rm -f client/quanquan

rm -rf client/res
cp -rf server/res client/res

go build -a -v -o client/quanquan server/src/main.go
if [ $? -ne 0 ]; then
	echo "Failed to build"
	exit 1
fi

cd client
./quanquan