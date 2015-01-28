echo off 

set GOPATH=%GOPATH%;%~dp0\..\server

taskkill /f /t /im quanquan.exe
cd ..\
del client\quanquan.exe

go build -a -v -o client\quanquan.exe server\src\main.go 
if %errorlevel% NEQ 0 (
	echo on
	echo "Failed to build"
	exit %errorlevel%
)

cd client
start quanquan.exe
