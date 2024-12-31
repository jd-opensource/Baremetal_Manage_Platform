#  [bmp安装指导](main.md) - 实用工具
tool目录下包含一些工具脚本，用于采集设备信息等

## 采集设备信息
collect_device_info.sh 使用说明
~~~
./collect_device_info.sh -h

Usage:  ./collect_device_info.sh --sn SN --iloIp iloIp --iloUser iloUser --iloPassword iloPassword --ip ip --mask mask --gateway gateway --mac MAC [--raidDriver raidDriver]

~~~

参数说明：
~~~
SN          设备sn
iloIp       设备带外网口ip
iloUser     设备带外管理用户名
iloPassword 设备带外管理密码
ip          设备管理网口1 ip
mask        设备管理网口1 ip所属子网掩码
gateway     设备管理网口1 ip所属子网网关
MAC         设备管理网口1 mac
raidDriver  设备raid卡驱动，目前支持的驱动有：sas2ircu sas3ircu megacli64 storcli64 perccli64 arcconf no_raid，默认使用megacli64
~~~

示例：
~~~
./collect_device_info.sh --sn 210235A3P6N1830002MH --iloIp 192.168.0.2 --iloUser xxx --iloPassword 'xxx' --ip 10.208.14.170 --mask 255.255.255.224 --gateway 10.208.14.161 --mac '3C:FD:FE:5E:C6:74'
~~~

采集设备信息任务完成后（需要等待2-5分钟），会生成文件 /var/log/bmp/bmp-scheduler/disk_location_dir/SN，该文件包含设备磁盘插槽信息。
~~~
# cat /var/log/bmp/bmp-scheduler/disk_location_dir/210235A3P6N1830002MH 
{"controllers":[{"adapter_id":"1","disk_info":"Controllers found: 1\n----------------------------------------------------------------------\nController information\n----------------------------------------------------------------------\n   Controller Status                        : Optimal\n   Controller Mode                          : Mixed\n   Channel description                      : SCSI\n   Controller Model                         : Adaptec SmartIOC 8i\n   Controller Serial Number                 : Unknown\n   Controller World Wide Name               : 5600B036FAEDD000\n   Physical Slot                            : 255\n   Temperature                              : 61 C/ 141 F (Normal)\n   Power Consumption                        : Not Applicable\n   Power Mode                               : Unknown\n   Power Mode Operational                   : Unknown\n   Survival Mode                            : Unknown\n   Host bus type                            : PCIe 3.0\n   Host bus speed                           : 7880 MBps\n   Host bus link width                      : 8 bit(s)/link(s)\n   PCI Address (Bus:Device:Function)        : 0:1:0:0\n   I2C Address                              : 0xDE\n   I2C Clock Speed ......
~~~

## 重试指令
retry_command.sh 使用说明
~~~
 ./retry_command.sh -h

Usage:  ./retry_command.sh Command_id

~~~

# 上一节 [故障排查](troubleshoot.md)
# 下一节 [配置管理](config.md)
