FROM golang:1.17-alpine AS builder
ENV GOPROXY https://goproxy.cn
WORKDIR /usr/src/bmp-monitor-proxy
COPY . .
RUN go build -o /tmp/bmp-monitor-proxy .

FROM alpine:3.16
WORKDIR /home/bmp/bmp-monitor-proxy
COPY --from=builder /tmp/bmp-monitor-proxy .
COPY ./conf ./conf
CMD ["./bmp-monitor-proxy"]
