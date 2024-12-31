FROM python:3.6-alpine
WORKDIR /home/bmp/bmp-dhcp-agent
COPY . .
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories  \
   && apk add dhcp  \
   && pip install -i https://mirrors.tuna.tsinghua.edu.cn/pypi/web/simple -r requirements.txt
ENV PYTHONPATH /home/bmp/bmp-dhcp-agent
CMD ["python3", "./bmpda/cmd/server.py"]
