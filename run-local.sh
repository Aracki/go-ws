#!/bin/bash

echo "building go binary..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o go-ws cmd/go-ws/main.go

echo "building docker image..."
docker build --no-cache -t go-web-server .

echo "deleting go-ws binary..."
rm go-ws

echo "running docker container..."
docker run --name go-ws -p 8080:8080 --rm go-web-server:latest
