FROM golang:1.17-alpine AS builder
ENV GOPROXY https://goproxy.cn
WORKDIR /usr/src/bmp-pronoea
COPY . .
RUN go build -o /tmp/bmp-pronoea .

FROM alpine:3.16
WORKDIR /home/bmp/bmp-pronoea
COPY --from=builder /tmp/bmp-pronoea .
COPY ./conf ./conf
CMD ["./bmp-pronoea"]

