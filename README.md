# go-ws

[![Build Status](https://travis-ci.org/Aracki/go-ws.svg?branch=master)](https://travis-ci.org/Aracki/go-ws)
[![](https://images.microbadger.com/badges/image/aracki/go-web-server.svg)](https://microbadger.com/images/aracki/go-web-server "Get your own image badge on microbadger.com")

An ultra simple Go App ready to be deployed via Kubernetes.

#### Build and run docker image
Use `./build-local.sh` to run Docker container locally.<br>
Use `./build.sh` to rebuild Docker image & push it to the public [Docker hub](https://hub.docker.com/r/aracki/).

#### Run mongo locally

`docker run --name mongo --rm -p 27017:27017 mongo:latest`
