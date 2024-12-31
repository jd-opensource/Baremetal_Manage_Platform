# BMP DHCP Agent
English | [简体中文](README.zh-CN.md) 

Receive the scheduler's commands through MQ and control the DHCP service

## Core Features
- MAC Address Binding to IP, Assign a fixed IP address to a device based on its MAC address. Example:
```
host b0-26-28-43-85-b0 {
  dynamic;
  hardware ethernet b0:26:28:43:85:b0;
  fixed-address 10.208.16.4;
}
```

- Subnet Configuration: Configure DHCP settings for a specific subnet. Example:
```
subnet 10.208.16.0 netmask 255.255.255.224  {
    option subnet-mask 255.255.255.224;
    option routers 10.208.16.1;
}
```