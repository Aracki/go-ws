# TODO if switch to alpine, find hostname command alternative
FROM ubuntu:16.04

COPY go-ws go-ws
ENTRYPOINT ["./go-ws"]
EXPOSE 8080
