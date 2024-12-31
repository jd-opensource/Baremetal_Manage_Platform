FROM golang:1.17-alpine AS builder
ENV GOPROXY https://goproxy.cn
WORKDIR /usr/src/bmp-openapi-console
COPY . .
RUN go build -o /tmp/bmp-openapi-console .

FROM alpine:3.16
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories  \
    && apk add tzdata \
    && apk add dmidecode
WORKDIR /home/bmp/bmp-openapi-console
COPY --from=builder /tmp/bmp-openapi-console .
COPY ./conf ./conf
CMD ["./bmp-openapi-console"]
