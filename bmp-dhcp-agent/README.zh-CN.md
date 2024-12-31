# BMP DHCP Agent
[English](README.md) | 简体中文

用于接受调度器的命令,对dhcp服务进行控制

## 核心特性
- MAC地址绑定IP，设置特定 MAC 地址的设备绑定固定IP。例如：
```
host b0-26-28-43-85-b0 {
  dynamic;
  hardware ethernet b0:26:28:43:85:b0;
  fixed-address 10.208.16.4;
}
```

- 网段配置，设置DHCP服务的网段和相关选项。例如
```
subnet 10.208.16.0 netmask 255.255.255.224  {
    option subnet-mask 255.255.255.224;
    option routers 10.208.16.1;
}
```