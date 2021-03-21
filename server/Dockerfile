FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
WORKDIR /go/src/go_base_server
COPY go_base_server/ ./

RUN go env && go list && go build -o go_base_server .

FROM alpine:latest
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /go/src/go_base_server

COPY --from=0 /go/src/go_base_server/go_base_server ./
COPY --from=0 /go/src/go_base_server/i18n ./i18n
COPY --from=0 /go/src/go_base_server/public ./public
COPY --from=0 /go/src/go_base_server/config ./config
COPY --from=0 /go/src/go_base_server/template ./template

ENTRYPOINT ./go_base_server
# 根据Dockerfile生成Docker镜像

# docker build -t gva-go_base_server:1.0 .

#- 根据Docker镜像启动Docker容器
#    - 后台运行
#    - ```
#    docker run -d -p 8888:8888 --name gva-go_base_server-v1 gva-go_base_server:1.0
#      ```
#    - 以可交互模式运行, Ctrl + p + q 后台运行
#    - ```
#    docker run -it -p 8888:8888 --name gva-go_base_server-v1 gva-go_base_server:1.0
#      ```