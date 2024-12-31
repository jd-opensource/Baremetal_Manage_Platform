# 服务启动失败
重新启动某服务
~~~
./bmp-deploy.sh status all          #检查哪些服务未正常运行
./bmp-deploy.sh start ServiceName   #启动服务
~~~

# pxe引导失败原因
pxe引导失败
机器的引导模式是否为bios
机器带外网络是否正确
机器的bmc是否正常
机器的网口连线是否正确

# docker未启动
报错：Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
重启manager节点之后，需要重启docker服务
~~~
systemctl start docker.service
~~~
