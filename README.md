# go-ws

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/3e4a3a4a337c4da5b7b2ccd5b144b47f)](https://app.codacy.com/app/Aracki/go-ws?utm_source=github.com&utm_medium=referral&utm_content=Aracki/go-ws&utm_campaign=Badge_Grade_Dashboard)
[![Go Report Card](https://goreportcard.com/badge/github.com/aracki/go-ws)](https://goreportcard.com/report/github.com/aracki/go-ws)
[![Build Status](https://travis-ci.org/Aracki/go-ws.svg?branch=master)](https://travis-ci.org/Aracki/go-ws)
[![](https://images.microbadger.com/badges/image/aracki/go-web-server.svg)](https://microbadger.com/images/aracki/go-web-server "Get your own image badge on microbadger.com")

An ultra simple Go App ready to be deployed via Kubernetes.

#### Build and run docker image
Install [dep tool](https://github.com/golang/dep#installation) and run `dep ensure`.<br>
Use `./run-local.sh` to run Docker container locally.<br>
Use `./push-public.sh` to rebuild Docker image & push it to the public [Docker hub](https://hub.docker.com/r/aracki/).

Continuous Integration: every push to Master branch will trigger Travis to run all these steps.

#### Run mongo locally

`docker run --name mongo --rm -p 27017:27017 mongo:latest`
