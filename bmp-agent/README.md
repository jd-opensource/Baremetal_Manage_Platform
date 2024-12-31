# BMP Agent
English | [简体中文](README.zh-CN.md) 

The BMP Agent runs on RamOS, receiving commands from a scheduler via MQ (Message Queue) and executing corresponding instructions to install and configure the operating system.

[Command Details](./doc/command.md)

## Capabilities
1. Collect Server Configuration Information: Gather details such as network interface information, disk information, switch information, etc.
2. Install OS on Servers and Write Configuration Information

## Deploy the Agent to the RamOS environment
Refer to [imagebuild-docker](./imagebuild-docker/README.md)

## Log Directory
The Agent log file under RamOS is located at '/tmp/bmpa.log'


## Service Control
In RamOS, you can control the Agent service with the following commands:

1. View service status
```shell
systemctl status bmpa
```
2. Start the service
```shell
systemctl start bmpa
```
3. Stop service
```shell
systemctl stop bmpa
```
## Environment configuration information
Environment configuration information is passed through kernel parameters

```shell
cat /proc/cmdline
```

