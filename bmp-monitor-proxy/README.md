# bmp-monitor-proxy


English | [简体中文](README.zh-CN.md) 

## Introduction

The in-band monitoring proxy module receives monitoring data reported by the bmp-monitor-agent of each device, processes it, and reports it to the push-gateway module.

The position of bmp-monitor-proxy in the entire bmp architecture is shown in：[bmp arch](../bmp-scheduler/README.md)



## 2，Core Features

- Receive monitoring data reported by bmp-monitor-agent of each device.
- After processing and cleaning, report to the push-gateway module.