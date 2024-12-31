# bmp-openapi-console


English | [简体中文](README.zh-CN.md) 

## Introduction

The core interface layer of the console focuses on the life cycle of the instance, provides interfaces for bmp-openapi-console and developers, and transmits user actions to bmp-scheduler through intermediate keys such as mq and redis.

The location of bmp-openapi-console in the entire bmp architecture is shown in：[bmp arch](../bmp-scheduler/README.md)



## 2，Core Features

- Instance create：In the console, you can specify various configurations for the listed devices to create a physical machine running an operating system.
- Instance destroy：Destroy the system disk of the physical machine running the operating system and restore it to a bare metal device.
- Instance reinstall：Reinstall the operating system on a device that has an operating system installed.
- Instance startup：Send a remote out-of-band command to power on the system.
- Instance shutdown：Send a remote out-of-band command to shut down the system.
- Instance restart：Send a remote out-of-band command to reboot.
- Instance reset password：Reset the root user login password of the operating system.