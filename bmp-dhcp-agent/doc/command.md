指令实例模板
===========================================

### ping ip

```json
{
    "sn" : "0",
    "action" : "PingHost",
    "ip" : "10.0.0.2"
}
```

### dhcp config add host

```json
{
    "action": "DHCPConfigAddHost",
    "sn": "0",
    "ip": "10.0.0.2",
    "mac": "00:01:6C:06:A6:29" 
}

```

### dhcp config delete host

```json
{
    "action": "DHCPConfigDelHost",
    "sn": "0",
    "ip": "10.0.0.2",
    "mac":"00:01:6C:06:A6:29"
}
```

### dhcp config add subnet

```json
{
    "action": "DHCPConfigAddSubnet",
    "sn": "0",
    "subnet": "10.0.0.1",
    "subnet_mask": "255.255.255.255",
    "routes": "10.209.10.1"
}

```

### dhcp config delete subnet

```json
{
    "action": "DHCPConfigDelSubnet",
    "sn": "0",
    "subnet": "10.0.0.1"
}
```

### tftp config add grub 

```json
{
    "action": "TFTPConfigAddGrub",
    "sn": "0",
    "mac": "00:01:6C:06:A6:29", //此mac和DHCPConfigAddHost的mac要一致
    "boot_mode": "bios", //bios、uefi。不传此参数则同时配置所有
    "arch": "x86_64", //系统架构 x86_64、aarch64、loongarch64。不传此参数则同时配置所有
    "kernel_name": "v2.0.3-centos_7_9-2024072217-vmlinuz",
    "initramfs_name": "v2.0.3-centos_7_9-2024072217-initramfs.gz"
}
```
备注：采集时可以不传boot_mode和arch参数，这样就会生成多个配置文件，满足当前机器在每个架构下的配置

### tftp config delete grub

```json
{
    "action": "TFTPConfigDelGrub",
    "sn": "0",
    "mac":"00:01:6C:06:A6:29"
}
```