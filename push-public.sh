#!/bin/bash

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o go-ws cmd/go-ws/main.go
docker build --no-cache -t aracki/go-web-server .
rm go-ws
docker push aracki/go-web-server
