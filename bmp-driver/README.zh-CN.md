# bmp-driver


[English](README.md) | 简体中文

## 介绍

带外控制模块，向上承接bmp-scheduler经mq发送过来的动作消息，向下通过ipmitool往设备发送带外远程命令。

bmp-driver在整个bmp架构中的位置见：[bmp架构图](../bmp-scheduler/README.zh-CN.md)



## 2，核心功能


- 实例开机：发送远程带外命令开机。
- 实例关机：发送远程带外命令关机。
- 实例重启：发送远程带外命令重启。

