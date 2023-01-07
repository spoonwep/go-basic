FROM golang:1.19.4-alpine3.16

MAINTAINER spoonwep<chenchao@imechos.com>

WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
CMD ["go", "run", "main.go"]