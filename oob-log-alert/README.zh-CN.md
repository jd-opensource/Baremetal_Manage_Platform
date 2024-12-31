# oob-log-alert


[English](README.md) | 简体中文

## 介绍

带外监控告警模块。订阅oob-log-agent发送的告警日志，经过指定规则过滤后，发送告警。

oob-log-alert在整个bmp架构中的位置见：[bmp架构图](../bmp-scheduler/README.zh-CN.md)



## 2，核心功能

- 接收oob-log-agent发送的告警日志。
- 规则过滤后，发送告警日志到控制台和运营平台。

