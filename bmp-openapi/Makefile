
BUILD_TIME := $(shell date +%Y%m%d%H%M%S)
VERSION := $(shell git describe --tags --always)
COMMITID := $(shell git rev-parse ${VERSION} | cut -c 1-32)
APP_NAME := ironic-api-golang

#GOARCH=amd64 GOOS=linux
ldflags ="-s -X main.VERSION=${COMMITID} -X main.BUILD_TIME=${BUILD_TIME}"
all: 
	GOARCH=amd64 GOOS=linux go build -ldflags ${ldflags} -o ${APP_NAME} main.go
