指令实例模板
===========================================

### ping 探测连接建立

```json
{
    "action" : "Ping",
    "sn" : "0"
}
```

### 收集硬件信息

```json
{
    "action" : "CollectHardwareInfo",
    "sn" : "0"
}
```

response
```json
{
    "action": "CollectHardwareInfo",
    "status": "OK",
    "message": "run CollectHardwareInfo success.",
    "data": {
        "raid_driver": "megacli64",
        "bmc_address": "10.208.16.165",
        "interfaces": [{
            "product": "0x1572",
            "vendor": "0x8086",
            "name": "eth0",
            "bus_info":"0000:01:00.0",//总线位置信息。0000:01:00代表着相应设备编号可用于标识网卡，.0是设备函数，通过.分割前面的可代表网卡设备
            "has_carrier": true,
            "ipv4_address": "10.208.14.174",//ip地址
            "switch_manage_ip": "10.208.13.11",//交换机ip
            "biosdevname": "p3p1",
            "mac_address": "3c:fd:fe:51:f2:30",//mac地址
            "switch_index": "Ten-GigabitEthernet1/0/7"//交换机端口号
        }, {
            "product": "0x1572",
            "vendor": "0x8086",
            "name": "eth1",
             "bus_info":"0000:01:00.1",
            "has_carrier": true,
            "ipv4_address": "10.208.14.175",
            "switch_manage_ip": "10.208.13.11",
            "biosdevname": "p3p2",
            "mac_address": "3c:fd:fe:51:f2:31",
            "switch_index": "Ten-GigabitEthernet1/0/8"
        }],
        "disks": [],
        "boot": {
            "current_boot_mode": "bios"
        },
        "system_vendor": {
            "serial_number": "J33H4YT",
            "product_name": "System x3650 M5: -[8871AC1]- ((none))",
            "manufacturer": "LENOVO"
        },
        "memory": {
            "physical_mb": 131072,
            "total": 134743052288
        },
        "cpu": {
            "count": 32,
            "frequency": "3000.0000",
            "flags": ["fpu", "vme", "de", "pse", "tsc", "msr", "pae", "mce", "cx8", "apic", "sep", "mtrr", "pge", "mca", "cmov", "pat", "pse36", "clflush", "dts", "acpi", "mmx", "fxsr", "sse", "sse2", "ss", "ht", "tm", "pbe", "syscall", "nx", "pdpe1gb", "rdtscp", "lm", "constant_tsc", "arch_perfmon", "pebs", "bts", "rep_good", "nopl", "xtopology", "nonstop_tsc", "aperfmperf", "eagerfpu", "pni", "pclmulqdq", "dtes64", "monitor", "ds_cpl", "vmx", "smx", "est", "tm2", "ssse3", "sdbg", "fma", "cx16", "xtpr", "pdcm", "pcid", "dca", "sse4_1", "sse4_2", "x2apic", "movbe", "popcnt", "tsc_deadline_timer", "aes", "xsave", "avx", "f16c", "rdrand", "lahf_lm", "abm", "3dnowprefetch", "epb", "cat_l3", "cdp_l3", "invpcid_single", "intel_ppin", "intel_pt", "tpr_shadow", "vnmi", "flexpriority", "ept", "vpid", "fsgsbase", "tsc_adjust", "bmi1", "hle", "avx2", "smep", "bmi2", "erms", "invpcid", "rtm", "cqm", "rdt_a", "rdseed", "adx", "smap", "xsaveopt", "cqm_llc", "cqm_occup_llc", "cqm_mbm_total", "cqm_mbm_local", "dtherm", "ida", "arat", "pln", "pts"],
            "model_name": "Intel(R) Xeon(R) CPU E5-2620 v4 @ 2.10GHz",
            "architecture": "x86_64"
        },
        "platform":{
            "machine":"x86_64" //系统架构 x86_64、aarch64
        }
    },
    "sn": "J33H4YT"
}
```

### 收集硬盘信息

```json
{
    "action" : "CollectDiskLocations",
    "sn" : "0",
    "raid_driver" : "" //可选。传参：sas2ircu，sas3ircu，megacli64，storcli64，perccli64，arcconf。不传参数则只可采集已适配机型，如果未适配则采集失败
}
```

response
``` json
{
    "action": "CollectDiskLocations",
    "status": "OK",
    "message": "run CollectDiskLocations success.",
    "data": {
        "controllers": [{
            "disks": [{
                "enclosure": 8,
                "slot": 0,
                "size": 278.464,
                "size_unit": "GB",
                "pd_type": "SAS",
                "media_type": "SDD"
            }, {
                "enclosure": 8, //int
                "slot": 1,// int
                "size": 278.464, //硬盘大小 float
                "size_unit": "GB", //硬盘大小单位 GB TB ,1024进制
                "pd_type": "SAS", //接口类型
                "media_type": "SDD" //磁盘类型
            }],
            "adapter_id": 0,
        }],
        "nvme":{
            "devices" : [// nvme无须做raid enclosure，slot都设置为0即可
                {
                "device_path" : "/dev/nvme0n1",
                "index" : 0,
                "serial" : "BTLJ133609V74P0VGN",
                "size" : 4.000787030016,
                "size_unit": "TB",
                "pd_type": "NVME",
                "media_type": "SDD"
                }
            ]
        },
        "devices":[
            {
                "device_path": "/dev/sda",//此盘符可能是直通盘符，也可能是做了raid后组成的盘符
                "media_type": "SSD", 
                "serial": "600062b20313f5c029d090b717000e67",
                "size": 798.999183360,
                "size_unit": "GB",
            },
            {
                "device_path": "/dev/sdb",
                "media_type": "SSD",
                "serial": "600062b20313f5c029d090b717f8cb7d",
                "size": 999.9999827968,
                "size_unit": "GB",
            }
        ]
    },
    "sn": "J33H4YT"
}
```

### make raid

```json
{
    "action": "MakeRaid",
    "sn": "0",
    "raid_datas": [{
        "raid_driver": "megacli64", //可选。sas2ircu，sas3ircu，megacli64，storcli64，perccli64，arcconf。不传时，系统自动判断使用哪个raid driver
        "adapter_id": 0,
        "volumes": [{
                "volume_id": "v-lloenltsstdkqraogf5cqconvbsg",//uuid 用于response数据返回时进行数据关联
                "is_root_device": true, // 可选。是否是系统盘，true表示是系统盘。不传时根据盘符由小到大排序，第一块为系统盘
                "raid_level": "RAID5", //raid类型:NORAID,RAID0,RAID1,RAID5,RAID10 
                "physical_disks": [{
                        "enclosure": 8,
                        "slot": 0,
                    },
                    {
                        "enclosure": 8,
                        "slot": 1
                    },
                    {
                        "enclosure": 8,
                        "slot": 2
                    },
                    {
                        "enclosure": 8,
                        "slot": 3
                    },
                    {
                        "enclosure": 8,
                        "slot": 4,
                    }
                ]
            },
            {
                "volume_id": "v-ptuu7zhtmqydeutrvciyfbjv3pjw",
                "is_data_device": true, //[字段废弃] 可选。是否是数据盘，true表示是数据盘。不传时根据盘符由小到大排序，第二块为数据盘
                "raid_level": "RAID1",
                "physical_disks": [{
                        "enclosure": 8,
                        "slot": 5
                    },
                    {
                        "enclosure": 8,
                        "slot": 6
                    }
                ]
            }
        ]
    }]
}
```

response
```json
{
	"sn": "J33RFRX",
	"action": "MakeRaid",
	"code": 200,
	"message": "Run MakeRaid success.",
	"data": {
		"raid_datas": [{
			"adapter_id": 0,
			"volumes": [{
				"volume_id": "v-lloenltsstdkqraogf5cqconvbsg",
				"disk_hints": {
					"name": "/dev/sda",
					"model": "ServeRAID M5210",
					"size": 298999349248,
					"rotational": true,
					"wwn": "0x600605b00d89bb80",
					"serial": "600605b00d89bb802de2e524599e42e9",
					"vendor": "IBM",
					"wwn_with_extension": "0x600605b00d89bb802de2e524599e42e9",
					"wwn_vendor_extension": "0x2de2e524599e42e9",
					"hctl": "0:2:0:0",
					"by_path": "/dev/disk/by-path/pci-0000:0b:00.0-scsi-0:2:0:0"
				}
			}, {
				"volume_id": "v-ptuu7zhtmqydeutrvciyfbjv3pjw",
				"disk_hints": {
					"name": "/dev/sdb",
					"model": "ServeRAID M5210",
					"size": 798999183360,
					"rotational": true,
					"wwn": "0x600605b00d89bb80",
					"serial": "600605b00d89bb802de2e52659c204d8",
					"vendor": "IBM",
					"wwn_with_extension": "0x600605b00d89bb802de2e52659c204d8",
					"wwn_vendor_extension": "0x2de2e52659c204d8",
					"hctl": "0:2:1:0",
					"by_path": "/dev/disk/by-path/pci-0000:0b:00.0-scsi-0:2:1:0"
				}
			}]
		}]
	},
	"status": "OK"
}
```

### clean raid

```json
{
    "action": "CleanRaid",
    "sn": "0",
    "raid_datas": [{
        "raid_driver": "megacli64", //可选。sas2ircu，sas3ircu，megacli64，storcli64，perccli64，arcconf。不传时，系统自动判断使用哪个raid driver
        "adapter_id": 0
    }]
}


```
创建分区
```json
{
    "action": "MakePartitions",
    "sn": "0",
    "version": "2.0",
    "boot_mode": "uefi",//可选，默认:空。不传此参数时，agent将自动判断模式
    "auto_create_efi": true, //可选，默认:true。uefi模式下自动创建efi分区
    "auto_create_bios_grub": true,//可选，默认:true。bios模式下自动创建bios_grub分区
    "volumes": [{
            "keep_data": false,//true:保留数据， false:不保留数据。新装机器时为false
            "is_root_device": true,//可选。相对于disk_hints优先级更高
            "disk_hints":{"serial": "600062b10030a4802c961499f2a09445"},//可选。可指定特定的盘符进行分区
            "disk_label": "gpt",
            "partitions": [
                {
                    "id": "ipt-haxcsxs570fhqew36zuuilks5w6v",//表uuid
                    "part_type":"esp",//可选
                    "fs_type":"fat32",//可选
                    "boot_flag": "boot",//可选。bios模式:引导分区需要设置此参数(有boot分区时是)
                    "number": 1,
                    "size": 512 //单位Mb
                },
                {
                    "id": "ipt-1bt811r0feavi4sfln5nyyz8qgvg",
                    "number": 2,
                    "size": 102400
                }
            ]
        },
        {
            "keep_data": true,  //true:保留数据， false:不保留数据。新装机器时为false
            "is_data_device": true,//[字段废弃]可选
            "disk_hints":{"serial": "600062b10030a4802c961499f2a09445"},
            "disk_label": "gpt",
            "partitions": [
                 {
                    "id": "ipt-ptuu7zhtmqydeutrvciyfbjv3pjw",
                    "number": 1,
                    "size": 512
                },
                {
                    "id": "ipt- llnm71kosh9l7yvwggungl2uawwx",
                    "number": 2,
                    "size": 102400
                }
            ]
        }
    ]
}
```
备注：重装系统保留数据时：只传系统盘分区，不要传数据盘等相关分区。格式化分区也只传系统盘相关。

格式化分区
```json
{
    "action": "FormatPartitions",
    "sn": "0",
    "version":"2.0",
    "partitions": [
        {
            "volume": "ipt-1bt811r0feavi4sfln5nyyz8qgvg",// 此参数来自，MakePartitions的 "id": "ipt-1bt811r0feavi4sfln5nyyz8qgvg" 或MakeSoftRaid的 "id": "ipg-yhzcyaxnw6rtieorpgfdnbfshfh9"
            "is_root": true, //是否是根分区
            "fs_type": "ext4",
            "label": "l_root"
        }
    ]
}
```

### 挂载目录
```json
{
    "action": "MountPartitions",
    "sn": "0",
    "version":"2.0",
    "auto_mount_efi":true, //可选，自动挂载efi分区
    "mounts": [{
            "is_root_device": true, 
            "disk_hints":{"serial": "600062b10030a4802c961499f2a09445"},//盘符特性信息。确定盘符信息由is_root_device或disk_hints确定is_root_device优先级更高
            "label":"l_root",//label只在每个盘符下唯一，所以同时需要确定盘符的相应信息
            "options": "errors=remount-ro",
            "mountpoint": "/"
        },
        {
            "is_data_device": true, //[字段废弃]
            "disk_hints":{"serial": "600062b10030a4802c961499f2a09445"},//盘符特性信息。确定盘符信息由is_root_device或disk_hints确定is_root_device优先级更高
            "label":"l_export",
            "options": "errors=remount-ro",
            "mountpoint": "/export"
        }
    ]
}
```

### 写镜像

```json
{
    "action" : "WriteImage",
    "sn" : "0",
    "url" : "{host}/ubuntu-trusty.qcow2",
    "format" : "qcow2",
    "hash" : "d3a470cc4d53fa89026ad7433964f769",
    "filename" : "ubuntu-trusty.qcow2"
}
```

```json
{
  "sn": "0",
  "action": "WriteImage",
  "code": 200,
  "message": "Run WriteImage success.",
  "data": {
    "root_device_hints": {
      "name": "/dev/sdp",
      "model": "LogicalDrv 0",
      "size": 600070840320,
      "rotational": true,
      "wwn": "0x600508b1001c3cfe",
      "serial": "600508b1001c3cfe2d5e6e4568e31521",
      "vendor": "SmartIO",
      "wwn_with_extension": "0x600508b1001c3cfe2d5e6e4568e31521",
      "wwn_vendor_extension": "0x2d5e6e4568e31521",
      "hctl": "0:0:0:0",
      "by_path": "/dev/disk/by-path/pci-0000:01:00.0-scsi-0:0:0:0"
    }
  },
  "status": "OK"
}
```

# 写Tar镜像
```json
{
    "action": "WriteImageTar",
    "sn": "0",
    "url": "{host}/v1.3.0-centos-6.6-2020112716.tar.xz",
    "format": "tar",
    "hash": "281e40975e2091b14c63586bb8a9301e",
    "filename": "v1.3.0-centos-6.6-2020112716.tar.xz"
}
```

```json
{
  "sn": "0",
  "action": "WriteImageTar",
  "code": 200,
  "message": "Run WriteImageTar success.",
  "data": {
    "root_device_hints": {
      "name": "/dev/sdp",
      "model": "LogicalDrv 0",
      "size": 600070840320,
      "rotational": true,
      "wwn": "0x600508b1001c3cfe",
      "serial": "600508b1001c3cfe2d5e6e4568e31521",
      "vendor": "SmartIO",
      "wwn_with_extension": "0x600508b1001c3cfe2d5e6e4568e31521",
      "wwn_vendor_extension": "0x2d5e6e4568e31521",
      "hctl": "0:0:0:0",
      "by_path": "/dev/disk/by-path/pci-0000:01:00.0-scsi-0:0:0:0"
    }
  },
  "status": "OK"
}
```

### 设置 root password

```json
{
    "action" : "SetPassword",
    "sn" : "0",
    "password": "beijing",
    "username": "root" //windows时是Administrator
}
```

### 清理所有磁盘数据

```json
{
    "action" : "CleanBlockDevice",
    "sn" : "0"
}
```

### 初始化设备信息

```json

{
	"action": "InitNode",
    "sn": "0",
	"node_data": {
		"root_device_hints": {"serial": "600062b10030a4802c961499f2a09445"},//可选。只需要传serial一个参数
        "data_device_hints": {"name": "/dev/nvme0n2"},//[字段废弃]可选。name字段预留
		"os_type": "Ubuntu",
		"os_version": "20.04"
	}
}

```

### windows设置网络bond

```json
{
	"action": "SetNetworkWindows",
    "sn": "0",
	"network": {
		"services": [{
			"type": "dns",
			"address": "172.16.16.16"
		}, {
			"type": "dns",
			"address": "10.16.16.16"
		}],
		"networks": [{
            	"link": "bond0",
				"type": "ipv4",
                "ip_address": "10.209.10.234",
				"netmask": "255.255.255.248",
				"routes": [{
                    "network": "0.0.0.0",
					"netmask": "0.0.0.0",
					"gateway": "10.209.10.233"
				}]
			},
			{
				"link": "bond0",
				"type": "ipv6",
				"ip_address": "fd01:0:1:0:2::4",
                "netmask":120,
				"routes": [{
                    "network": "::",
					"netmask": "::",
					"gateway": "fd01:0:1:0:1"
				}]
			}
		],
		"links": [{
			"ethernet_mac_address": "0c:42:a1:d8:64:6a",
			"type": "phy",
			"id": "eth0",
			"mtu": 1500
		}, {
			"ethernet_mac_address": "0c:42:a1:d8:64:6b",
			"type": "phy",
			"id": "eth1",
			"mtu": 1500
		}, {
			"bond_miimon": 100,
			"ethernet_mac_address": "0c:42:a1:d8:64:6a",
			"bond_mode": "802.3ad",
			"bond_links": ["eth0", "eth1"],
			"type": "bond",
			"id": "bond0"
		}]
	}
}
```
### windows设置网络，网口非聚合
```json
{
	"action": "SetNetworkWindows",
    "sn": "0",
	"network": {
		"services": [{
			"type": "dns",
			"address": "172.16.16.16"
		}, {
			"type": "dns",
			"address": "10.16.16.16"
		}],
		"networks": [{
			"ip_address": "10.209.10.234",
			"type": "ipv4",
			"netmask": "255.255.255.248",
			"link": "eth0",
			"routes": [{
                "network": "10.209.10.234",
				"netmask": "255.255.255.248",
				"gateway": "10.209.10.233"
			}, {
                "network": "0.0.0.0",
				"netmask": "0.0.0.0",
				"gateway": "10.209.10.233"
			}
            ]
		},
        {
			"ip_address": "2403:1EC0:8549:60C0::4",
			"type": "ipv6",
			"netmask": "64",
			"link": "eth0",
			"routes": [{
                "network": "::",
				"netmask": "::",
				"gateway": "2403:1EC0:8549:60C0::1"
			}
            ]
		}
        ],
		"links": [{
			"ethernet_mac_address": "0c:42:a1:d8:64:6a",
			"type": "phy",
			"id": "eth0",
			"mtu": 1500
		}]
	}
}
```

### linux 设置 network bond
```json
{
    "action": "SetNetwork",
    "sn": "0",
    "network":  {
        "ethernets": {
            "eth0": {//网卡key 使用网卡名称
                "match": {
                    "macaddress": "38:68:dd:04:8e:68" //网卡mac地址
                },
                "set-name": "eth0" //网卡名称
            },
            "eth1": {
                "match": {
                    "macaddress": "38:68:dd:04:8e:69"
                },
                "set-name": "eth1"
            }
        },
        "bonds": {
            "bond0": {
                "macaddress": "38:68:dd:04:8e:68", //eth0 对应网卡mac
                "addresses": ["10.0.0.2/16", "2402:db40:51ba:f1::3/64"], //私有地址ip+子网cidr 支持ipv4和ipv6同时设置
                "gateway4": "10.0.0.1", //子网网关
                "gateway6":"2402:db40:51ba:af::1",  //设置ipv6地址时必传
                "nameservers": {
                    "addresses": ["114.114.114.114", "8.8.8.8"] //dns地址
                },
                "interfaces": ["eth0", "eth1"], //关联的网卡key
                "parameters": {
                    "mode": "802.3ad", //当前传此固定值
                    "mii-monitor-interval": 100 //整型 当前传此固定值
                }
            }
        },
        "version": 2 //整型
    }
}

```

### linux 设置 network 非聚合
#### 说明
- 双网口vpc时，辅网口不要设置gateway4，否则cloud-init会为其设置默认路由0.0.0.0与主网卡冲突
- 双网口基础网络环境，开启外网时(eth1是外网)，此时eth0不设置gateway4，eth1设置gateway4
```json

{
    "action": "SetNetwork",
    "sn": "0",
    "network":  {
        "ethernets": {
            "eth0": {//网卡key 使用网卡名称
                "match": {
                    "macaddress": "38:68:dd:04:8e:68" //网卡mac地址
                },
                "set-name": "eth0", //set-name 是中划线不是下划线
                "addresses": ["10.0.0.2/16", "2402:db40:51ba:f1::3/64"], //私有地址ip+子网cidr 支持ipv4和ipv6同时设置
                "netmask":"255.255.0.0",
                "gateway4": "10.0.0.1",
                "gateway6": "2402:db40:51ba:af::1", //设置ipv6地址时必传
                "nameservers": {
                    "addresses": ["103.224.222.222", "103.224.222.223"]
                },
                "routes":[
                         {
                        "to":"0.0.0.0/0", //固定值全局路由
                        "via":"10.0.0.1" //子网网关
                    },
                    {
                        "to":"10.0.0.0/16",//vpc网段
                        "via":"10.0.0.1" //子网网关
                    }
                ]
            },
            "eth1": {
                "match": {
                    "macaddress": "38:68:dd:04:8e:69" //网卡mac地址
                },
                "set-name": "eth1",
                "addresses": ["192.168.0.4/27", "2402:db40:51ba:af::3/64"], //ipv6辅网卡无需设置gateway6
                "netmask":"255.255.255.224",
                "nameservers": {
                    "addresses": ["103.224.222.222", "103.224.222.223"]
                },
                "routes":[
                    {
                        "to":"192.168.0.0/26", //vpc网段
                        "via":"192.168.0.1"  //子网网关
                    }
                ]
            }
        },
        "version": 2 //整型
    }
}
```

### 设置 metadata
```json
{
    "action": "SetMetaData",
    "sn": "0",
    "meta_data": {
        "instance-id": "cps-s4r3e3exknzbo8wxr2nkdwuz5v94", //required 实例id
        "local-hostname": "host001", //required 主机名 
        "public-keys": ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC0vbep6l/mGZh8XItDj2oaldOPYiZXb1MkQlpC1zhda9YXzdFoufKrqFH4o5Gch6ve7/wncEUarnGb49IEjGdN77ooLQ+wWTw1HnJCKy1CPYVFL1oMYt2+89gh9HlyjbpyHnlHXG2hvklYvKVNtqQqDTwOM8794/W/Ik3H8iES8dCrD1+XoljSxDSaM6ou/h/W7tVflLI5TN74I/uxdOdRzbisB662HZXIXysGR6jMvUpvbENolB7j1S1ttWGs2sNKOcpExu76EcvOcLoTqdwbZ0fOG/q5xsgcTuPtyECzOK336u+VsBnp1vxdLzLcQi8Ry1+2NrnAN9U4pW0gEdab"]//公钥值 不需设置时，不要传此参数
    }
}
```
### 设置 userdata
```json
{
    "action" : "SetUserData",//不设置用户自定义脚本时，Linux系统也需要传此指令，user_data传空字符串。Nocloud必须要有user_data文件才认为读取数据正常。Windows系统可不传此指令
    "sn" : "0",
    "user_data":"IyEvYmluL2Jhc2gKZWNobyAndGVzdCB1c2VyIGRhdGEnID4gL3Jvb3QvZmlsZQ==" //用户自定义数据
}
```

### 设置 vendordata
```json
{
    "action": "SetVendorData",//此指令暂时可不使用，未来扩展使用
    "sn": "0",
    "vendor_data": {
        "local-hostname": "hostname001"
     }
}
```

### 设置 cloudinit 配置文件
```json
{
    "action" : "SetCloudinitConf",
    "sn" : "0",
    "ssh_pwauth":"yes"   //yes:启用密码登陆  //no禁用密码登陆
}
```

#### raid_driver 可用值
- sas2ircu
- sas3ircu
- megacli64
- storcli64
- perccli64
- arcconf

### TAR镜像，装机模版
#### 不做raid，新装
- DHCPConfigAddHost
- SetPXEBoot
- Ping
- InitNode
- MakePartitions
- FormatPartitions
- MountPartitions
- WriteImageTar
- SetPassword
- SetCloudinitConf
- SetMetaData
- SetNetwork
- SetUserData
- DHCPConfigDelHost
- SetDISKBoot
- PowerCycle
#### 不做raid，重装
- DHCPConfigAddHost
- SetPXEBoot
- Ping
- InitNode
- MakePartitions
- FormatPartitions   1.保留数据时，只传系统盘相关信息；2.不保留数据时，同新装传全量数据
- MountPartitions
- WriteImageTar
- SetPassword
- SetCloudinitConf
- SetMetaData
- SetNetwork
- SetUserData
- DHCPConfigDelHost
- SetDISKBoot
- PowerCycle
#### 做raid，新装
- DHCPConfigAddHost
- SetPXEBoot
- Ping
- MakeRaid
- InitNode
- MakePartitions
- FormatPartitions
- MountPartitions
- WriteImageTar
- SetPassword
- SetCloudinitConf
- SetMetaData
- SetNetwork
- SetUserData
- DHCPConfigDelHost
- SetDISKBoot
- PowerCycle
#### 做raid，重装
- DHCPConfigAddHost
- SetPXEBoot
- Ping
- InitNode
- MakePartitions
- FormatPartitions   1.保留数据时，只传系统盘相关信息；2.不保留数据时，同新装传全量数据
- MountPartitions
- WriteImageTar
- SetPassword
- SetCloudinitConf
- SetMetaData
- SetNetwork
- SetUserData
- DHCPConfigDelHost
- SetDISKBoot
- PowerCycle
#### 修改密码
- DHCPConfigAddHost
- SetPXEBoot
- Ping
- InitNode
- SetPassword
- DHCPConfigDelHost
- PowerOff
### QCOW2,TAR通用模版
#### 不做raid，销毁机器
- DHCPConfigAddHost
- SetPXEBoot
- Ping
- CleanBlockDevice
- DHCPConfigDelHost
- PowerOff
#### 做raid，销毁机器
- DHCPConfigAddHost
- SetPXEBoot
- Ping
- CleanBlockDevice
- CleanRaid
- DHCPConfigDelHost
- PowerOff
### WINDOWS QCOW2镜像，装机模版
SetMetaData，部分传参变化
- hostname变更为local-hostname
- uuid变更为instance-id
#### 不做raid，新装，重装相同
- DHCPConfigAddHost
- SetPXEBoot
- Ping
- InitNode
- WriteImage
- SetPassword
- SetMetaData
- SetNetworkWindows
- SetUserData(设置自定义脚本设置需要此指令，不设置时不需要)
- DHCPConfigDelHost
- SetDISKBoot
- PowerCycle
#### 做raid，新装，重装相同
- DHCPConfigAddHost
- SetPXEBoot
- Ping
- MakeRaid
- InitNode
- WriteImage
- SetPassword
- SetMetaData
- SetNetworkWindows
- SetUserData(设置自定义脚本设置需要此指令，不设置时不需要)
- DHCPConfigDelHost
- SetDISKBoot
- PowerCycle
#### 修改密码
- DHCPConfigAddHost
- SetPXEBoot
- Ping
- InitNode
- SetPassword
- DHCPConfigDelHost
- PowerOff


## 备注
对于特殊机器新增TFTPConfigAddGrub、TFTPConfigDelGrub指令
```
TFTPConfigDelGrub、TFTPConfigAddGrub放在DHCPConfigAddHost前面
TFTPConfigDelGrub放在DHCPConfigDelHost后面
```

