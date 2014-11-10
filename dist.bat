echo off 

taskkill /f /t /im quanquan.exe

set GOPATH=%GOPATH%;%~dp0
echo %GOPATH%
go build -a -v -o quanquan.exe src\main.go 
if %errorlevel% NEQ 0 (
	echo on
	echo "failed to build!"
	exit %errorlevel%
)

start quanquan.exe
