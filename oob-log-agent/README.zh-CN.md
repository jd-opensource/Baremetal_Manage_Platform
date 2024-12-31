# oob-log-agent

[English](README.md) | 简体中文


## 介绍

带外监控采集模块，负责通过ipmitool命令从设备获取系统日志，并publish到redis中

oob-log-agent在整个bmp架构中的位置见：[bmp架构图](../bmp-scheduler/README.zh-CN.md)



## 2，核心功能

- 通过ipmitool命令采集目标设备的系统日志
- 过滤后publish到redis中

