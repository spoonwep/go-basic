FROM golang:1.19.4-alpine3.16 as base

MAINTAINER spoonwep<chenchao@imechos.com>

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add g++ && apk add make

WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN make buildfile

FROM alpine:latest

WORKDIR /app/
COPY --from=base /app/main .
ENTRYPOINT ["./main"]