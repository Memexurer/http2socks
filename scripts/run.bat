@echo off
set GOARCH=amd64
set GOOS=windows
go build -o build_win/ .
set LISTEN_ADDR=127.0.0.1:12321
"build_win/httptosocks.exe"