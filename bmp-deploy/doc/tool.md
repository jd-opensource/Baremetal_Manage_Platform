# [BMP Installation Guide](main.md) - Useful Tools
The `tool` directory contains various tool scripts for collecting device information, etc.

## Collecting Device Information
`collect_device_info.sh` Usage Guide
~~~
./collect_device_info.sh -h

Usage: ./collect_device_info.sh --sn SN --iloIp iloIp --iloUser iloUser --iloPassword iloPassword --ip ip --mask mask --gateway gateway --mac MAC [--raidDriver raidDriver]

~~~

Parameter Explanation:
~~~
SN          Device SN
iloIp       Device iLO IP
iloUser     Device iLO management username
iloPassword Device iLO management password
ip          Device management interface 1 IP
mask        Device management interface 1 IP subnet mask
gateway     Device management interface 1 IP subnet gateway
MAC         Device management interface 1 MAC
raidDriver  Device RAID card driver, currently supported drivers include: sas2ircu sas3ircu megacli64 storcli64 perccli64 arcconf no_raid, default is megacli64
~~~

Example:
~~~
./collect_device_info.sh --sn 210235A3P6N1830002MH --iloIp 192.168.0.2 --iloUser xxx --iloPassword 'xxx' --ip 10.208.14.170 --mask 255.255.255.224 --gateway 10.208.14.161 --mac '3C:FD:FE:5E:C6:74'
~~~

After the device information collection task is completed (which may take 2-5 minutes), a file `/var/log/bmp/bmp-scheduler/disk_location_dir/SN` will be generated. This file contains the device's disk slot information.
~~~
# cat /var/log/bmp/bmp-scheduler/disk_location_dir/210235A3P6N1830002MH 
{"controllers":[{"adapter_id":"1","disk_info":"Controllers found: 1\n----------------------------------------------------------------------\nController information\n----------------------------------------------------------------------\n   Controller Status                        : Optimal\n   Controller Mode                          : Mixed\n   Channel description                      : SCSI\n   Controller Model                         : Adaptec SmartIOC 8i\n   Controller Serial Number                 : Unknown\n   Controller World Wide Name               : 5600B036FAEDD000\n   Physical Slot                            : 255\n   Temperature                              : 61 C/ 141 F (Normal)\n   Power Consumption                        : Not Applicable\n   Power Mode                               : Unknown\n   Power Mode Operational                   : Unknown\n   Survival Mode                            : Unknown\n   Host bus type                            : PCIe 3.0\n   Host bus speed                           : 7880 MBps\n   Host bus link width                      : 8 bit(s)/link(s)\n   PCI Address (Bus:Device:Function)        : 0:1:0:0\n   I2C Address                              : 0xDE\n   I2C Clock Speed......
~~~

## Retry Command
`retry_command.sh` Usage Guide
~~~
./retry_command.sh -h

Usage: ./retry_command.sh Command_id

~~~

### Previous section [Troubleshooting](troubleshoot.md)
### Next section [Configuration Management](config.md)
