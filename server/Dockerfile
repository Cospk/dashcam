FROM golang:alpine as builder

WORKDIR /server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o dashcamServer .

FROM alpine:latest

LABEL MAINTAINER="xie18115@outlook.com"

WORKDIR /server

COPY --from=0 /server/dashcamServer ./
COPY --from=0 /server/config.yaml ./
COPY --from=0 /server/wait-for-it.sh ./

RUN mkdir -p ./resource/upload && chmod +x ./wait-for-it.sh

EXPOSE 8888

CMD ["./wait-for-it.sh", "./dashcamServer"]
