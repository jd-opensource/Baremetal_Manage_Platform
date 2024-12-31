# bmp-openapi-console

English | [简体中文](README.zh-CN.md) 


## Introduction

The out-of-band monitoring collection module is responsible for obtaining system logs from the device through the ipmitool command and publishing them to redis

The location of oob-log-agent in the entire bmp architecture is shown in：[bmp arch](../bmp-scheduler/README.md)



## 2，Core Features

- Collect the system log of the target device through the ipmitool command.
- Publish to redis after filtering.