FROM alpine:latest

MAINTAINER 一汎

WORKDIR /usr/src/goods

COPY ./cmd/cmd WORKDIR

ENTRYPOINT  /usr/src/goods/cmd --conf /config/
