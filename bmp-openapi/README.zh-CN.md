# bmp-openapi

[English](README.md) | 简体中文


## 介绍

运营平台核心接口层，围绕设备的生命周期，向上为bmp-openapi和开发者提供接口，向下通过mq、redis等中间键将用户动作传递到bmp-scheduler。

bmp-openapi在整个bmp架构中的位置见：[bmp架构图](../bmp-scheduler/README.zh-CN.md)



## 2，核心功能

- 设备导入：裸金属设备可经运营平台导入到bmp系统中。
- 设备采集：采集裸金属设备的磁盘、网卡等信息。
- 设备上架：裸金属设备采集后绑定合适机型，即可上架，上架后的设备可以在控制台创建实例。

