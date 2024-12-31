# bmp-monitor-agent

English | [简体中文](README.zh-CN.md) 

## 1，Introduction

The in-band monitoring agent module runs on the bm operating system, collects bm's disk, network and other data, and reports it to bmp-monitor-proxy.

The position of bmp-monitor-agent in the entire bmp architecture is shown in：[bmp arch](../bmp-scheduler/README.md)



## 2，Core Features

Collects data from device disks, networks, etc., and reports the monitoring data to bmp-monitor-proxy.


## 3，Usage

Package linux amd x_64 version agent:
make linux 
Package the Windows version agent:
make windows

download bmp-monitor-agent：
curl -O http://..../download/agent/bmp-agent.bin

Adding operation permissions ：
chmod +x bmp-agent.bin 

run：
./bmp-agent.bin 

service name:bmpd
service status:service bmpd status 
start service:service start bmpd 
stop service:service stop bmpd
