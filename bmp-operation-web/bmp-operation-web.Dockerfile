FROM node:16.19-alpine3.16 AS builder
WORKDIR /usr/src/bmp-operation-web
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk add python3 make gcc g++
COPY . .
RUN npm config set registry https://registry.npmmirror.com && npm install --legacy-peer-deps && npm run build-pre

FROM nginx:1.22-alpine
WORKDIR /home/bmp/bmp-operation-web
COPY --from=builder /usr/src/bmp-operation-web/dist .
