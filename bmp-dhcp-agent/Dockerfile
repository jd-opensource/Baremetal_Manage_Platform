FROM python:3.6.8-alpine
WORKDIR /home/bmp/bmp-dhcp-agent
ARG PIP_INDEX_URL=https://mirrors.aliyun.com/pypi/simple
COPY . .
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.jd.com/g' /etc/apk/repositories \
    && apk add dhcp \
    && apk add gcc musl-dev linux-headers \
    && pip install -r requirements.txt -i $PIP_INDEX_URL
ENV PYTHONPATH /home/bmp/bmp-dhcp-agent
CMD ["python3", "./bmpda/cmd/server.py"]
