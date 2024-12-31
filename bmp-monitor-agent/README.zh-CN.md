# 带内监控bmp-monitor-agent

[English](README.md) | 简体中文

## 介绍

带内监控agent模块，运行在bm操作系统上，采集bm的磁盘，网络等数据，上报到bmp-monitor-proxy。

bmp-monitor-agent在整个bmp架构中的位置见：[bmp架构图](../bmp-scheduler/README.zh-CN.md)



## 2，核心功能

- 设备磁盘，网络等数据的采集，并上报监控数据到bmp-monitor-proxy。



## 3，使用

打包linux amd x_64版本agent:
make linux 
打包windows版本agent:
make windows

下载agent：
curl -O http://..../download/agent/bmp-agent.bin

增加操作权限：
chmod +x bmp-agent.bin 

运行：
./bmp-agent.bin 

服务名称:bmpd
查看服务状态:service bmpd status 
start服务:service start bmpd 
stop服务:service stop bmpd
