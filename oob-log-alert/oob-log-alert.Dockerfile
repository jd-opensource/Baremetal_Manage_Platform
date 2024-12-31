FROM quay.io/goswagger/swagger AS SDK
ENV GOPROXY https://goproxy.cn
WORKDIR /usr/src/bmp-openapi
COPY ./bmp-openapi .
RUN swagger generate spec -o ./swagger.yml
RUN swagger validate swagger.yml
RUN mkdir -p ../bmp-openapi-sdk
RUN cd ../bmp-openapi-sdk && go mod init coding.jd.com/aidc-bmp/bmp-openapi-sdk
RUN cd ../bmp-openapi
RUN swagger generate client -f swagger.yml --target=../bmp-openapi-sdk

FROM golang:1.17-alpine AS builder
ENV GOPROXY https://goproxy.cn
WORKDIR /usr/src/bmp-oob-alert
COPY ./oob-log-alert .
COPY --from=SDK /usr/src/bmp-openapi-sdk ./bmp_vendor/bmp-openapi-sdk
RUN go mod tidy -compat=1.17
RUN go mod vendor
RUN go build -o /tmp/bmp-oob-alert .

FROM alpine:3.16
WORKDIR /home/bmp/bmp-oob-alert
COPY --from=builder /tmp/bmp-oob-alert .
COPY ./oob-log-alert/conf ./conf
CMD ["./bmp-oob-alert"]
