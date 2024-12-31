FROM quay.io/goswagger/swagger AS SDK
ENV GOPROXY https://goproxy.cn
WORKDIR /usr/src/bmp-openapi-tmp
COPY ./bmp-openapi .
RUN swagger generate spec -o ./swagger.yml
RUN swagger validate swagger.yml
RUN mkdir -p ../bmp-openapi-sdk
RUN cd ../bmp-openapi-sdk && go mod init coding.jd.com/aidc-bmp/bmp-openapi-sdk
RUN cd ../bmp-openapi-tmp
RUN swagger generate client -f swagger.yml --target=../bmp-openapi-sdk

FROM golang:1.17-alpine AS builder
ENV GOPROXY https://goproxy.cn
WORKDIR /usr/src/bmp-operation-api
COPY ./bmp-operation-api .
COPY --from=SDK /usr/src/bmp-openapi-sdk ./bmp_vendor/bmp-openapi-sdk
RUN go mod tidy -compat=1.17
RUN go mod vendor
RUN go build -o /tmp/bmp-operation-api .

FROM alpine:3.16
WORKDIR /home/bmp/bmp-operation-api
COPY --from=builder /tmp/bmp-operation-api .
COPY ./bmp-operation-api/conf ./conf
CMD ["./bmp-operation-api"]
