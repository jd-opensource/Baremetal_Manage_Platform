FROM golang:1.17-alpine AS builder
ENV GOPROXY https://goproxy.cn
WORKDIR /usr/src/bmp-oob-agent
COPY . .
RUN go build -o /tmp/bmp-oob-agent .

FROM alpine:3.16
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk add ipmitool
WORKDIR /home/bmp/bmp-oob-agent
COPY --from=builder /tmp/bmp-oob-agent .
COPY ./conf ./conf
CMD ["./bmp-oob-agent"]
