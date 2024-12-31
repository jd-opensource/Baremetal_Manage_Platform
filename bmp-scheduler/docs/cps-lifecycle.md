
* [机型录入](#机型录入)
	* [机型录入相关数据表](#机型录入对应数据表)
	* [机型录入api](#机型录入对应api)
	* [机型录入界面入口](#机型录入对应界面入口)
* [设备录入](#设备录入)
	* [设备录入相关数据表](#设备录入对应数据表)
	* [设备录入api](#设备录入对应api)
	* [设备录入界面入口](#设备录入对应界面入口)
* [厂商和硬盘信息录入](#厂商和硬盘信息录入)
	* [厂商和硬盘信息录入相关数据表](#厂商和硬盘信息录入对应数据表)
	* [厂商和硬盘信息录入api](#厂商和硬盘信息录入对应api)
	* [厂商和硬盘信息录入界面入口](#厂商和硬盘信息录入对应界面入口)
* [镜像录入](#镜像录入)
	* [镜像录入相关数据表](#镜像录入对应数据表)
	* [镜像录入api](#镜像录入对应api)
	* [镜像录入界面入口](#镜像录入对应界面入口)
* [机型和镜像关联](#机型和镜像关联)
	* [机型和镜像关联相关数据表](#机型和镜像关联相关表)
	* [机型和镜像关联api](#机型和镜像关联api)
	* [机型和镜像关联界面入口](#机型和镜像关联界面入口)
* [raid录入](#raid录入)
	* [raid录入相关数据表](#raid录入相关数据表)
	* [raid录入api](#raid录入相关api)
	* [raid录入界面入口](#raid录入相关界面入口)
* [机型和raid关联](#机型和raid关联)
	* [机型和raid关联相关数据表](#机型和raid关联相关数据表)
	* [机型和raid关联api](#机型和raid关联相关api)
	* [机型和raid关联界面入口](#机型和raid关联相关界面入口)
* [设备上架售卖](#设备上架售卖)
	* [设备上架售卖api](#设备上架售卖相关api)
	* [设备上架售卖界面入口](#设备上架售卖相关界面入口)
* [设备下架](#设备下架)
	* [设备下架api](#设备下架相关api)
	* [设备下架界面入口](#设备下架相关界面入口)
* [设备退库回收](#设备退库回收)
	* [设备退库回收api](#设备退库回收相关api)
	* [设备退库回收界面入口](#设备退库回收相关界面入口)
* [创建实例](#创建实例)
	* [创建实例api](#创建实例相关api)
	* [创建实例界面入口](#创建实例相关界面入口)
* [实例开机](#实例开机)
	* [实例开机api](#实例开机相关api)
	* [实例开机界面入口](#实例开机相关界面入口)
* [实例关机](#实例关机)
	* [实例关机api](#实例关机相关api)
	* [实例关机界面入口](#实例关机相关界面入口)
* [实例重启](#实例重启)
	* [实例重启api](#实例重启相关api)
	* [实例重启界面入口](#实例重启相关界面入口)
* [实例重装系统](#实例重装系统)
	* [实例重装系统api](#实例重装系统相关api)
	* [实例重装系统界面入口](#实例重装系统相关界面入口)
* [实例重置密码](#实例重置密码)
	* [实例重置密码api](#实例重置密码相关api)
	* [实例重置密码界面入口](#实例重置密码相关界面入口)
* [实例销毁](#实例销毁)
	* [实例销毁api](#实例销毁相关api)
	* [实例销毁界面入口](#实例销毁相关界面入口)





## 机型录入


新采购一批裸机后，首先要做的就是机型(即实例类型)的录入

#### 机型录入对应数据表

机型包含的主要信息在cps_ironic.device_type表中：



| id  | az          | device_type   | name_en                                    | name_zh                        | cpu_amount | cpu_cores | cpu_manufacturer | cpu_model | cpu_frequency | mem_amount | mem_size | nic_amount | nic_rate | system_volume_amount | system_volume_size | data_volume_amount | data_volume_size | gpu_amount | gpu_manufacturer | gpu_model | create_time         | update_time         | is_del |
|-----|-------------|---------------|--------------------------------------------|--------------------------------|------------|-----------|------------------|-----------|---------------|------------|----------|------------|----------|----------------------|--------------------|--------------------|------------------|------------|------------------|-----------|---------------------|---------------------|--------|
| 101 | cn-north-1a | cps.c.perf1   | Compute Performance Ⅰ                      | 计算效能型Ⅰ                    |          2 |        16 | Intel            | 2683V4    | 2.1           |          8 |       32 |          2 |       10 |                    2 |                300 |                  1 |             4000 |       NULL | NULL             | NULL      | 2018-12-03 13:53:25 | 2018-12-03 13:53:25 |      0 |
| 102 | cn-north-1c | cps.c2.perf1  | Compute PerformanceⅠ（2nd generation）     | 计算效能型Ⅰ（二代）            |          2 |        20 | Intel            | GOLD6148  | 2.4           |         12 |       32 |          2 |       10 |                    1 |                240 |                  1 |             2000 |       NULL | NULL             | NULL      | 2018-12-03 13:53:25 | 2018-12-03 13:53:25 |      0 |
| 103 | cn-north-1a | cps.c.perf2   | Compute Performance Ⅱ                      | 计算效能型Ⅱ                    |          2 |        16 | Intel            | 2683V4    | 2.1           |          8 |       32 |          2 |       10 |                    1 |                240 |                 16 |              960 |       NULL | NULL             | NULL      | 2018-12-03 13:53:25 | 2018-12-03 13:53:25 |      0 |
| 104 | cn-north-1c | cps.s2.normal | Storage Standard（2nd generation）         | 标准存储型（二代）             |          2 |        12 | Intel            | 2683V4    | 2.1           |          8 |       32 |          2 |       10 |                    2 |                300 |                 12 |           120000 |       NULL | NULL             | NULL      | 2018-12-03 13:53:25 | 2018-12-03 13:53:25 |      0 |




#### 机型录入对应api

私有化本次封装api [post /v1/offline/CreateBmClazz](https://git.jd.com/minping/ironic-console/blob/master/docs/offline-api.md#%E6%9C%BA%E5%9E%8B%E5%BD%95%E5%85%A5)
#### 机型录入对应界面入口
cps这个录入操作没有页面，pm提供excel，rd直接插表，qa校验。


## 设备录入


每次采购了新机器，pm都会发类似如下的excel表格：








| SN   | 系统IP | 管理IP | 品牌 | 型号 | 机房 | 机柜 | U位 | 套餐 |
| ---  | ----- | ----- | ---- | --- | --- | --- | --- | --- |
| 210235A2CTH219000001 | 10.209.75.139 | 10.208.75.139 | H3C | R4900 G3 | 北京_中国联通_有孚 | B5-L5-1-I-13 | 30--31 | D2001 |
| 210235A2CTH219000002 | 10.209.75.138 | 10.208.75.138 | H3C | R4900 G3 | 北京_中国联通_有孚 | B5-L5-1-I-13 | 27--28 | D2001 |
| 210235A2CTH219000003 | 10.209.75.137 | 10.208.75.137 | H3C | R4900 G3 | 北京_中国联通_有孚 | B5-L5-1-I-13 | 24--25 | D2001 |








然后第一部要做的就是录入设备，可通过ironic-api模块的/devices/CreateDevices接口来进行设备入库(cps也可以通过idc信息同步来入库)，其实就是插入到device表中


#### 设备录入对应数据表


就是device表，内容如下：




| id | sn       | region      | az           | device_type          | manufacturer_id | ilo_ip        | cabinet | u_position | sale_status | create_time         | update_time         | is_del | remark |
|----|----------|-------------|--------------|----------------------|-----------------|---------------|---------|------------|-------------|---------------------|---------------------|--------|--------|
| 20 | J300GT9H | cn-east-tz1 | cn-east-tz1a | edcps.c2.perf1       |              15 | 10.208.13.79  |         |            | sold        | 2019-02-20 16:31:32 | 2020-09-17 15:17:30 |      0 | NULL   |
| 54 | J300GZ8P | cn-north-1  | cn-north-1c  | cps.s2.normal        |              21 | 10.208.13.107 | NULL    | NULL       | sold        | 2019-07-18 10:32:01 | 2021-06-24 10:38:40 |      0 | NULL   |
| 63 | J301B9B6 | cn-north-1  | cn-north-1c  | cps.c2.perf2         |              26 | 10.208.13.197 | NULL    | NULL       | sold        | 2020-01-02 18:09:11 | 2021-09-30 11:36:46 |      0 | NULL   |
| 65 | 2CVZ1K2  | cn-east-tz1 | cn-east-tz1a | edcps.huya.c.normal1 |              27 | 10.208.16.41  | NULL    | NULL       | sold        | 2020-02-04 18:23:34 | 2020-08-18 19:55:25 |      0 | NULL   |
| 66 | 2D642K2  | cn-north-1  | cn-north-1a  | cps.c.normal         |              32 | 10.208.11.98  | NULL    | NULL       | sold        | 2020-02-04 18:23:34 | 2021-07-13 14:23:16 |      0 | NULL   |




注意录入设备这一步只插入device_type， ilo_ip等信息，manufacturer_id是后面硬件设备信息采集才插入进去的。device_type就是对应的上面需求表中的套餐，两者的对应关系成琳在cps库中新加了一张physical_machine_package表来记录对应关系。


设备录入完成后，sale_status为上架中putawaying


#### 设备录入对应api


私有化本次封装api [post v1/offline/createDevices](https://git.jd.com/minping/ironic-console/blob/master/docs/offline-api.md#%E8%AE%BE%E5%A4%87%E5%BD%95%E5%85%A5)




#### 设备录入对应界面入口


无


## 厂商和硬盘信息录入

上面的机型device_type表包含cpu，cpu_core，mem，nic，gpu等各种信息，但是唯独没有磁盘和厂商相关信息。其实这部分信息是放在manufacturer表和disk表中的。

#### 厂商和硬盘信息录入对应数据表


manufacturer表内容：




| id | device_type  | manufacturer                        | product_name                                                  | create_time         | update_time         | is_del |
|----|--------------|-------------------------------------|---------------------------------------------------------------|---------------------|---------------------|--------|
|  1 | cps.c.normal | Huawei                              | RH1288 V3 (Type1Sku0)                                         | 2018-12-26 13:28:10 | 2018-12-26 13:28:13 |      0 |
|  2 | cps.c.perf2  | Dell Inc.                           | PowerEdge R730xd (SKU=NotProvided;ModelName=PowerEdge R730XD) | 2018-12-26 13:28:16 | 2018-12-26 13:28:18 |      0 |
|  3 | cps.c.perf1  | Dell Inc.                           | PowerEdge R630 (SKU=NotProvided;ModelName=PowerEdge R630)     | 2018-12-26 13:28:20 | 2018-12-26 13:28:22 |      0 |
|  4 | cps.s.normal | Unis Huashan Technologies Co., Ltd. | R4900 G2 (0)                                                  | 2018-12-26 13:28:24 | 2018-12-26 13:28:27 |      0 |
|  5 | cps.gpu.1    | Sugon                               | W580-G20 (Default string)                                     | 2018-12-26 13:28:29 | 2018-12-26 13:28:31 |      0 |
| 26 | cps.c2.perf2 | Lenovo       | ThinkSystem SR650 -[7X06CTO1WW]- (7X06CTO1WW) | 2020-01-02 18:14:17 | 2020-01-02 18:14:17 |      0 |
| 42 | cps.c2.perf2 | Inspur       | SA5212M5 (Default string)                     | 2020-12-01 11:36:44 | 2020-12-01 11:36:44 |      0 |






disk表内容：




| id | manufacturer_id | enclosure | slot | disk_type | create_time         | update_time         | is_del |
|----|-----------------|-----------|------|-----------|---------------------|---------------------|--------|
| 18 |               1 | 252       |    0 | system    | 2018-12-04 10:57:29 | 2018-12-04 10:57:29 |      0 |
| 19 |               1 | 252       |    1 | system    | 2018-12-04 10:57:29 | 2018-12-04 10:57:29 |      0 |
| 20 |               1 | 252       |    2 | data      | 2018-12-04 10:57:29 | 2018-12-04 10:57:29 |      0 |
| 21 |               1 | 252       |    3 | data      | 2018-12-04 10:57:29 | 2018-12-04 10:57:29 |      0 |
| 22 |               3 | 32        |    0 | system    | 2018-12-04 10:57:29 | 2018-12-04 10:57:29 |      0 |




用户选择一种机型一种镜像创建实例时，如果这种机型对应多种manufacturer(比如上面的cps.c2.perf2机型有Lenovo和Inspur两个厂商可以提供)，那么现在的逻辑是随机选一个。其实manufacturer对用户是透明的，而且这里的manufacturer信息只是给disk提供服务的，所以个人觉得manufacturer表完全可以去掉，而且manufacturer表其实是一个机型-manufacturer关联表，应该叫r_device_type_disk这种关联表名字更合适。但是agent和监控好像依赖manufacturer信息。


#### 厂商和硬盘信息录入对应api


这里分两步，先接口采集厂商等基本信息，后手动登录到机器获取磁盘槽位信息，手动录入信息到disk表


- 接口采集厂商等基本信息这一步是手动调用ironic-api模块的/devices/collectDeviceInfo接口，到ironic-scheduler模块下发collecthardwareinfo命令到agent，agent采集了上传到scheduler模块，插入manufacturer表中。
- 然后通过device表该设备的带外ip登录进去，用/usr/sbin/MegaCli64 -PDList -aAll | egrep 'Enclosure | Slot'等命令获取槽位信息，将槽位信息手动插入到disk表中。


私有化本次封装api [post v1/offline/collectManufacturerAndDisk](https://git.jd.com/minping/ironic-console/blob/master/docs/offline-api.md#%E5%8E%82%E5%95%86%E5%92%8C%E7%A1%AC%E7%9B%98%E4%BF%A1%E6%81%AF%E5%BD%95%E5%85%A5)

#### 厂商和硬盘信息录入对应界面入口


采集设备厂商基本信息和采集槽位信息均无界面入口，纯手工。



## 镜像录入
上面的机型是硬件相关的，这里的镜像就是跟操作系统相关了，用户在创建实例的时候，需要选择使用哪个镜像系统。所以需要管理员先在运营平台上先录入可用的镜像。cps方面，都是唐湘先制作好镜像，上传到nginx对应镜像下载目录上，然后通知rd录入到image表中。


#### 镜像录入对应数据表


镜像这里有两个相关表，os表和image表，os表内容如下：




| id | uuid                           | name_en                                                          | name_zh                                          | os_type | platform | architecture | bits | version | sys_user | create_time         | update_time         | is_del |
|----|--------------------------------|------------------------------------------------------------------|--------------------------------------------------|---------|----------|--------------|------|---------|----------|---------------------|---------------------|--------|
|  1 | o-1vc0l7zpnrt1rpnssudanmq01jbv | CentOS 6.6 64-bit                                                | CentOS 6.6 64位                                  | CentOS  | linux    | x86          |   64 | 6.6     | root     | 2018-12-17 17:40:56 | 2018-12-17 17:40:56 |      0 |
|  2 | o-hea0bysnymuquyhyrfgtxj6qtc8e | CentOS 7.1 64-bit                                                | CentOS 7.1 64位                                  | CentOS  | linux    | x86          |   64 | 7.1     | root     | 2018-12-17 17:40:56 | 2018-12-17 17:40:56 |      0 |
|  3 | o-fspugrsoajxqzt1tncelmodsbsuc | CentOS 7.2 64-bit                                                | CentOS 7.2 64位                                  | CentOS  | linux    | x86          |   64 | 7.2     | root     | 2018-12-17 17:40:56 | 2018-12-17 17:40:56 |      0 |
|  4 | o-zqxfcdf99c6fgkihaaayd1uiaofk | CentOS 7.5 64-bit                                                | CentOS 7.5. 64位                                 | CentOS  | linux    | x86          |   64 | 7.5     | root     | 2018-12-17 17:40:56 | 2018-12-17 17:40:56 |      0 |
|  5 | o-1c9ypijcvtsbqonhn2ukw5ljxgpq | Ubuntu 14.04 64-bit                                              | Ubuntu 14.04 64位                                | Ubuntu  | linux    | x86          |   64 | 14.04   | root     | 2018-12-17 17:40:56 | 2018-12-17 17:40:56 |      0 |
|  6 | o-1uz8qa06mtz2lve8ibtabkh5kwna | Ubuntu 16.04 64-bit                                              | Ubuntu 16.04 64位                                | Ubuntu  | linux    | x86          |   64 | 16.04   | root     | 2018-12-17 17:40:56 | 2018-12-17 17:40:56 |      0 |
|  7 | o-s916eda0byt6mihuxp77axmhybpn | Ubuntu 18.04 64-bit                                              | Ubuntu 18.04 64位                                | Ubuntu  | linux    | x86          |   64 | 18.04   | root     | 2018-12-17 17:40:56 | 2018-12-17 17:40:56 |      0 |
|  9 | o-t4gru7kqadqaftv1mgubvum6904a | CentOS 7.5 64-bit TD Plus                                        | CentOS 7.5 64位 TD增强版                         | CentOS  | linux    | x86          |   64 | 7.5     | root     | 2019-05-24 17:04:38 | 2019-05-24 17:04:38 |      0 |
| 10 | o-82oueik4ihvosopcoc1m4ipifv3l | Ubuntu 18.04 64-bit ARM                                          | Ubuntu 18.04 64位 ARM                            | Ubuntu  | linux    | x86          |   64 | 18.04   | root     | 2019-11-12 11:36:37 | 2019-11-12 11:36:37 |      0 |
| 11 | o-pwrdu2andavqyp1yqxlgaqrg4ywn | CentOS 7.2 64-bit                                                | CentOS 7.2 64-bit定制版                          | CentOS  | linux    | x86          |   64 | 7.2     | root     | 2019-12-04 18:07:26 | 2019-12-04 18:07:26 |      0 |
| 14 | o-fb398c1dc66d427fa0da35a898e5 | CentOS 7.6-bit                                                   | CentOS 7.6 64位                                  | CentOS  | linux    | x86          |   64 | 7.6     | root     | 2019-12-23 17:54:39 | 2019-12-23 17:54:39 |      0 |
| 15 | o-16a13d56847f4b0e9755d9c0cdb1 | Ubuntu 16.04 Huya 64-bit                                         | Ubuntu 16.04 Huya 64位                           | Ubuntu  | linux    | x86          |   64 | 16.04   | root     | 2020-02-19 11:36:37 | 2020-02-19 11:36:37 |      0 |
| 16 | o-17ea6e54932c455c9e51909b568e | Windows Server 2016 Standard Version, 64-bit, Chinese Version    | Windows Server 2016 标准版 64位 中文版           | Windows | linux    | x86          |   64 | 2016    | root     | 2020-06-11 10:20:20 | 2020-08-12 18:50:17 |      0 |
| 17 | o-caecaf97b62a4cb7baab799dfd3c | Windows Server 2012 R2 Standard Version, 64-bit, Chinese Version | Windows Server 2012 R2 标准版 64位 中文版        | Windows | linux    | x86          |   64 | 2012    | root     | 2020-07-08 18:20:41 | 2020-08-12 18:53:05 |      0 |
| 18 | o-a628710f69d2486e9842f822ff56 | Windows 2012 normal                                              | Windows 2012 normal                              | Windows | linux    | x86          |   64 | 2012    | root     | 2020-07-14 11:50:29 | 2020-07-14 11:50:29 |      0 |
| 19 | o-50b9621c383b4d47b8f19e385630 | CentOS 7.4 64-bit BDP customize                                  | CentOS 7.4 64位 BDP定制版                        | CentOS  | linux    | x86          |   64 | 7.4     | root     | 2020-09-01 20:57:41 | 2020-09-01 20:57:41 |      0 |
| 20 | o-b4969d4f9b464b67a46ebb6f1368 | CentOS 8.2 64-bit                                                | CentOS 8.2 64位                                  | CentOS  | linux    | x86          |   64 | 8.2     | root     | 2020-10-14 20:39:30 | 2020-10-14 20:39:30 |      0 |
| 21 | o-dfaeea19e2c64f17b7c5ace23c3a | CentOS 8.0 64-bit                                                | CentOS 8.0 64位                                  | CentOS  | linux    | x86          |   64 | 8.0     | root     | 2020-10-30 11:08:16 | 2020-10-30 11:08:16 |      0 |
| 22 | o-eda5ef18380f4e9a980777ddb263 | CentOS 7.4 64-bit                                                | CentOS 7.4 64位                                  | CentOS  | linux    | x86          |   64 | 7.4     | root     | 2020-11-11 10:36:50 | 2020-11-11 10:36:50 |      0 |
| 23 | o-s916eda0byt6mihuxp77axmhubun | Ubuntu 20.04 64-bit                                              | Ubuntu 20.04 64位                                | Ubuntu  | linux    | x86          |   64 | 20.04   | root     | 2018-12-17 17:40:56 | 2018-12-17 17:40:56 |      0 |




image表内容如下：






| id   | uuid                               | name_en                                                          | name_zh                                                         | version | os_id                          | format | filename                                               | url                                                           | hash                             | source    | create_time         | update_time         | is_del |
|------|------------------------------------|------------------------------------------------------------------|-----------------------------------------------------------------|---------|--------------------------------|--------|--------------------------------------------------------|---------------------------------------------------------------|----------------------------------|-----------|---------------------|---------------------|--------|
|   21 | i-d656757978ec4ad78a7d4f515037     | Windows Server 2012 R2 Standard Version, 64-bit, Chinese Version | Windows Server 2012 R2 标准版 64位 中文版                       | 2012    | o-caecaf97b62a4cb7baab799dfd3c | qcow2  | windows-server-2012-cn.qcow2                           | {host}/windows-server-2012-cn.qcow2                           | 7eaeb8d9a8140cda22dd2174b2236ab4 | common    | 2020-07-08 18:20:41 | 2020-08-12 19:01:50 |      0 |
| 1006 | i-usk5qjugotg56r0jlaruy57nzswwwtar | Ubuntu 20.04 64-bit                                              | Ubuntu 20.04 64位 UEFI TAR                                      | 20.04   | o-s916eda0byt6mihuxp77axmhubun | tar    | ubuntu-20.04-uefi.tar.xz                               | {host}/ubuntu-20.04-uefi.tar.xz                               | 433a199dd041ac6a048fd6f3df2fa88b | common    | 2018-11-27 16:41:51 | 2018-11-27 16:41:51 |      0 |
| 1118 | i-rodl94xo5nevgovh5y8k8rnofbhv     | CentOS 7.5 64-bit TD Plus                                        | CentOS 7.5 64位 TD 增强版                                       | 0       | o-t4gru7kqadqaftv1mgubvum6904a | qcow2  | v1.3.0-centos-7.5-td-plus-2019081316.qcow2             | {host}/v1.3.0-centos-7.5-td-plus-2019081316.qcow2             | 58e880243671a60f4fcc8805a333844a | customize | 2019-05-30 21:00:21 | 2019-08-26 21:01:56 |      0 |
| 1135 | i-tvuhgdqltetmuhem9fh7wq3dtmuy     | CentOS 8.2 64-bit for mysql and es                               | CentOS 8.2 64位 零售MySQL&ES定制版                              | 8.2     | o-b4969d4f9b464b67a46ebb6f1368 | tar    | v1.3.0-centos-8.2-mysql-es-2020120416.tar.xz           | {host}/v1.3.0-centos-8.2-mysql-es-2020120416.tar.xz           | da981058f8f3fbe59dcf77ed21b587f7 | customize | 2020-10-29 16:44:00 | 2020-10-29 16:44:00 |      0 |


可以看到，os表是基础表，image表信息包含全部的os表信息，业务上几乎没有用到os表，用的都是image表。
image表包含如下信息：


- uuid
- 中、英文镜像名称
- version版本
- 关联os_id
- 创建/修改时间
- 镜像下载url
- 镜像文件hash
- 来源 common/customize




#### 镜像录入对应api


私有化本次封装api [post v1/offline/createImage](https://git.jd.com/minping/ironic-console/blob/master/docs/offline-api.md#%E9%95%9C%E5%83%8F%E5%BD%95%E5%85%A5)


#### 镜像录入对应界面入口


无



## 机型和镜像关联

有些机型只能使用特定image镜像来装系统，这种关系称为机型和镜像关联。在页面上创建实例的时候也是先选机型，根据选择的机型给出可用的镜像供选择，等机型和镜像选好了后才可以创建实例装机。


#### 机型和镜像关联相关表


镜像和机型关联表为 r_device_type_image表和r_device_type_image_partition表。
在创建实例，选择了机型后，列出可用的镜像就是查r_device_type_image表，表内容为：




| id | image_id                       | device_type  | create_time         | update_time         | is_del |
|----|--------------------------------|--------------|---------------------|---------------------|--------|
| 37 | i-vtgpscwagbvd7urak6jk4ykstmff | cps.s.normal | 2019-01-04 20:07:51 | 2019-01-04 20:07:51 |      0 |
| 38 | i-yhzcyaxnw6rtieorpgfdnbfshfh9 | cps.s.normal | 2019-01-04 20:07:51 | 2019-01-04 20:07:51 |      0 |
| 44 | i-vtgpscwagbvd7urak6jk4ykstmff | cps.c.normal | 2019-01-04 20:07:51 | 2019-01-04 20:07:51 |      0 |
| 45 | i-yhzcyaxnw6rtieorpgfdnbfshfh9 | cps.c.normal | 2019-01-04 20:07:51 | 2019-01-04 20:07:51 |      0 |
| 51 | i-vtgpscwagbvd7urak6jk4ykstmff | cps.c.perf1  | 2019-01-04 20:07:51 | 2020-06-23 21:05:18 |      0 |
| 52 | i-yhzcyaxnw6rtieorpgfdnbfshfh9 | cps.c.perf1  | 2019-01-04 20:07:51 | 2019-01-04 20:07:51 |      0 |





r_device_type_image_partition表看名字其实和镜像关联相关，但其实跟镜像关联的意义不大，这个表其实是给装机过程中的做分区这个子command来使用的。只要是tar格式的镜像，一般在r_device_type_image_partition表中都有一行system(系统盘)的分区信息记录(如果没有则装机在make partition这一步会失败，导致实例创建失败)，data(数据盘)分区信息记录可有可没有。r_device_type_image_partition表内容如下：




| id | image_id                       | device_type               | boot_mode | partition_type | partition_size | partition_fs_type | partition_mount_point | partition_label | system_disk_label | data_disk_label | create_time         | update_time         | is_del |
|----|--------------------------------|---------------------------|-----------|----------------|----------------|-------------------|-----------------------|-----------------|-------------------|-----------------|---------------------|---------------------|--------|
|  1 | i-8f5fc65dd2894c85b9865fddee97 | cps.s.normal.optimized    | bios      | system         |      999999999 | xfs               | /export               | l_export        | gpt               | gpt             | 2020-11-16 18:15:29 | 2020-11-16 18:15:29 |      0 |
|  2 | i-5a1bf8e3f0844bd19894ec21b612 | cps.c2.perf2.netenhanced  | bios      | data           |         102400 | xfs               | /export               | l_export        | gpt               | gpt             | 2020-11-16 18:15:29 | 2020-11-16 18:15:29 |      0 |
|  3 | i-1688c97caa9b414085d91c38b5dd | cps.c2.perf2.netenhanced  | bios      | data           |         102400 | xfs               | /export               | l_export        | gpt               | gpt             | 2020-11-16 18:15:29 | 2020-11-16 18:15:29 |      0 |
|  4 | i-276dfcc0dbc74925bfca67b5a497 | cps.c2.perf2.storenhanced | bios      | system         |      999999999 | xfs               | /data0                | l_data0         | gpt               | gpt             | 2020-11-16 18:15:29 | 2020-11-16 18:15:29 |      0 |
|  5 | i-1688c97caa9b414085d91c38b5dd | cps.c2.perf2.storenhanced | bios      | system         |         102400 | xfs               | /export               | l_export        | gpt               | gpt             | 2020-11-16 18:15:29 | 2020-11-16 18:15:29 |      0 |


#### 机型和镜像关联api

[POST v1/offline/deviceTypes/associatedImage](https://git.jd.com/minping/ironic-console/blob/master/docs/offline-api.md#%E5%85%B3%E8%81%94%E9%95%9C%E5%83%8F)





#### 机型和镜像关联界面入口


运营平台-机型管理-管理镜像-关联镜像/移除





## raid录入

常见的raid方案有NORAID， RAID0， RAID1， RAID10， RAID51H， soft RAID1等，这些信息一般开始就录入进表了，后面基本不会动。


#### raid录入相关数据表


raid表的内容如下：




| id  | uuid                           | name_en | name_zh | description_en                                                                                                                                                                                                                                                          | description_zh                                                                                                                                                                                                                        | create_time         | update_time         | is_del |
|-----|--------------------------------|---------|---------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------|---------------------|--------|
| 100 | r-qytwf9r5h0yn9x4evjkyr0n1cwyb | NORAID  | NORAID  | NO RAID means that each physical disk is displayed independently as a logical disk. There is no block stripping and no data redundancy.                                                                                                                                 | NO RAID是每个物理磁盘独立显示为一个逻辑盘，没有数据块分条（no block stripping），不提供数据冗余。                                                                                                                                     | 2019-01-11 11:55:14 | 2019-01-11 11:55:12 |      0 |
| 101 | r-sfykwibcvcyztuez3mk1klyha9nn | RAID0   | RAID0   | RAID0 means that the entire logical disk data is striped among multiple physical disks, it allows read/write in parallel and provides the fastest speed with no redundancy.                                                                                             | RAID0是整个逻辑盘的数据是被分条（stripped）分布在多个物理磁盘上，可以并行读/写，提供最快的速度，但没有冗余能力。                                                                                                                      | 2018-11-27 10:14:26 | 2018-11-27 10:14:26 |      0 |
| 102 | r-wtzluqacgzzxgunnabdkpnpjew3d | RAID1   | RAID1   | RAID1 also known as mirroring method, provides data redundancy. In the entire mirroring process, only half of the disk capacity is valid, the other half is used to store the same data. RAID1 has taken security into consideration with half capacity and full speed. | RAID1又称镜像方式，提供数据的冗余。在整个镜像过程中，只有一半的磁盘容量是有效的（另一半磁盘容量用来存放同样的数据），RAID1考虑了安全性，容量减半、速度不变。                                                                          | 2018-11-27 10:15:26 | 2018-11-27 10:15:26 |      0 |
| 103 | r-l6pounvfife0njlinhk6ztf2xyio | RAID10  | RAID10  | RAID10 is used for achieving the high speed and security purposes. RAID 10 can be simply interpreted as the RAID 0 array composed of several disks, which then provides the image.                                                                                      | RAID10是为了达到既高速又安全，可以把RAID 10简单地理解成由多个磁盘组成的RAID 0阵列再进行镜像。                                                                                                                                         | 2018-11-27 10:16:08 | 2018-11-27 10:16:08 |      0 |
| 104 | r-lowdd8nf8y7v6jk5ocr5paqhvhre | RAID51H | RAID51H | RAID 5+1 hotspare offers you a resolution featuring high performance, data security and reasonable cost in storage, with added hot backup disks for better data security.                                                                                               | RAID 5+1 hotspare是一种存储性能、数据安全和存储成本兼顾的存储解决方案，同时增加热备盘来提升数据安全。                                                                                                                                 | 2019-05-20 19:50:23 | 2019-05-20 19:50:23 |      0 |
| 105 | r-softraiddpo98m7v6jkferhwpaqh | RAID1   | RAID1   | soft RAID1                                                                                                                                                                                                                                                              | 软RAID1                                                                                                                                                                                                                               | 2021-08-13 17:37:19 | 2021-08-13 17:37:23 |      0 |


#### raid录入相关api

无

#### raid录入相关界面入口


无






## 机型和raid关联


哪些机型(即device_type，也即套餐)的系统盘可以做哪种raid，数据盘可以做哪种raid，是存在r_device_type_raid表中的。其实主要是跟机型的磁盘块数有关系。
	
#### 机型和raid关联相关数据表


r_device_type_raid表的内容如下：






| id  | raid_id                        | device_type  | volume_type | volume_detail       | system_partition_count | create_time         | update_time         | is_del |
|-----|--------------------------------|--------------|-------------|---------------------|------------------------|---------------------|---------------------|--------|
| 101 | r-wtzluqacgzzxgunnabdkpnpjew3d | cps.s.normal | system      | 300GB (2*300GB SAS) |                      1 | 2019-01-14 18:21:07 | 2019-01-14 18:21:07 |      0 |
| 102 | r-qytwf9r5h0yn9x4evjkyr0n1cwyb | cps.s.normal | data        | 72TB (12*6TB SATA)  |                      1 | 2019-01-14 18:21:07 | 2019-01-14 18:21:07 |      0 |
| 103 | r-wtzluqacgzzxgunnabdkpnpjew3d | cps.c.normal | system      | 300GB (2*300GB SAS) |                      1 | 2019-01-14 18:21:07 | 2019-01-14 18:21:07 |      0 |
| 104 | r-qytwf9r5h0yn9x4evjkyr0n1cwyb | cps.c.normal | data        | 1.6TB (2*800GB SSD) |                      1 | 2019-01-14 18:21:07 | 2019-01-14 18:21:07 |      0 |
| 105 | r-sfykwibcvcyztuez3mk1klyha9nn | cps.c.normal | data        | 1.6TB (2*800GB SSD) |                      1 | 2019-01-14 18:21:07 | 2019-01-14 18:21:07 |      0 |





#### 机型和raid关联相关api


私有化本次封装api [post v1/offline/deviceTypes/associatedRaid](https://git.jd.com/minping/ironic-console/blob/master/docs/offline-api.md#%E6%9C%BA%E5%9E%8B%E5%92%8Craid%E5%85%B3%E8%81%94)
	
#### 机型和raid关联相关界面入口

无

## 机型和raid解关联


#### 机型和raid解关联相关api


私有化本次封装api [post v1/offline/deviceTypes/dissociatedRaid](https://git.jd.com/minping/ironic-console/blob/master/docs/offline-api.md#%E6%9C%BA%E5%9E%8B%E5%92%8Craid%E8%A7%A3%E5%85%B3%E8%81%94)


## 设备上架售卖

#### 设备上架售卖相关api


[/v1/offline/putawayDevices](https://git.jd.com/minping/ironic-console/blob/master/docs/offline-api.md#%E8%AE%BE%E5%A4%87%E4%B8%8A%E6%9E%B6)


sale_status前置状态：上架中putawaying
sale_status后置状态：售卖中selling

#### 设备上架售卖相关界面入口


资源管理-开放售卖


## 设备下架

#### 设备下架相关api


[/v1/offline/unPutawayDevices](https://git.jd.com/minping/ironic-console/blob/master/docs/offline-api.md#%E8%AE%BE%E5%A4%87%E4%B8%8B%E6%9E%B6)


sale_status前置状态：售卖中selling
sale_status后置状态：上架中putawaying


#### 设备下架相关界面入口

资源管理-设备下架


## 设备退库回收

#### 设备退库回收相关api

[v1/offline/instance](https://git.jd.com/minping/ironic-console/blob/master/docs/offline-api.md#%E8%AE%BE%E5%A4%87%E5%9B%9E%E6%94%B6)


sale_status前置状态：putawaying
回收完后is_del置为1

#### 设备退库回收相关界面入口
资源管理-设备回收

## 创建实例

#### 创建实例相关api


[POST /api/cps/createInstance](https://git.jd.com/minping/ironic-console/blob/master/docs/console-api.md#%E5%88%9B%E5%BB%BA%E7%89%A9%E7%90%86%E6%9C%BA)

#### 创建实例相关界面入口


## 实例开机
	
#### 实例开机相关api


[PUT /api/cps/startInstance](https://git.jd.com/minping/ironic-console/blob/master/docs/console-api.md#%E5%BC%80%E6%9C%BA)

#### 实例开机相关界面入口



## 实例关机
	
#### 实例关机相关api


[PUT /api/cps/stopInstance](https://git.jd.com/minping/ironic-console/blob/master/docs/console-api.md#%E5%85%B3%E6%9C%BA)

#### 实例关机相关界面入口


## 实例重启

#### 实例重启相关api


[PUT /api/cps/restartInstance](https://git.jd.com/minping/ironic-console/blob/master/docs/console-api.md#%E9%87%8D%E5%90%AF)
#### 实例重启相关界面入口



## 实例重装系统
#### 实例重装系统相关api


[POST /api/cps/reinstallInstance](https://git.jd.com/minping/ironic-console/blob/master/docs/console-api.md#%E9%87%8D%E8%A3%85%E7%89%A9%E7%90%86%E6%9C%BA)
#### 实例重装系统相关界面入口


## 实例重置密码


#### 实例重置密码相关api


[PUT /api/cps/resetPassword](https://git.jd.com/minping/ironic-console/blob/master/docs/console-api.md#%E9%87%8D%E7%BD%AE%E5%AF%86%E7%A0%81)
##### 实例重置密码相关界面入口


## 实例销毁


实例销毁和运营平台的设备回收其实是一样的意思。
#### 实例销毁相关api


[DELETE /api/cps/deleteInstance](https://git.jd.com/minping/ironic-console/blob/master/docs/console-api.md#%E5%88%A0%E9%99%A4%E7%89%A9%E7%90%86%E6%9C%BA)
#### 实例销毁相关界面入口







