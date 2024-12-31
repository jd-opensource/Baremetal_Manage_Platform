FROM node:16.19-alpine3.16 AS builder
WORKDIR /usr/src/bmp-console-web
COPY . .
RUN npm config set registry https://registry.npmmirror.com && npm install --legacy-peer-deps && npm run build-pre

FROM nginx:1.22-alpine
WORKDIR /home/bmp/bmp-console-web
COPY --from=builder /usr/src/bmp-console-web/dist .
