# oob-log-alert


English | [简体中文](README.zh-CN.md) 

## Introduction

Out-of-band monitoring alarm module. Subscribe to the alarm log sent by oob-log-agent, and send alarms after filtering by specified rules.

The location of oob-log-alert in the entire bmp architecture is shown in：[bmp arch](../bmp-scheduler/README.md)



## 2，Core Features

- Receives alarm logs sent by oob-log-agent.
- After rule filtering, send alarm logs to the console and operation platform.