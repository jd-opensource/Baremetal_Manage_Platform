# BMP Agent
[English](README.md) | 简体中文

运行于RamOS中，通过MQ接收调度器下发的指令，并执行相应命令, 进行操作系统的安装及配置。

[指令详情](./doc/command.md)

## 能力
1. 采集服务器相关配置信息。如网卡信息、硬盘信息、交换机信息等
2. 为服务器安装OS系统，并写入相关配置信息

## 将Agent部署到RamOS中
详见[imagebuild-docker](./imagebuild-docker/README.md)

## 日志目录
RamOS下Agent日志文件位于'/tmp/bmpa.log'

## 服务控制
在RamOS中，可以通过以下命令控制Agent服务：
1. 查看服务状态
```shell
systemctl status bmpa
```
2. 启动服务
```shell
systemctl start bmpa
```
3. 停止服务
```shell
systemctl stop bmpa
```
## 环境配置信息
环境配置信息通过内核参数传递，详见
```shell
cat /proc/cmdline
```

