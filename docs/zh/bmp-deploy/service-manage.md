# [bmp安装指导](main.md) - 服务管理
bmp-deploy.sh 使用说明
~~~
# ./bmp-deploy.sh -h

Usage:  ./bmp-deploy.sh [-t|--tag TAG] COMMAND SERVICE...

COMMAND:
    start       Install and start SERVICE
    stop        Stop SERVICE
    restart     restart SERVICE
    status      check SERVICE status
    images      display SERVICE image
    uninstall   uninstall SERVICE
    reinstall   reinstall SERVICE
    config      get SERVICE config

SERVICE:
    all                 All SERVICES
    bmp-db
    bmp-redis
    bmp-mq
    bmp-image
    bmp-pushgateway
    bmp-alertmanager
    bmp-prometheus
    bmp-tftp
    bmp-rsyslog
    bmp-console-web
    bmp-operation-web
    bmp-console-api
    bmp-operation-api
    bmp-openapi
    bmp-openapi-console
    bmp-scheduler
    bmp-driver
    bmp-dhcp-agent
    bmp-oob-alert
    bmp-oob-agent
    bmp-monitor-proxy
    bmp-pronoea
~~~
## 停止服务  
./bmp-deploy.sh stop all | [service...]
~~~
#停止bmp-console-web服务
./bmp-deploy.sh stop bmp-console-web

#停止bmp-console-web和bmp-operation-web服务
./bmp-deploy.sh stop bmp-console-web bmp-operation-web

#停止所有服务
./bmp-deploy.sh stop all
~~~

## 启动服务  
./bmp-deploy.sh start all | [service...]
~~~
#启动bmp-console-web服务
./bmp-deploy.sh start bmp-console-web

#启动bmp-console-web和bmp-operation-web服务
./bmp-deploy.sh start bmp-console-web bmp-operation-web

#启动所有服务
./bmp-deploy.sh start all
~~~

## 重启服务
./bmp-deploy.sh restart all | [service...]
~~~
#重启bmp-console-web服务
./bmp-deploy.sh restart bmp-console-web
~~~

## 查看服务运行状态 
./bmp-deploy.sh status all | [service...]
~~~
#查看所有服务运行状态
./bmp-deploy.sh status all
~~~

## 查看服务关联的容器镜像
./bmp-deploy.sh images all | [service...]
~~~
#查看bmp-console-web服务关联的容器镜像
./bmp-deploy.sh images bmp-console-web
~~~

## 删除服务
./bmp-deploy.sh uninstall all | [service...]

说明： 相关数据也会一并删除
~~~
#删除单个服务，如bmp-console-web
./bmp-deploy.sh uninstall bmp-console-web

#删除所有服务
./bmp-deploy.sh uninstall all
~~~

## 卸载bmp(删除所有服务)
~~~
cd /root/bmp-deploy/
./bmp-deploy.sh uninstall all
cd /root
rm -rf bmp-deploy bmp-deploy*.tar.gz
rm -rf /var/log/bmp
~~~

## 重装服务(删除+按装)
./bmp-deploy.sh reinstall all | [service...]
~~~
#重装单个服务，如bmp-console-web
./bmp-deploy.sh reinstall bmp-console-web
~~~

## 获取服务配置
./bmp-deploy.sh config all | [service...]

## 部署bmp时指定版本
./bmp-deploy.sh start all | [service...]    -t|--tag TAG
~~~
#部署1.0版本
./bmp-deploy.sh start all --tag 1.0
~~~

