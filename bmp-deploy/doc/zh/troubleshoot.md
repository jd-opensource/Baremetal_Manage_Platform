# [bmp安装指导](main.md) - 故障排查
* pxe引导失败
1. 机器的引导模式是否为bios
2. 机器带外网络是否正确
3. 机器的bmc是否正常
4. 机器的网口连线是否正确
* 某个服务未正常运行
~~~
./bmp-deploy.sh status all          #检查哪些服务未正常运行
./bmp-deploy.sh start ServiceName   #启动服务
~~~

# 上一节 [装机流程](bm-deploy.md)
# 下一节 [实用工具](tool.md)