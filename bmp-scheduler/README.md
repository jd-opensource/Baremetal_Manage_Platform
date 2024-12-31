# bmp-scheduler

English | [简体中文](README.zh-CN.md) 


## Introduction

The scheduling module mainly maintains the status flow of devices and instances.

bmp arch：
![bmp arch](docs/resources/bmp(open-source)arch.png)

Module Introduction：

- bmp-console-web：Console front end
- bmp-operation-web:Operation platform front end
- bmp-console-api：Adaptation layer for bmp-openapi-console and bmp-console-web
- bmp-operation-api：Adaptation layer for bmp-openapi and bmp-operation-web
- bmp-openapi-console:The core openapi interface layer used by the console
- bmp-openapi：The core openapi interface layer used by the operation platform
- bmp-scheduler：The core scheduling layer is responsible for the state transfer of devices and instances.
- bmp-driver：The out-of-band control layer is responsible for sending ipmitool commands to bm
- bmp-agent：Responsible for starting liveos and interacting with bmp-scheduler
- inbond-agent：In-band monitoring collection agent, responsible for bm data collection and reporting to inbond-agent-proxy
- bmp-dhcp-agent：Responsible for adding and deleting dhcp configuration
- inbond-agent-proxy：Responsible for cleaning the data collected by inbond-agent, reporting it to pushgateway, and finally storing it in prometheus
- push-gateway：Prometheus built-in module, responsible for receiving in-band monitoring data
- prometheus：Open source prometheus components
- pronoea:A self-made module that receives alarm rules and sends the generated alarms to bmp-openapi-console.
- oob-agent：Out-of-band monitoring collection agent
- oob-alert：Responsible for filtering the data collected by oob-agent according to specified rules and issuing business alarms when alarms are triggered.


bmp device and instance status flow：

![bmp LifeCycle](docs/resources/cps-lifecycle.png)

## 2，Core Features

- Equipment import：Bare metal equipment can be imported into the BMP system through the operation platform。
- Equipment Collection：Collect information about the disk, network card, etc. of bare metal devices.
- Equipment putaway：After collecting bare metal devices, bind them to the appropriate model and put them on the shelf. After the devices are put on the shelf, you can create instances in the console.
- Instance create：In the console, you can specify various configurations for the listed devices to create a physical machine running an operating system.
- Instance destroy：Destroy the system disk of the physical machine running the operating system and restore it to a bare metal device.
- Instance reinstall：Reinstall the operating system on a device that has an operating system installed.
- Instance startup：Send a remote out-of-band command to power on the system.
- Instance shutdown：Send a remote out-of-band command to shut down the system.
- Instance restart：Send a remote out-of-band command to reboot.
- Instance reset password：Reset the root user login password of the operating system.

