FROM alpine:3.16
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk --no-cache update \
    && apk add --no-cache tftp-hpa        
CMD	["sh", "/tftp_init.sh"]
