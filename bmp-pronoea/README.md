# PRONOEA (bmp-pronoea)

English | [简体中文](README.zh-CN.md) 

## Introduction

`Bmp Pronoea` In order to further improve the ability of private bare metal platform to manage and monitor instances, 
and realize integrated management of monitoring and alarm, the performance monitoring (in-band monitoring) function was proposed. 
That is, it supports monitoring and alarm management of resources under user names. Therefore, 
the third-party alarm service Prometheus (official website: https://prometheus.io/) was introduced. 
The self-developed version bmp-pronoea (Pronoa) is based on the Prometheus service, integrates product requirements, 
and realizes resource monitoring and alarm management.

![bmp-pronoea.png](../docs/zh/bmp-deploy/picture/bmp-pronoea.png)
## Core functions
### Pronoea System Module
- **Alarm rule management (prometheus-rule-agent)**
  Provides the function of dynamically setting alarm rules. Receives the alarm rule information sent by the upper-layer BMP-API, 
  converts the alarm rule information into a yml file that Prometheus can recognize, and writes it to the specified directory.
- **Alarm information management (prometheus-alert-transform)**
  Receives the alert information from the Alertmanager module, converts the information format, and returns the alert information to the downstream module. 
  It can be the event center, BMP-API, or other possible system modules.
- **Monitoring data management (prometheus-monitor-data)**
  Prometheus provides an API-client to query the collected monitoring data. 
  The original intention of this module is to encapsulate the API-clieng layer and provide data query services for BMP-API.
- **Monitoring data clean (cache-del-server)**
  After the monitoring data received by Pushgateway is pulled by the Prometheus main service, 
  it will not be deleted actively, which will cause the accumulation of useless data and waste resources. 
  cache-del-server is designed to actively delete this junk data at regular intervals.

### External Dependencies
- **Proxy（bmp-monitor-proxy）**
Integrate the monitoring data of the instance, integrate the data into a specific recognizable data format, 
and finally push it to Pushgateway. For details, see bmp-monitor-proxy.
- **BMP API**
The interaction between the console, operation management background and Prometheus monitoring system relies on BMP-API, for example:setting alarm rules, querying monitoring data, receiving and processing alarm information, etc.
Here, BMP API includes bmp-console-api and bmp-openapi.