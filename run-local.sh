#!/bin/bash

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o go-ws cmd/go-ws/main.go
docker build --no-cache -t go-web-server .
docker run --name go-ws -p 8080:8080 --rm go-web-server:latest
rm go-ws
