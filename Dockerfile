FROM ubuntu:16.04

ADD go-ws go-ws
ENTRYPOINT ["./go-ws"]
EXPOSE 8080
