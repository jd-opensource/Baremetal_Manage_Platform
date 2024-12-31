# bmp-openapi-console


[English](README.md) | 简体中文

## 介绍

控制台核心接口层，围绕实例的生命周期，向上为bmp-openapi-console和开发者提供接口，向下通过mq、redis等中间键将用户动作传递到bmp-scheduler。

bmp-openapi-console在整个bmp架构中的位置见：[bmp架构图](../bmp-scheduler/README.zh-CN.md)



## 2，核心功能

- 创建实例：在控制台可以将上架的设备指定各种配置，创建成一台运行操作系统的物理机。
- 销毁实例：将运行操作系统的物理机销毁系统盘，还原成裸金属设备。
- 重装实例：对已安装操作系统的设备重新安装操作系统。
- 实例开机：发送远程带外命令开机。
- 实例关机：发送远程带外命令关机。
- 实例重启：发送远程带外命令重启。
- 实例重置密码：重置操作系统root用户登陆密码。

