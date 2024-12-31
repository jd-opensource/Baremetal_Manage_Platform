# bmp-driver


English | [简体中文](README.zh-CN.md) 

## Introduction

The out-of-band control module receives action messages sent by bmp-scheduler via MQ and sends out-of-band remote commands to the device via ipmitool.

The position of bmp-driver in the entire bmp architecture is shown in：[bmp arch](../bmp-scheduler/README.md)



## 2，Core Features


- Instance startup：Send a remote out-of-band command to power on the system.
- Instance shutdown：Send a remote out-of-band command to shut down the system.
- Instance restart：Send a remote out-of-band command to reboot.

