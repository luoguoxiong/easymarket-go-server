FROM alpine:latest

MAINTAINER 一汎

#设置工作目录
WORKDIR /app

#设置时区
RUN echo "https://mirrors.ustc.edu.cn/alpine/v3.6/main" > /etc/apk/repositories && \
  echo "https://mirrors.ustc.edu.cn/alpine/v3.6/community" >> /etc/apk/repositories && \
  apk update && apk add ca-certificates && \
  apk add tzdata && \
  ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
  echo "Asia/Shanghai" > /etc/timezone

ADD . /app

ENTRYPOINT ./cmd/cmd --conf ./configs
