# Table of Contents

- [机型录入](#机型录入)
- [设备录入](#设备录入)
- [厂商和硬盘信息录入](#厂商和硬盘信息录入)
- [镜像录入](#镜像录入)
- [机型和镜像关联](#关联镜像)
- [机型和镜像关联移除](#移除镜像)
- [raid录入](#raid录入)
- [机型和raid关联](#机型和raid关联)
- [机型和raid解关联](#机型和raid解关联)
- [机型管理列表](#机型管理列表)
- [查询设备类型列表](#查询设备类型列表)
- [查询镜像列表](#查询镜像列表)
- [查询 RAID 列表](#查询raid列表)
- [查询用户列表](#查询用户列表)
- [查询用户详情](#查询用户详情)
- [查询当前登录用户](#查询当前登录用户)
- [查询实例列表-已废弃](#查询实例列表-已废弃)
- [导出实例列表-已废弃](#导出实例列表-已废弃)
- [查询设备实例列表](#查询设备实例列表)
- [查询设备实例详情](#查询设备实例详情)
- [导出设备实例列表](#导出设备实例列表)
- [查询设备库存](#查询设备库存)
- [导出设备库存](#导出设备库存)
- [设备上架售卖](#设备上架)
- [设备下架](#设备下架)
- [设备回收](#设备回收)
- [售卖状态查询](#售卖状态查询)
- [实例状态查询](#实例状态查询)

## changelog

| 版本 | 变更日期   | 变更内容 |  变更人 | 
| ----| --------- | ------- | ------ |
| v1.0.0 | 2021-09-18 | 初始化 | 闵平(minping@jd.com)|


## 约定
除了GET请求外，其他格式请求的requestbody格式一律为 Content-Type=application/json


_通用数据结构:

```json
{
    "requestId": "xxx",
    "data": {

    },
    "error": {
        "code": 500,
        "message": "invalid xxx parameter",
        "status": "InvalidParameter",
    }
}
```

_后续接口定义中的内容全部都是 result 属性中的结构_

### 机型录入

##### url
```http
/v1/offline/CreateBmClazz
```

##### method

```http
POST
```

##### 请求参数

```json
{
	"region" : "cn-north-1", //required,
	"deviceType" :  "cps.c.perf1", //required
	"nameEN" : "Compute Performance Ⅰ", //required
	"nameZH" : "计算效能型Ⅰ", //required"
	"family" : "compute", //设备所属规格系列：存储storage 计算compute GPUgpu
	"canRaid" : true, //能否做RAID：true可以, false不可以
	"isSale" : true, //是否可以买，true可以, false不可以,废弃
	"cpuAmount" : 1, //cpu数量
	"cpuCores" : 1,  //required,（单个）cpu核心数
	"cpuManufacturer" : "Intel", //required, cpu厂商
	"cpuModel" : "E5-2683V4", //required cpu处理器型号
	"cpuFrequency" : "2.1", //required cpu频率(G)
	"memAmount" : "8",  // required 内存条数量
	"memCapacity" : "32", // required 单个内存容量(GB)
	"memModel" : "DDR4", //required 
	"ifAmount" : "2", //required 网卡数量
	"ifRate" : "10", //required 网卡接口传输速率（GE）
	"gpuAmount" : "4", //GPU数量
	"gpuManufacturer" : "NVIDIA", //GPU厂商
	"gpuModel" : "V100", //GPU处理器型号
	"systemVolumeAmount" : "6",
	"systemVolumeSize" : "E5-2683V4",
	"systemVolumeModel" : "SSD",
	"dataVolumeAmount" : "6",
	"dataVolumeSize" : "6000",
	"dataVolumeModel" : "SATA",
	"diskAmount" : "4",
	"source" : "common", //required 来源 common|customize
}
```

##### 响应示例

```json
{
  "success" : true
}
```

### 设备录入

##### url

```http
v1/offline/createDevices
```

##### method

```http
POST
```

##### 请求参数

```json
{
  "devices" : [
      {
          "region" : "cn-north-1",
          "az" : "cn-north-1a",
          "sn" : "1B3MXP2", // required
          "device_type" : "cps.c.normal", //required
          "u_position" : "2",//槽位
          "cabinet" : "1",//机柜
          "ilo_ip":"10.208.11.131", //required 带外ip
      }
  ]
}
```

##### 响应示例

{
  "success" : true
}

### 厂商和硬盘信息录入

##### url

```http
v1/offline/collectManufacturerAndDisk
```

##### method

```http
POST
```

##### 请求参数

```json
{
  "sns" : [
    "DP50CP2",
    "1B3MXP2"
  ],
  "network_type" : "basic"
}
```

##### 响应示例

```json
{
  "success" : true
}
```

### 镜像录入

##### url

```http
v1/offline/createImage
```

##### method

```http
POST
```

##### 请求参数

```json
{
  "name_en" : "CentOS 6.6 64-bit", // required 长度介于1-200
	"name_zh" : "CentOS 6.6 64位", // required 长度介于1-200
	"version" : "1.0", // required
	"os_id" : "o-1vc0l7zpnrt1rpnssudanmq01jbv", // required
	"format" : "qcow2", // required qcow2|tar
	"filename" : "v1.3.0-centos-6.6-2019081310.qcow2", // required
	"url" : "{host}/v1.3.0-centos-6.6-2019081310.qcow2", // required
	"hash" : "2058a93ca59798b96f446ef306bcf690",// required
	"source" "common" // required common|customize
}
```

##### 响应示例

```json
{
  "success" : true
}
```

### raid录入

常见的raid方案有NORAID， RAID0， RAID1， RAID10， RAID51H， soft RAID1等，这些信息一般开始就录入进表了，后面基本不会动。故暂无接口

### 机型和raid关联

##### url

```http
v1/offline/deviceTypes/associatedRaid
```

##### method

```http
POST
```

##### 请求参数

```json
{
  "raid_id" : "r-sfykwibcvcyztuez3mk1klyha9nn", //required
	"device_type" : "cps.gpu.1", ////required
	"volume_type" : "data", ////required data|system
	"volume_detail" : "300GB (2*300GB SAS)", ////required
}
```

##### 响应示例

```json
{
  "success" : true
}
```

### 机型和raid解关联

##### url

```http
v1/offline/deviceTypes/dissociatedRaid
```

##### method

```http
DELETE
```

##### 请求参数

```json
{
  "raid_id" : "r-sfykwibcvcyztuez3mk1klyha9nn", //required
	"device_type" : "cps.gpu.1", //required
	"volume_type" : "data", //required data|system
	"volume_detail" : "300GB (2*300GB SAS)", //
}
```

##### 响应示例

```json
{
  "success" : true
}
```



### 查询设备类型列表

##### url

```http
v1/offline/deviceTypes
```

##### method

```http
GET
```

##### 请求参数

参数名|类型|是否必须|默认值|描述
---|---|---|---|---
region|String|否| |region
deviceType|String|否| |实例类型


##### 请求响应

```json
{
  "deviceTypes": [
    {
      "deviceType": "cps.gpu.4", //规格
      "name": "GPU IV", //实例类型
      "region": "cn-east-2", //地域
      "family": "gpu", //系列
      "cpuConcise": "12核", //cpu概述
      "cpuDetail": "2*Intel 2683V4（12核 2.2G）", //cpu详情
      "memConcise": "256G", //内存概述
      "memDetail": "256G（8*32G）", //内存详情
      "ifConcise": "2*10GE", //网卡概述
      "ifDetail": "2*10GE", //网卡详情
      "gpuConcise": "NVIDIA P40 * 4", //gpu概述
      "gpuDetail": "NVIDIA P40 * 4", //gpu详情
      "systemDiskAmount": 2, //系统盘数量
      "systemDiskSize": 300, //系统盘大小
      "systemDiskModel": "SAS", //系统盘类型
      "dataDiskAmount": 1, //数据盘数量
      "dataDiskSize": 4000, //数据盘大小
      "dataDiskModel": "SSD NVME", //数据盘类型
      "source": "common" //来源：通用common,定制customize,自定义user_defined
    }
  ]
}
```

### 机型管理列表

##### url

```http
v1/offline/deviceTypes/plus
```

##### method

```http
GET
```

##### 请求参数

参数名|类型|是否必须|默认值|描述
---|---|---|---|---
region|String|否| |region
deviceType|String|否|cps.s.normal |实例类型
pageNumber|int|否| 1|第 几 页
pageSize|int|否| 20|每页大小


##### 请求响应

```json
{
  "deviceTypes": [
    {
      "deviceType": "cps.gpu.4", //规格
      "name": "GPU IV", //实例类型名称
      "region": "cn-east-2", //地域
      "family": "gpu", //系列
      "cpuConcise": "12核", //cpu概述
      "cpuDetail": "2*Intel 2683V4（12核 2.2G）", //cpu详情
      "memConcise": "256G", //内存概述
      "memDetail": "256G（8*32G）", //内存详情
      "ifConcise": "2*10GE", //网卡概述
      "ifDetail": "2*10GE", //网卡详情
      "gpuConcise": "NVIDIA P40 * 4", //gpu概述
      "gpuDetail": "NVIDIA P40 * 4", //gpu详情
      "systemDiskAmount": 2, //系统盘数量
      "systemDiskSize": 300, //系统盘大小
      "systemDiskModel": "SAS", //系统盘类型
      "dataDiskAmount": 1, //数据盘数量
      "dataDiskSize": 4000, //数据盘大小
      "dataDiskModel": "SSD NVME", //数据盘类型
      "source": "common", //来源：通用common,定制customize,自定义user_defined
      "osCount": 3, //关联镜像数
      "dataRaids": ["RAID0", "RAID1"],
      "sysRaids": ["RAID0", "RAID1"],
      "serviceCode": "cps" //所属产品，cps（物理服务器）/edcps（分布式物理服务器）
    }
  ],
  "pageNumber": 1,
  "pageSize": 20,
  "totalCount": 12
}
```

### 查询镜像列表

##### url

```http
v1/offline/oss
```

##### method

```http
get
```

##### 请求参数

参数名|类型|是否必须|默认值|描述
---|---|---|---|---
region|String|否| |region
deviceType|String|否| cps.s.normal|实例类型（规格）
imageType|String|否| standard|镜像类型，标准镜像，非必填
osType|String|否| CentOS|操作系统：CentOS，Ubuntu，非必填
source|String|否| common|来源：通用 common,定制 customize,自定义user_defined，非必填
pageNumber|int|否| 1|第 几 页
pageSize|int|否| 20|每页大小

##### 响应参数

```json
{
  "oss": [
    {
      "osTypeId": "i-abcdadoiadld", //镜像id
      "osName": "CentOS7.4 64位", //操作系统名称
      "deviceType": "cps.c.normal", //实例类型（规格）
      "osType": "CentOS", //操作系统：CentOS,Ubuntu
      "osVersion": "7.4", //版本
      "source": "common" //来源：通用common,定制customize,自定义user_defined
    }
  ],
  "pageNumber": 1,
  "pageSize": 20,
  "totalCount": 12
}
```

### 关联镜像

##### url

```http
v1/offline/deviceTypes/associatedImage
```

##### method

```http
POST
```

##### 请求参数
```json
{
  "region": "cn-north-1", //必填
  "deviceType": "cps.s.normal", //必填，实例类型（规格）
  "osTypeIds": ["i-abcdadoiadld","i-adfaladida"], //镜像id,必填
  "pin": "test01" //给用户关联镜像时必填，其他情况不填
}
```

##### 响应参数

```json
{
  "success": true
}
```

### 移除镜像

##### url

```http
v1/offline/deviceTypes/dissociatedImage
```

##### method

```http
Delete
```

##### 请求参数

```json
{
  "region": "cn-north-1", //必填
  "deviceType": "cps.s.normal", //必填，实例类型（规格）
  "osTypeId": "i-abcdadoiadld", //镜像id，必填
  "pin": "test01" //给用户移除镜像时必填，其他情况不填
}
```

##### 响应参数

```json
{
  "success": true
}
```

### 查询raid列表

##### url

```http
v1/offline/raids
```

##### method

```http
get
```

##### 请求参数

参数名|类型|是否必须|默认值|描述
---|---|---|---|---
deviceType|String|否| cps.s.normal|实例类型（规格）


##### 响应参数

```json
{
  "raids": [
    {
      "raidTypeId": "r-lowdd8nf8y7v6jk5ocr5paqhvhre", //raid id
      "raidType": "RAID1", //raid类型名称
      "deviceType": "cps.c.normal", //实例类型
      "volumeType": "system", //磁盘类型：system,data
      "volumeDetail": "6720GB (14*960GB SATA)",
      "description": "描述"
    }
  ]
}
```

### 查询用户列表

##### url

```http
v1/offline/users
```

##### method

```http
get
```

##### 请求参数

参数名|类型|是否必须|默认值|描述
---|---|---|---|---
pageNumber|int|否| 1|第 几 页
pageSize|int|否| 20|每页大小

- pageNumber=1 第 1 页
- pageSize=20 每页大小

##### 响应参数

```json
{
  "users": [
    {
      "pin": "cloudid001",
      "companyName": "东东有限公司"
    }
  ],
  "pageNumber": 1,
  "pageSize": 10,
  "totalCount": 2
}
```

### 查询用户详情

##### url

```http
v1/offline/user
```

##### method

```http
get
```

##### 请求参数

参数名|类型|是否必须|默认值|描述
---|---|---|---|---
username|String|否| |用户名


##### 响应参数

```json
{
  "user": {
    "pin": "cloudid001",
    "companyName": "",
    "phoneNumber": "13811622092",
    "email": "hexinlei@jd.com",
    "companyAuthStatus": false
  }
}
```





### 查询当前登录用户

##### url

```http
v1/offline/userLogin
```

##### method

```http
get
```

##### 请求参数

_无_

##### 响应参数

```json
{
  "userName": "xxx"
}
```

### 查询实例列表-已废弃

##### url

```http
v1/offline/instanceList
```

##### method

```http
get
```

##### 请求参数

参数名|类型|是否必须|默认值|描述
---|---|---|---|---
pageNumber|int|否| 1|第 几 页
pageSize|int|否| 20|每页大小


##### 响应参数

```json
{
  "instances": [
    {
      "pin": "xxx",
      "instanceId": "cps-vibbcnxisj768nccslap8zqjj3lf",
      "region": "华北-北京",
      "az": "可用区B",
      "deviceType": "标准计算型",
      "regionEn": "cn-north-1",
      "azEn": "cn-north-1b",
      "deviceTypeEn": "cps.c.normal",
      "name": "bm_test",
      "description": "weqe",
      "bandwidth": 1, //带宽
      "extraUplinkBandwidth": 20, //额外上行带宽
      "lineType": "bgp", //带宽类型（中文）
      "lineTypeEn": "bgp", //带宽类型（英文）
      "sn": "1B3MXP2", //设备号
      "privateIp": "10.0.0.2", //内网IP
      "publicIp": "10.0.0.2", //公网ip
    }
  ],
  "pageNumber": 1,
  "pageSize": 10,
  "totalCount": 2
}
```

### 导出实例列表-已废弃

##### url

```http
v1/offline/exportInstances
```

##### method

```http
get
```

##### 请求参数

_无_

##### 响应参数

_无_

### 查询设备实例列表

##### url

```http
v1/offline/deviceInstanceList
```

##### method

```http
get
```

##### 请求参数

参数名|类型|是否必须|默认值|描述
---|---|---|---|---
instanceId|String|否| |实例 id
sn|String|否| |实例 sn(支持多选，逗号分隔）
saleStatus|String|否| |售卖状态
status|String|否| |实例状态
iloIp|String|否| |带外ip(支持多选，逗号分隔）
publicIp|String|否| |公网ip(支持多选，逗号分隔）
privateIp|String|否| |内网ip(支持多选，逗号分隔）
startTime|int64|否| |开始时间
endTime|int64|否| |开始时间
pageNumber|int|否| 1|第 几 页
pageSize|int|否| 20|每页大小

##### 响应参数

```json
{
  "deviceInstances": [
    {
      "pin": "xxx",
      "instanceId": "cps-vibbcnxisj768nccslap8zqjj3lf",
      "region": "华北-北京",
      "az": "可用区B",
      "deviceType": "标准计算型",
      "regionEn": "cn-north-1",
      "azEn": "cn-north-1b",
      "deviceTypeEn": "cps.c.normal",
      "name": "bm_test",
      "description": "weqe",
      "status": "销毁中", //实例状态
      "statusEn": "destroying",
      "sn": "1B3MXP2",
      "iloIp": "10.208.11.131",
      "cabinet": "23234343",
      "uPosition": "24-28",
      "saleStatus": "已售卖",
      "saleStatusEn": "sold",
      "osName": "CentOS 6.6 64位",
      "osType": "CentOS",
      "privateIp": "10.0.0.2",
      "publicIp": "10.0.0.2",
      "bandwidth": 10, //带宽
      "extraUplinkBandwidth": 10, //额外上行带宽
      "switchIpOne": "10.208.11.35", //交换机1 ip
      "switchIpTwo": "10.208.11.35", //交换机2 ip
      "eth0SwitchPort": "Ten-GigabitEthernet1/0/2", // eth0
      "eth1SwitchPort": "Ten-GigabitEthernet1/0/26", // eth1
      "secondaryIp": ["xxx", "xxxx"], //别名ip
    }
  ]
}
```

### 查询设备实例详情

##### url

```http
v1/offline/deviceInstance
```

##### method

```http
get
```

##### 请求参数

参数名|类型|是否必须|默认值|描述
---|---|---|---|---
sn|String|否| |实例 sn



##### 响应参数

```json
{
  "deviceInstance": {
    "pin": "xxx",
    "instanceId": "cps-vibbcnxisj768nccslap8zqjj3lf",
    "region": "华北-北京",
    "az": "可用区B",
    "deviceType": "标准计算型",
    "regionEn": "cn-north-1",
    "azEn": "cn-north-1b",
    "deviceTypeEn": "cps.c.normal",
    "name": "bm_test",
    "description": "weqe",
    "status": "销毁中", //实例状态
    "statusEn": "destroying",
    "sn": "1B3MXP2",
    "iloIp": "10.208.11.131",
    "cabinet": "23234343",
    "uPosition": "24-28",
    "saleStatus": "已售卖",
    "saleStatusEn": "sold",
    "osName": "CentOS 6.6 64位",
    "osType": "CentOS",
    "privateIp": "10.0.0.2",
    "publicIp": "10.0.0.2",
    "bandwidth": 10, //带宽
    "extraUplinkBandwidth": 10, //额外上行带宽
    "switchIpOne": "10.208.11.35", //交换机1 ip
    "switchIpTwo": "10.208.11.35", //交换机2 ip
    "eth0SwitchPort": "Ten-GigabitEthernet1/0/2", // eth0
    "eth1SwitchPort": "Ten-GigabitEthernet1/0/26", // eth1
    "secondaryIp": ["xxx", "xxxx"], //别名ip
  }
}
```

### 导出设备实例列表

##### url

```http
v1/offline/exportDeviceInstanceList
```

##### method

```http
get
```

##### 请求参数

参数名|类型|是否必须|默认值|描述
---|---|---|---|---
instanceId|String|否| |实例 id
sn|String|否| |实例 sn(支持多选，逗号分隔）
saleStatus|String|否| |售卖状态
status|String|否| |实例状态
iloIp|String|否| |带外ip(支持多选，逗号分隔）
publicIp|String|否| |公网ip(支持多选，逗号分隔）
privateIp|String|否| |内网ip(支持多选，逗号分隔）
startTime|int64|否| |开始时间
endTime|int64|否| |开始时间


##### 响应参数

_无_

### 查询设备库存

##### url

```http
v1/offline/devicesStock
```

##### method

```http
get
```

##### 请求参数


参数名|类型|是否必须|默认值|描述
---|---|---|---|---
deviceType|String|否| |设备类型
pageNumber|int|否| 1|第 几 页
pageSize|int|否| 20|每页大小


##### 响应参数

```json
{
  "devicesStock": [
    {
      "region": "华北-北京",
      "az": "可用区B",
      "deviceType": "标准计算型",
      "regionEn": "cn-north-1",
      "azEn": "cn-north-1b",
      "deviceTypeEn": "cps.c.normal",
      "total": 100, //总数
      "alreadyOnline": 10, //开放售卖数量
      "alreadySold": 20, //已售卖数量
      "unSold": 20, //剩余数量
      "unOnline": 20 //未开放售卖数量
    }
  ],
  "pageNumber": 1,
  "pageSize": 10,
  "totalCount": 2
}
```

### 导出设备库存

##### url

```http
v1/offline/exportDevicesStock
```

##### method

```http
get
```

##### 请求参数

参数名|类型|是否必须|默认值|描述
---|---|---|---|---
deviceType|String|否| |设备类型

##### 响应参数

_无_

### 设备上架

##### url

```http
v1/offline/putawayDevices
```

##### method

```http
put
```

##### 请求 body

参数名|类型|是否必须|默认值|描述
---|---|---|---|---
sn|[]String|是| |设备sn

##### 响应参数

```json
{
  "success": true
}
```

### 设备下架

##### url

```http
v1/offline/unPutawayDevices
```

##### method

```http
put
```

##### 请求 body

参数名|类型|是否必须|默认值|描述
---|---|---|---|---
sn|[]String|是| |设备sn

##### 响应参数

```json
{
  "success": true
}
```

### 设备回收

##### url

```http
v1/offline/instance
```

##### method

```http
delete
```

##### 请求参数

参数名|类型|是否必须|默认值|描述
---|---|---|---|---
instanceId|String|是| |实例id



##### 响应参数

```json
{
  "success": true
}
```

### 售卖状态查询

##### url

```http
v1/offline/saleStatus
```

##### method

```http
get
```

##### 请求参数

_无_

##### 响应参数

```json
{
  "saleStatus": [
    {
      "status": "sold",
      "statusName": "已售卖"
    },
    {
      "status": "selling",
      "statusName": "售卖中"
    },
    {
      "status": "putawaying",
      "statusName": "上架"
    }
  ]
}
```

### 实例状态查询

##### url

```http
v1/offline/instanceStatus
```

##### method

```http
get
```

##### 请求参数

_无_

##### 响应参数

```json
{
  "instanceStatus": [
    {
      "status": "creating",
      "statusName": "创建中"
    },
    {
      "status": "starting",
      "statusName": "开机中"
    },
    {
      "status": "running",
      "statusName": "运行"
    },
    {
      "status": "stopping",
      "statusName": "关机中"
    },
    {
      "status": "stopped",
      "statusName": "关机"
    },
    {
      "status": "restarting",
      "statusName": "重启中"
    },
    {
      "status": "reinstalling",
      "statusName": "重装系统中"
    },
    {
      "status": "destroying",
      "statusName": "销毁中"
    },
    {
      "status": "error",
      "statusName": "错误"
    },
    {
      "status": "upgrading",
      "statusName": "调整配置中"
    },
    {
      "status": "resetting_password",
      "statusName": "重置密码中"
    }
  ]
}
```

### 机型管理列表

##### url

```http
v1/offline/deviceTypes/plus
```

##### method

```http
GET
```

##### 请求参数


参数名|类型|是否必须|默认值|描述
---|---|---|---|---
deviceType|String|否|cps.s.normal |实例类型，非必填
pageNumber|int|否| 1|第 几 页
pageSize|int|否| 20|每页大小


##### 请求响应

```json
{
  "deviceTypes": [
    {
      "deviceType": "cps.gpu.4", //规格
      "name": "GPU IV", //实例类型
      "region": "cn-east-2", //地域
      "family": "gpu", //系列
      "cpuConcise": "12核", //cpu概述
      "cpuDetail": "2*Intel 2683V4（12核 2.2G）", //cpu详情
      "memConcise": "256G", //内存概述
      "memDetail": "256G（8*32G）", //内存详情
      "ifConcise": "2*10GE", //网卡概述
      "ifDetail": "2*10GE", //网卡详情
      "gpuConcise": "NVIDIA P40 * 4", //gpu概述
      "gpuDetail": "NVIDIA P40 * 4", //gpu详情
      "systemDiskAmount": 2, //系统盘数量
      "systemDiskSize": 300, //系统盘大小
      "systemDiskModel": "SAS", //系统盘类型
      "dataDiskAmount": 1, //数据盘数量
      "dataDiskSize": 4000, //数据盘大小
      "dataDiskModel": "SSD NVME", //数据盘类型
      "source": "common", //来源：通用common,定制customize,自定义user_defined
      "osCount": 3, //关联镜像数
      "dataRaids": ["RAID0", "RAID1"],
      "sysRaids": ["RAID0", "RAID1"],
      "serviceCode": "cps" //所属产品，cps/edcps
    }
  ],
  "pageNumber": 1,
  "pageSize": 20,
  "totalCount": 12
}
```
