# bmp-openapi


English | [简体中文](README.zh-CN.md) 

## Introduction

The core interface layer of the operation platform focuses on the life cycle of the device, provides interfaces for bmp-openapi and developers, and transmits user actions to bmp-scheduler through intermediate keys such as mq and redis.

See the location of bmp-openapi in the entire bmp architecture.：[bmp arch](../bmp-scheduler/README.md)



## 2，Core Features

- Equipment import：Bare metal equipment can be imported into the BMP system through the operation platform。
- Equipment Collection：Collect information about the disk, network card, etc. of bare metal devices.
- Equipment putaway：After collecting bare metal devices, bind them to the appropriate model and put them on the shelf. After the devices are put on the shelf, you can create instances in the console.