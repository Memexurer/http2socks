@echo off
set GOOS=linux
go build -o build/ .
set DOCKER_HOST=tcp://127.0.0.1:2375
docker build --network=host --tag localhost:5000/http2socks .
docker push localhost:5000/http2socks:latest