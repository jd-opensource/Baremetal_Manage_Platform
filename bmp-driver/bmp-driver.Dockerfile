FROM golang:1.17-alpine AS builder
ENV GOPROXY https://goproxy.cn
WORKDIR /usr/src/bmp-driver
COPY . .
RUN go build -o /tmp/bmp-driver .

FROM alpine:3.16
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk add --no-cache ipmitool
WORKDIR /home/bmp/bmp-driver
COPY --from=builder /tmp/bmp-driver .
COPY ./conf ./conf
CMD ["./bmp-driver"]
