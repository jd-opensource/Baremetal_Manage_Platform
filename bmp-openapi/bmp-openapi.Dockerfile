FROM golang:1.17-alpine AS builder
ENV GOPROXY https://goproxy.cn
WORKDIR /usr/src/bmp-openapi
COPY . .
RUN go build -o /tmp/bmp-openapi .

FROM alpine:3.16
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
&& apk add tzdata \
&& apk add dmidecode
WORKDIR /home/bmp/bmp-openapi
COPY --from=builder /tmp/bmp-openapi .
COPY ./conf ./conf
CMD ["./bmp-openapi"]
