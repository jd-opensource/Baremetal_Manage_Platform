FROM quay.io/goswagger/swagger AS SDK
ENV GOPROXY https://goproxy.cn
WORKDIR /usr/src/bmp-openapi-console
COPY ./bmp-openapi-console .
RUN swagger generate spec -o ./swagger.yml
RUN swagger validate swagger.yml
RUN mkdir -p ../bmp-openapi-console-sdk
RUN cd ../bmp-openapi-console-sdk && go mod init coding.jd.com/aidc-bmp/bmp-openapi-console-sdk
RUN cd ../bmp-openapi-console
RUN swagger generate client -f swagger.yml --target=../bmp-openapi-console-sdk

FROM golang:1.17-alpine AS builder
ENV GOPROXY https://goproxy.cn
WORKDIR /usr/src/bmp-console-api
COPY ./bmp-console-api .
COPY --from=SDK /usr/src/bmp-openapi-console-sdk ./bmp_vendor/bmp-openapi-console-sdk
RUN go mod tidy -compat=1.17
RUN go mod vendor
RUN go build -o /tmp/bmp-console-api .

FROM alpine:3.16
WORKDIR /home/bmp/bmp-console-api
COPY --from=builder /tmp/bmp-console-api .
COPY ./bmp-console-api/conf ./conf
CMD ["./bmp-console-api"]

