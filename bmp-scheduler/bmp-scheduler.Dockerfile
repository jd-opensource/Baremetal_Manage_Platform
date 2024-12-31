FROM golang:1.17-alpine AS builder
ENV GOPROXY https://goproxy.cn
WORKDIR /usr/src/bmp-scheduler
COPY . .
RUN go build -o /tmp/bmp-scheduler .

FROM alpine:3.16
WORKDIR /home/bmp/bmp-scheduler
COPY --from=builder /tmp/bmp-scheduler .
COPY ./conf ./conf
CMD ["./bmp-scheduler"]
