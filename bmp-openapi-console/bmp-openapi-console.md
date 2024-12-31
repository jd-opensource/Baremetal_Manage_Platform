
 <h1 class="curproject-name"> bmp-openapi-console </h1> 
 


# deviceType

## /deviceTypes
<a id=/deviceTypes> </a>
### 基本信息

**Path：** /v1/deviceTypes

**Method：** GET

**接口描述：**
DescribeDeviceTypes 获取机型列表

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| pageNumber | 否  |   |  页码 |
| pageSize | 否  |   |  每页数量 |
| deviceTypeId | 否  |   |  机型id |
| deviceType | 否  |   |  机型规格 |
| idcId | 否  |   |  机房id |
| name | 否  |   |  机型名称 |
| deviceSeries | 否  |   |  机型类型 |
| isAll | 否  |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> deviceTypes</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> architecture</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">体系架构，如i386/x86_64/ ARM64 (aarch64)，默认 x86_64</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> boot_mode</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">boot模式</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuCores</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">单个cpu内核数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuFrequency</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu频率(G)</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">拼装信息</span></td><td key=5></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuManufacturer</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu厂商</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuModel</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu处理器型号</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuSpec</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">CPU 规格,预置还是其它</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeInterfaceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘接口类型（SATA,SAS,NVME）</span></td><td key=5></td></tr><tr key=0-0-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘单盘大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-14><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘类型（SSD，HDD）</span></td><td key=5></td></tr><tr key=0-0-0-15><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeUnit</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘单位（GB，TB）</span></td><td key=5></td></tr><tr key=0-0-0-16><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">描述</span></td><td key=5></td></tr><tr key=0-0-0-17><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-18><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceSeries</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型类型，如计算型，存储型</span></td><td key=5></td></tr><tr key=0-0-0-19><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceSeriesName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型类型，如计算型，存储型</span></td><td key=5></td></tr><tr key=0-0-0-20><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型规格, cps.c.normal</span></td><td key=5></td></tr><tr key=0-0-0-21><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceTypeId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型uuid</span></td><td key=5></td></tr><tr key=0-0-0-22><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dvInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-23><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-24><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-25><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuManufacturer</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu厂商</span></td><td key=5></td></tr><tr key=0-0-0-26><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuModel</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu处理器型号</span></td><td key=5></td></tr><tr key=0-0-0-27><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> height</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">显示机型高度</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-28><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-29><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房uuid</span></td><td key=5></td></tr><tr key=0-0-0-30><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房名称</span></td><td key=5></td></tr><tr key=0-0-0-31><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcNameEn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房名称</span></td><td key=5></td></tr><tr key=0-0-0-32><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-33><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceStatus</span></td><td key=1><span>string []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>string</span></p></td></tr><tr key=array-19><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> </span></td><td key=1><span></span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-34><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> interfaceMode</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">bond单网口,dual双网口</span></td><td key=5></td></tr><tr key=0-0-0-35><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-36><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memFrequency</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存主频（MHz)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-37><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-38><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">单个内存大小(GB)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-39><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memSpec</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存 规格,预置还是其它</span></td><td key=5></td></tr><tr key=0-0-0-40><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存接口（如DDR3，DDR4）</span></td><td key=5></td></tr><tr key=0-0-0-41><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型名称，如计算效能型,标准计算型</span></td><td key=5></td></tr><tr key=0-0-0-42><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网卡数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-43><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-44><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicRate</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网卡传输速率(GE)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-45><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> raid</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘raid</span></td><td key=5></td></tr><tr key=0-0-0-46><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> stockAvailable</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">StockAvailable 可用库存</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-47><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> svInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-48><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-49><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeInterfaceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘接口类型（SATA,SAS,NVME）</span></td><td key=5></td></tr><tr key=0-0-0-50><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘单盘大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-51><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘类型（SSD，HDD）</span></td><td key=5></td></tr><tr key=0-0-0-52><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeUnit</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘单位（GB，TB）</span></td><td key=5></td></tr><tr key=0-0-0-53><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-54><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageNumber</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> totalCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr>
               </tbody>
              </table>
            
## /deviceTypes/deviceTypeImage
<a id=/deviceTypes/deviceTypeImage> </a>
### 基本信息

**Path：** /v1/deviceTypes/deviceTypeImage

**Method：** GET

**接口描述：**
QueryDeviceTypeImage 根据机型获取镜像

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| pageNumber | 否  |   |  页码 |
| pageSize | 否  |   |  每页数量 |
| deviceTypeId | 否  |   |  机型id |
| imageId | 否  |   |  镜像id |
| architecture | 否  |   |  体系架构 |
| osType | 否  |   |  操作系统平台 |
| imageName | 否  |   |  镜像名称 |
| version | 否  |   |  版本号 |
| osId | 否  |   |  操作系统ID |
| source | 否  |   |  镜像类型，预置，自定义 |
| isAll | 否  |   |  是否显示全部 |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> images</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像实体列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> architecture</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">架构:x86/x64/i386/</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataPartition</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据分区信息</span></td><td key=5></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">描述</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceTypeNum</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">绑定了机型数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> filename</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像文件名称</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> format</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像格式（qcow2、tar）</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> hash</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像校验码</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>uint64</span></p></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> imageId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像uuid</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> imageName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像名称</span></td><td key=5></td></tr><tr key=0-0-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> isBind</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">是否绑定了某个机型</span></td><td key=5></td></tr><tr key=0-0-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">源os uuid</span></td><td key=5></td></tr><tr key=0-0-0-14><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">CentOS 7.2 64-bit</span></td><td key=5></td></tr><tr key=0-0-0-15><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作系统分类:linux/windows</span></td><td key=5></td></tr><tr key=0-0-0-16><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osVersion</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作系统版本</span></td><td key=5></td></tr><tr key=0-0-0-17><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> source</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像来源(common通用、customize定制、user_defined自定义)</span></td><td key=5></td></tr><tr key=0-0-0-18><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> sourceName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像来源(common通用、customize定制、user_defined自定义)</span></td><td key=5></td></tr><tr key=0-0-0-19><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemPartition</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统分区信息（系统盘：“/ ” ：50GiB，xfs；swap：10GiB，swap）</span></td><td key=5></td></tr><tr key=0-0-0-20><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-21><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-0-22><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> url</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像源路径</span></td><td key=5></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageNumber</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">页数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">页大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> totalCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">总条数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr>
               </tbody>
              </table>
            
## /deviceTypes/deviceTypeImagePartition
<a id=/deviceTypes/deviceTypeImagePartition> </a>
### 基本信息

**Path：** /v1/deviceTypes/deviceTypeImagePartition

**Method：** GET

**接口描述：**
QueryDeviceTypeImagePartition 根据机型，镜像，获取分区

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| deviceTypeId | 否  |   |  机型id |
| imageId | 否  |   |  镜像id |
| isAll | 否  |   |  是否显示所有，isAll=1显示所有 |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> dataPartition</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘分区列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> format</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">format,如[swap, xfs]</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> point</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">point,如[swap, /, /var]</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> size</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">分区大小, MB为单位</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> systemPartition</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘分区列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-1-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> format</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">format,如[swap, xfs]</span></td><td key=5></td></tr><tr key=0-0-1-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> point</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">point,如[swap, /, /var]</span></td><td key=5></td></tr><tr key=0-0-1-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> size</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">分区大小, MB为单位</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr>
               </tbody>
              </table>
            
## /deviceTypes/deviceTypeRaid
<a id=/deviceTypes/deviceTypeRaid> </a>
### 基本信息

**Path：** /v1/deviceTypes/deviceTypeRaid

**Method：** GET

**接口描述：**
DescribeDeviceTypeRaids 根据机型获取raid

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| deviceTypeId | 否  |   |  机型id |
| volumeType | 否  |   |  系统盘还是数据盘 |
| raidId | 否  |   |  raidID |
| isAll | 否  |   |  是否显示所有 |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> DiskInterfaceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">磁盘类型：SAS/SATA/SSD/NVME</span></td><td key=5></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> availableValue</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">可用分区值，单位GB</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-4><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> deviceTypeId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备类型id</span></td><td key=5></td></tr><tr key=0-0-5><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> diskType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">磁盘类型：SAS/SATA/SSD/NVME</span></td><td key=5></td></tr><tr key=0-0-6><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> raidId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">RAID uuid</span></td><td key=5></td></tr><tr key=0-0-7><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> systemPartitionCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘noraid模式真实数量;此模式多块盘只能用一块</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-8><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-9><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-10><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> volumeDetail</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">磁盘详细信息</span></td><td key=5></td></tr><tr key=0-0-11><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> volumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">磁盘类型 system/data</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /deviceTypes/{device_type_id}
<a id=/deviceTypes/{device_type_id}> </a>
### 基本信息

**Path：** /v1/deviceTypes/{device_type_id}

**Method：** GET

**接口描述：**
DescribeDeviceType 获取机型详情

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| device_type_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> deviceType</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> architecture</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">体系架构，如i386/x86_64/ ARM64 (aarch64)，默认 x86_64</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> boot_mode</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">boot模式</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuCores</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">单个cpu内核数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuFrequency</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu频率(G)</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">拼装信息</span></td><td key=5></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuManufacturer</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu厂商</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuModel</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu处理器型号</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuSpec</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">CPU 规格,预置还是其它</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeInterfaceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘接口类型（SATA,SAS,NVME）</span></td><td key=5></td></tr><tr key=0-0-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘单盘大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-14><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘类型（SSD，HDD）</span></td><td key=5></td></tr><tr key=0-0-0-15><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeUnit</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘单位（GB，TB）</span></td><td key=5></td></tr><tr key=0-0-0-16><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">描述</span></td><td key=5></td></tr><tr key=0-0-0-17><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-18><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceSeries</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型类型，如计算型，存储型</span></td><td key=5></td></tr><tr key=0-0-0-19><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceSeriesName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型类型，如计算型，存储型</span></td><td key=5></td></tr><tr key=0-0-0-20><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型规格, cps.c.normal</span></td><td key=5></td></tr><tr key=0-0-0-21><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceTypeId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型uuid</span></td><td key=5></td></tr><tr key=0-0-0-22><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dvInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-23><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-24><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-25><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuManufacturer</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu厂商</span></td><td key=5></td></tr><tr key=0-0-0-26><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuModel</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu处理器型号</span></td><td key=5></td></tr><tr key=0-0-0-27><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> height</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">显示机型高度</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-28><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-29><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房uuid</span></td><td key=5></td></tr><tr key=0-0-0-30><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房名称</span></td><td key=5></td></tr><tr key=0-0-0-31><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcNameEn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房名称</span></td><td key=5></td></tr><tr key=0-0-0-32><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-33><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceStatus</span></td><td key=1><span>string []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>string</span></p></td></tr><tr key=array-20><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> </span></td><td key=1><span></span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-34><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> interfaceMode</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">bond单网口,dual双网口</span></td><td key=5></td></tr><tr key=0-0-0-35><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-36><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memFrequency</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存主频（MHz)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-37><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-38><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">单个内存大小(GB)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-39><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memSpec</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存 规格,预置还是其它</span></td><td key=5></td></tr><tr key=0-0-0-40><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存接口（如DDR3，DDR4）</span></td><td key=5></td></tr><tr key=0-0-0-41><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型名称，如计算效能型,标准计算型</span></td><td key=5></td></tr><tr key=0-0-0-42><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网卡数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-43><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-44><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicRate</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网卡传输速率(GE)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-45><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> raid</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘raid</span></td><td key=5></td></tr><tr key=0-0-0-46><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> stockAvailable</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">StockAvailable 可用库存</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-47><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> svInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-48><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-49><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeInterfaceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘接口类型（SATA,SAS,NVME）</span></td><td key=5></td></tr><tr key=0-0-0-50><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘单盘大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-51><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘类型（SSD，HDD）</span></td><td key=5></td></tr><tr key=0-0-0-52><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeUnit</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘单位（GB，TB）</span></td><td key=5></td></tr><tr key=0-0-0-53><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-54><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
# device

## /devices
<a id=/devices> </a>
### 基本信息

**Path：** /v1/devices

**Method：** GET

**接口描述：**
DescribeDevices 获取设备列表

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| pageNumber | 否  |   |  页码 |
| pageSize | 否  |   |  每页数量 |
| idcId | 否  |   |  机房id |
| sn | 否  |   |  设备sn |
| instanceId | 否  |   |  实例id |
| deviceTypeId | 否  |   |  机型id |
| manageStatus | 否  |   |  设备管理状态 |
| iloIp | 否  |   |  带外ip |
| ipv4 | 否  |   |  ipv4 |
| ipv6 | 否  |   |  ipv6 |
| deviceSeries | 否  |   |  机型类型 |
| userId | 否  |   |  所属用户id |
| userName | 否  |   |  所属用户 |
| isAll | 否  |   |  是否显示全部，1不分页 |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> devices</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> Enclosure2</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘2背板号</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> adapterId</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">adapter_id</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> architecture</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">体系架构，如i386/x86_64/ ARM64 (aarch64)，默认 x86_64</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> brand</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">品牌</span></td><td key=5></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cabinet</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机柜编码</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuCores</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">单个cpu内核数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuFrequency</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu频率(G)</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">拼装信息</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuManufacturer</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu厂商</span></td><td key=5></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuModel</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu处理器型号</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuRoads</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">单个cpu路数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int8</span></p></td></tr><tr key=0-0-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-14><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-15><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeInterfaceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘接口类型（SATA,SAS,NVME）</span></td><td key=5></td></tr><tr key=0-0-0-16><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘单盘大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-17><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘类型（SSD，HDD）</span></td><td key=5></td></tr><tr key=0-0-0-18><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">描述</span></td><td key=5></td></tr><tr key=0-0-0-19><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备uuid</span></td><td key=5></td></tr><tr key=0-0-0-20><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceSeries</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">计算型，存储型</span></td><td key=5></td></tr><tr key=0-0-0-21><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceSeriesName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">计算型，存储型</span></td><td key=5></td></tr><tr key=0-0-0-22><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">型号，机型规格cps.normal</span></td><td key=5></td></tr><tr key=0-0-0-23><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceTypeId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型uuid</span></td><td key=5></td></tr><tr key=0-0-0-24><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceTypeName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型名称</span></td><td key=5></td></tr><tr key=0-0-0-25><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dvInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-26><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> enclosure1</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">添加到disk表, 系统盘1背板号</span></td><td key=5></td></tr><tr key=0-0-0-27><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gateway</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网关地址</span></td><td key=5></td></tr><tr key=0-0-0-28><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-29><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-30><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuManufacturer</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu厂商</span></td><td key=5></td></tr><tr key=0-0-0-31><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuModel</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu处理器型号</span></td><td key=5></td></tr><tr key=0-0-0-32><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备ID编号</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-33><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcId</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房uuid</span></td><td key=5></td></tr><tr key=0-0-0-34><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房名称</span></td><td key=5></td></tr><tr key=0-0-0-35><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcNameEn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">idcname</span></td><td key=5></td></tr><tr key=0-0-0-36><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> iloIp</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">带外管理IP</span></td><td key=5></td></tr><tr key=0-0-0-37><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> iloPassword</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">带外账号密码</span></td><td key=5></td></tr><tr key=0-0-0-38><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> iloUser</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">带外账号</span></td><td key=5></td></tr><tr key=0-0-0-39><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> imageName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例镜像名称</span></td><td key=5></td></tr><tr key=0-0-0-40><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceCreatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例创建时间</span></td><td key=5></td></tr><tr key=0-0-0-41><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceDescription</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例描述</span></td><td key=5></td></tr><tr key=0-0-0-42><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例id</span></td><td key=5></td></tr><tr key=0-0-0-43><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例名称</span></td><td key=5></td></tr><tr key=0-0-0-44><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceReason</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例失败原因</span></td><td key=5></td></tr><tr key=0-0-0-45><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceStatus</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例状态</span></td><td key=5></td></tr><tr key=0-0-0-46><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceStatusName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例状态</span></td><td key=5></td></tr><tr key=0-0-0-47><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> interfaceMode</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">【网口模式】【网络设置】: bond单网口,dual双网口</span></td><td key=5></td></tr><tr key=0-0-0-48><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> locked</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例是否锁定 锁定locked 未锁定unlocked</span></td><td key=5></td></tr><tr key=0-0-0-49><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> mac1</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">MAC1（eth0）</span></td><td key=5></td></tr><tr key=0-0-0-50><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> mac2</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">MAC2（eth2）</span></td><td key=5></td></tr><tr key=0-0-0-51><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> manageStatus</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备状态: 未装机，已装机</span></td><td key=5></td></tr><tr key=0-0-0-52><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> manageStatusName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备状态: 未装机，已装机</span></td><td key=5></td></tr><tr key=0-0-0-53><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> mask</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">子网掩码</span></td><td key=5></td></tr><tr key=0-0-0-54><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-55><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memFrequency</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存主频（MHz)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-56><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-57><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">单个内存大小(GB)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-58><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存接口（如DDR3，DDR4）</span></td><td key=5></td></tr><tr key=0-0-0-59><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> model</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">型号</span></td><td key=5></td></tr><tr key=0-0-0-60><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网卡数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-61><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-62><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicRate</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网卡传输速率(GE)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-63><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> privateIpv4</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内网IPV4</span></td><td key=5></td></tr><tr key=0-0-0-64><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> privateIpv6</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内网IPV6</span></td><td key=5></td></tr><tr key=0-0-0-65><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> raidDriver</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">raid工具：（megacli64等）</span></td><td key=5></td></tr><tr key=0-0-0-66><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> reason</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备状态变更失败原因</span></td><td key=5></td></tr><tr key=0-0-0-67><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> slot1</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘1槽位</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-68><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> slot2</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘2槽位</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-69><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> sn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备SN</span></td><td key=5></td></tr><tr key=0-0-0-70><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> svInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-71><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchIp</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网口交换机IP</span></td><td key=5></td></tr><tr key=0-0-0-72><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchIp1</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机1ip</span></td><td key=5></td></tr><tr key=0-0-0-73><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchIp2</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机2ip</span></td><td key=5></td></tr><tr key=0-0-0-74><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchPassword1</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机1登录密码</span></td><td key=5></td></tr><tr key=0-0-0-75><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchPassword2</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机2登录密码</span></td><td key=5></td></tr><tr key=0-0-0-76><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchPort1</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机1port</span></td><td key=5></td></tr><tr key=0-0-0-77><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchPort2</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机2port</span></td><td key=5></td></tr><tr key=0-0-0-78><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchUser1</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机1登录账号，如果为空，取所在机房的值</span></td><td key=5></td></tr><tr key=0-0-0-79><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchUser2</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机2登录账号，如果为空，取所在机房的值</span></td><td key=5></td></tr><tr key=0-0-0-80><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-81><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeInterfaceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘接口类型（SATA,SAS,NVME）</span></td><td key=5></td></tr><tr key=0-0-0-82><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘单盘大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-83><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘类型（SSD，HDD）</span></td><td key=5></td></tr><tr key=0-0-0-84><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> uPosition</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">U位</span></td><td key=5></td></tr><tr key=0-0-0-85><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-86><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-0-87><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例归属用户uuid</span></td><td key=5></td></tr><tr key=0-0-0-88><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例归属用户名称</span></td><td key=5></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageNumber</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> totalCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr>
               </tbody>
              </table>
            
## /devices/stock
<a id=/devices/stock> </a>
### 基本信息

**Path：** /v1/devices/stock

**Method：** DELETE

**接口描述：**
DescribeDeviceStock 获取指定机型的设备库存

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| deviceTypeId | 否  |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> stocks</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">库存量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr>
               </tbody>
              </table>
            
## /devices/{device_id}
<a id=/devices/{device_id}> </a>
### 基本信息

**Path：** /v1/devices/{device_id}

**Method：** GET

**接口描述：**
DescribeDevice 获取设备详情

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| device_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> device</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> Enclosure2</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘2背板号</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> adapterId</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">adapter_id</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> architecture</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">体系架构，如i386/x86_64/ ARM64 (aarch64)，默认 x86_64</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> brand</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">品牌</span></td><td key=5></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cabinet</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机柜编码</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuCores</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">单个cpu内核数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuFrequency</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu频率(G)</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">拼装信息</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuManufacturer</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu厂商</span></td><td key=5></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuModel</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu处理器型号</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuRoads</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">单个cpu路数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int8</span></p></td></tr><tr key=0-0-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-14><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-15><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeInterfaceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘接口类型（SATA,SAS,NVME）</span></td><td key=5></td></tr><tr key=0-0-0-16><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘单盘大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-17><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘类型（SSD，HDD）</span></td><td key=5></td></tr><tr key=0-0-0-18><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">描述</span></td><td key=5></td></tr><tr key=0-0-0-19><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备uuid</span></td><td key=5></td></tr><tr key=0-0-0-20><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceSeries</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">计算型，存储型</span></td><td key=5></td></tr><tr key=0-0-0-21><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceSeriesName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">计算型，存储型</span></td><td key=5></td></tr><tr key=0-0-0-22><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">型号，机型规格cps.normal</span></td><td key=5></td></tr><tr key=0-0-0-23><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceTypeId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型uuid</span></td><td key=5></td></tr><tr key=0-0-0-24><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceTypeName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型名称</span></td><td key=5></td></tr><tr key=0-0-0-25><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dvInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-26><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> enclosure1</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">添加到disk表, 系统盘1背板号</span></td><td key=5></td></tr><tr key=0-0-0-27><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gateway</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网关地址</span></td><td key=5></td></tr><tr key=0-0-0-28><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-29><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-30><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuManufacturer</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu厂商</span></td><td key=5></td></tr><tr key=0-0-0-31><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuModel</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu处理器型号</span></td><td key=5></td></tr><tr key=0-0-0-32><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备ID编号</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-33><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcId</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房uuid</span></td><td key=5></td></tr><tr key=0-0-0-34><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房名称</span></td><td key=5></td></tr><tr key=0-0-0-35><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcNameEn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">idcname</span></td><td key=5></td></tr><tr key=0-0-0-36><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> iloIp</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">带外管理IP</span></td><td key=5></td></tr><tr key=0-0-0-37><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> iloPassword</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">带外账号密码</span></td><td key=5></td></tr><tr key=0-0-0-38><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> iloUser</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">带外账号</span></td><td key=5></td></tr><tr key=0-0-0-39><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> imageName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例镜像名称</span></td><td key=5></td></tr><tr key=0-0-0-40><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceCreatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例创建时间</span></td><td key=5></td></tr><tr key=0-0-0-41><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceDescription</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例描述</span></td><td key=5></td></tr><tr key=0-0-0-42><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例id</span></td><td key=5></td></tr><tr key=0-0-0-43><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例名称</span></td><td key=5></td></tr><tr key=0-0-0-44><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceReason</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例失败原因</span></td><td key=5></td></tr><tr key=0-0-0-45><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceStatus</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例状态</span></td><td key=5></td></tr><tr key=0-0-0-46><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceStatusName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例状态</span></td><td key=5></td></tr><tr key=0-0-0-47><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> interfaceMode</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">【网口模式】【网络设置】: bond单网口,dual双网口</span></td><td key=5></td></tr><tr key=0-0-0-48><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> locked</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例是否锁定 锁定locked 未锁定unlocked</span></td><td key=5></td></tr><tr key=0-0-0-49><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> mac1</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">MAC1（eth0）</span></td><td key=5></td></tr><tr key=0-0-0-50><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> mac2</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">MAC2（eth2）</span></td><td key=5></td></tr><tr key=0-0-0-51><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> manageStatus</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备状态: 未装机，已装机</span></td><td key=5></td></tr><tr key=0-0-0-52><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> manageStatusName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备状态: 未装机，已装机</span></td><td key=5></td></tr><tr key=0-0-0-53><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> mask</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">子网掩码</span></td><td key=5></td></tr><tr key=0-0-0-54><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-55><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memFrequency</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存主频（MHz)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-56><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-57><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">单个内存大小(GB)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-58><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存接口（如DDR3，DDR4）</span></td><td key=5></td></tr><tr key=0-0-0-59><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> model</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">型号</span></td><td key=5></td></tr><tr key=0-0-0-60><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网卡数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-61><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-62><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicRate</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网卡传输速率(GE)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-63><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> privateIpv4</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内网IPV4</span></td><td key=5></td></tr><tr key=0-0-0-64><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> privateIpv6</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内网IPV6</span></td><td key=5></td></tr><tr key=0-0-0-65><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> raidDriver</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">raid工具：（megacli64等）</span></td><td key=5></td></tr><tr key=0-0-0-66><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> reason</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备状态变更失败原因</span></td><td key=5></td></tr><tr key=0-0-0-67><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> slot1</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘1槽位</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-68><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> slot2</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘2槽位</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-69><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> sn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备SN</span></td><td key=5></td></tr><tr key=0-0-0-70><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> svInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-71><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchIp</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网口交换机IP</span></td><td key=5></td></tr><tr key=0-0-0-72><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchIp1</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机1ip</span></td><td key=5></td></tr><tr key=0-0-0-73><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchIp2</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机2ip</span></td><td key=5></td></tr><tr key=0-0-0-74><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchPassword1</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机1登录密码</span></td><td key=5></td></tr><tr key=0-0-0-75><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchPassword2</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机2登录密码</span></td><td key=5></td></tr><tr key=0-0-0-76><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchPort1</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机1port</span></td><td key=5></td></tr><tr key=0-0-0-77><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchPort2</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机2port</span></td><td key=5></td></tr><tr key=0-0-0-78><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchUser1</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机1登录账号，如果为空，取所在机房的值</span></td><td key=5></td></tr><tr key=0-0-0-79><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchUser2</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机2登录账号，如果为空，取所在机房的值</span></td><td key=5></td></tr><tr key=0-0-0-80><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-81><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeInterfaceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘接口类型（SATA,SAS,NVME）</span></td><td key=5></td></tr><tr key=0-0-0-82><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘单盘大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-83><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘类型（SSD，HDD）</span></td><td key=5></td></tr><tr key=0-0-0-84><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> uPosition</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">U位</span></td><td key=5></td></tr><tr key=0-0-0-85><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-86><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-0-87><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例归属用户uuid</span></td><td key=5></td></tr><tr key=0-0-0-88><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例归属用户名称</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
# idc

## /idcs
<a id=/idcs> </a>
### 基本信息

**Path：** /v1/idcs

**Method：** GET

**接口描述：**
DescribeIdcs 获取idc列表

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| pageNumber | 否  |   |  页码 |
| pageSize | 否  |   |  每页数量 |
| name | 否  |   |  机房名称 |
| nameEn | 否  |   |  机房英文名称 |
| level | 否  |   |  机房等级 |
| isAll | 否  |   |  是否显示所有 |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> idcs</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> address</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房地址</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房uuid</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> iloPassword</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房公用带外管理password</span></td><td key=5></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> iloUser</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房公用带外管理user</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> level</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房等级</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房名称</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nameEn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房名称en</span></td><td key=5></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> shortname</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">shortname</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchPassword1</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机密码1</span></td><td key=5></td></tr><tr key=0-0-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchPassword2</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机密码2</span></td><td key=5></td></tr><tr key=0-0-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchUser1</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机用户名1</span></td><td key=5></td></tr><tr key=0-0-0-14><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchUser2</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机用户名2</span></td><td key=5></td></tr><tr key=0-0-0-15><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updateTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">修改时间</span></td><td key=5></td></tr><tr key=0-0-0-16><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">修改用户</span></td><td key=5></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageNumber</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">页数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">页大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> totalCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">总条数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr>
               </tbody>
              </table>
            
## /idcs/{idc_id}
<a id=/idcs/{idc_id}> </a>
### 基本信息

**Path：** /v1/idcs/{idc_id}

**Method：** GET

**接口描述：**
DescribeIdc 获取idc详情

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| idc_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> idc</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> address</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房地址</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房uuid</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> iloPassword</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房公用带外管理password</span></td><td key=5></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> iloUser</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房公用带外管理user</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> level</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房等级</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房名称</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nameEn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房名称en</span></td><td key=5></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> shortname</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">shortname</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchPassword1</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机密码1</span></td><td key=5></td></tr><tr key=0-0-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchPassword2</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机密码2</span></td><td key=5></td></tr><tr key=0-0-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchUser1</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机用户名1</span></td><td key=5></td></tr><tr key=0-0-0-14><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> switchUser2</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">交换机用户名2</span></td><td key=5></td></tr><tr key=0-0-0-15><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updateTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">修改时间</span></td><td key=5></td></tr><tr key=0-0-0-16><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">修改用户</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
# image

## /images
<a id=/images> </a>
### 基本信息

**Path：** /v1/images

**Method：** GET

**接口描述：**
DescribeImages 获取镜像列表

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| pageNumber | 否  |   |  页码 |
| pageSize | 否  |   |  每页数量 |
| imageId | 否  |   |  镜像ID |
| imageName | 否  |   |  镜像名称 |
| deviceTypeId | 否  |   |  设备类型id |
| version | 否  |   |  版本号 |
| osId | 否  |   |  操作系统ID |
| imageIds | 否  |   |  镜像ID，数组，支持多个 |
| source | 否  |   |  镜像类型，预置，自定义 |
| architecture | 否  |   |  体系架构 |
| osType | 否  |   |  操作系统平台 |
| isAll | 否  |   |  是否显示全部 |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> images</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像实体列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> architecture</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">架构:x86/x64/i386/</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataPartition</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据分区信息</span></td><td key=5></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">描述</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceTypeNum</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">绑定了机型数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> filename</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像文件名称</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> format</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像格式（qcow2、tar）</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> hash</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像校验码</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>uint64</span></p></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> imageId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像uuid</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> imageName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像名称</span></td><td key=5></td></tr><tr key=0-0-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> isBind</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">是否绑定了某个机型</span></td><td key=5></td></tr><tr key=0-0-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">源os uuid</span></td><td key=5></td></tr><tr key=0-0-0-14><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">CentOS 7.2 64-bit</span></td><td key=5></td></tr><tr key=0-0-0-15><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作系统分类:linux/windows</span></td><td key=5></td></tr><tr key=0-0-0-16><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osVersion</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作系统版本</span></td><td key=5></td></tr><tr key=0-0-0-17><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> source</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像来源(common通用、customize定制、user_defined自定义)</span></td><td key=5></td></tr><tr key=0-0-0-18><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> sourceName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像来源(common通用、customize定制、user_defined自定义)</span></td><td key=5></td></tr><tr key=0-0-0-19><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemPartition</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统分区信息（系统盘：“/ ” ：50GiB，xfs；swap：10GiB，swap）</span></td><td key=5></td></tr><tr key=0-0-0-20><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-21><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-0-22><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> url</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像源路径</span></td><td key=5></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageNumber</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">页数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">页大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> totalCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">总条数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr>
               </tbody>
              </table>
            
## /images/imageDeviceTypes
<a id=/images/imageDeviceTypes> </a>
### 基本信息

**Path：** /v1/images/imageDeviceTypes

**Method：** GET

**接口描述：**
DescribeImageDeviceTypes 查看镜像绑定的机型

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| pageNumber | 否  |   |  页码 |
| pageSize | 否  |   |  每页数量 |
| imageId | 是  |   |  镜像ID |
| architecture | 否  |   |  体系架构 |
| isBind | 否  |   |  镜像是否绑定了机型，0查询该镜像没有绑定的机型列表 1查询该镜像已经绑定了的机型列表 |
| isAll | 否  |   |  是否显示全部 |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> deviceTypes</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> architecture</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">体系架构，如i386/x86_64/ ARM64 (aarch64)，默认 x86_64</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> boot_mode</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">boot模式</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuCores</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">单个cpu内核数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuFrequency</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu频率(G)</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">拼装信息</span></td><td key=5></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuManufacturer</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu厂商</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuModel</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu处理器型号</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuSpec</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">CPU 规格,预置还是其它</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeInterfaceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘接口类型（SATA,SAS,NVME）</span></td><td key=5></td></tr><tr key=0-0-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘单盘大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-14><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘类型（SSD，HDD）</span></td><td key=5></td></tr><tr key=0-0-0-15><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeUnit</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘单位（GB，TB）</span></td><td key=5></td></tr><tr key=0-0-0-16><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">描述</span></td><td key=5></td></tr><tr key=0-0-0-17><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-18><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceSeries</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型类型，如计算型，存储型</span></td><td key=5></td></tr><tr key=0-0-0-19><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceSeriesName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型类型，如计算型，存储型</span></td><td key=5></td></tr><tr key=0-0-0-20><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型规格, cps.c.normal</span></td><td key=5></td></tr><tr key=0-0-0-21><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceTypeId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型uuid</span></td><td key=5></td></tr><tr key=0-0-0-22><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dvInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-23><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-24><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-25><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuManufacturer</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu厂商</span></td><td key=5></td></tr><tr key=0-0-0-26><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuModel</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu处理器型号</span></td><td key=5></td></tr><tr key=0-0-0-27><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> height</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">显示机型高度</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-28><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-29><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房uuid</span></td><td key=5></td></tr><tr key=0-0-0-30><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房名称</span></td><td key=5></td></tr><tr key=0-0-0-31><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcNameEn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房名称</span></td><td key=5></td></tr><tr key=0-0-0-32><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-33><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceStatus</span></td><td key=1><span>string []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>string</span></p></td></tr><tr key=array-21><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> </span></td><td key=1><span></span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-34><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> interfaceMode</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">bond单网口,dual双网口</span></td><td key=5></td></tr><tr key=0-0-0-35><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-36><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memFrequency</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存主频（MHz)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-37><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-38><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">单个内存大小(GB)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-39><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memSpec</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存 规格,预置还是其它</span></td><td key=5></td></tr><tr key=0-0-0-40><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存接口（如DDR3，DDR4）</span></td><td key=5></td></tr><tr key=0-0-0-41><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型名称，如计算效能型,标准计算型</span></td><td key=5></td></tr><tr key=0-0-0-42><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网卡数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-43><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-44><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicRate</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网卡传输速率(GE)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-45><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> raid</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘raid</span></td><td key=5></td></tr><tr key=0-0-0-46><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> stockAvailable</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">StockAvailable 可用库存</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-47><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> svInfo</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-48><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-49><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeInterfaceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘接口类型（SATA,SAS,NVME）</span></td><td key=5></td></tr><tr key=0-0-0-50><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘单盘大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-51><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘类型（SSD，HDD）</span></td><td key=5></td></tr><tr key=0-0-0-52><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeUnit</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘单位（GB，TB）</span></td><td key=5></td></tr><tr key=0-0-0-53><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-54><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageNumber</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> totalCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr>
               </tbody>
              </table>
            
## /images/{image_id}
<a id=/images/{image_id}> </a>
### 基本信息

**Path：** /v1/images/{image_id}

**Method：** GET

**接口描述：**
DescribeImage 获取镜像详情

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| image_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> image</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> architecture</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">架构:x86/x64/i386/</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataPartition</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据分区信息</span></td><td key=5></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">描述</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceTypeNum</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">绑定了机型数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> filename</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像文件名称</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> format</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像格式（qcow2、tar）</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> hash</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像校验码</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>uint64</span></p></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> imageId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像uuid</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> imageName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像名称</span></td><td key=5></td></tr><tr key=0-0-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> isBind</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">是否绑定了某个机型</span></td><td key=5></td></tr><tr key=0-0-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">源os uuid</span></td><td key=5></td></tr><tr key=0-0-0-14><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">CentOS 7.2 64-bit</span></td><td key=5></td></tr><tr key=0-0-0-15><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作系统分类:linux/windows</span></td><td key=5></td></tr><tr key=0-0-0-16><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osVersion</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作系统版本</span></td><td key=5></td></tr><tr key=0-0-0-17><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> source</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像来源(common通用、customize定制、user_defined自定义)</span></td><td key=5></td></tr><tr key=0-0-0-18><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> sourceName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像来源(common通用、customize定制、user_defined自定义)</span></td><td key=5></td></tr><tr key=0-0-0-19><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemPartition</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统分区信息（系统盘：“/ ” ：50GiB，xfs；swap：10GiB，swap）</span></td><td key=5></td></tr><tr key=0-0-0-20><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-21><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-0-22><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> url</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像源路径</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
# user

## /local/users
<a id=/local/users> </a>
### 基本信息

**Path：** /v1/local/users

**Method：** GET

**接口描述：**
DescribeLocalUser 控制台获取用户详情

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> error</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> code</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">错误码</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> message</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">错误信息</span></td><td key=5></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> status</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">错误状态</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> user</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-1-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-1-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> defaultProjectId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户默认项目uuid</span></td><td key=5></td></tr><tr key=0-1-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> defaultProjectName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户默认项目名称</span></td><td key=5></td></tr><tr key=0-1-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">描述</span></td><td key=5></td></tr><tr key=0-1-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> email</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">邮箱</span></td><td key=5></td></tr><tr key=0-1-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户id</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-1-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> language</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">默认语言（en_US/zh_CN）</span></td><td key=5></td></tr><tr key=0-1-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> phoneNumber</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">手机号</span></td><td key=5></td></tr><tr key=0-1-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> phonePrefix</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">国家地区码，如86</span></td><td key=5></td></tr><tr key=0-1-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色uuid</span></td><td key=5></td></tr><tr key=0-1-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色名称</span></td><td key=5></td></tr><tr key=0-1-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> timezone</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">时区</span></td><td key=5></td></tr><tr key=0-1-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-1-0-14><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-1-0-15><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户uuid</span></td><td key=5></td></tr><tr key=0-1-0-16><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户名，唯一</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /local/users
<a id=/local/users> </a>
### 基本信息

**Path：** /v1/local/users

**Method：** PUT

**接口描述：**
ModifyLocalUser 控制台修改除密码外的个人信息

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> defaultProjectId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户默认项目uuid</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> email</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">邮箱</span></td><td key=5></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> language</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">语言[en_US, zh_CN]</span></td><td key=5></td></tr><tr key=0-3><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> phoneNumber</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">手机号</span></td><td key=5></td></tr><tr key=0-4><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> phonePrefix</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">国家地区码，如86</span></td><td key=5></td></tr><tr key=0-5><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> timezone</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">时区 Asia/Shanghai</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /local/users/password
<a id=/local/users/password> </a>
### 基本信息

**Path：** /v1/local/users/password

**Method：** PUT

**接口描述：**
ModifyLocalUserPassword 控制台修改个人密码

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> oldPassword</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">旧密码，明文</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> password</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">新密码 明文，密码复杂度校验，大写，小写，数字，特殊字符至少出现3种，且密码长度{8,30}</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /users
<a id=/users> </a>
### 基本信息

**Path：** /v1/users

**Method：** GET

**接口描述：**
DescribeUsers 获取用户列表

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| pageNumber | 否  |   |  页码 |
| pageSize | 否  |   |  每页数量 |
| roleId | 否  |   |  角色uuid |
| defaultProjectId | 否  |   |  项目uuid |
| userName | 否  |   |  用户名 |
| isAll | 否  |   |  是否显示全部, isAll=1表示全部 |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageNumber</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">页数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">页大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> totalCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">总条数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> users</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">user实体列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-3-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-3-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-3-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> defaultProjectId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户默认项目uuid</span></td><td key=5></td></tr><tr key=0-0-3-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> defaultProjectName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户默认项目名称</span></td><td key=5></td></tr><tr key=0-0-3-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">描述</span></td><td key=5></td></tr><tr key=0-0-3-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> email</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">邮箱</span></td><td key=5></td></tr><tr key=0-0-3-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户id</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-3-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> language</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">默认语言（en_US/zh_CN）</span></td><td key=5></td></tr><tr key=0-0-3-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> phoneNumber</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">手机号</span></td><td key=5></td></tr><tr key=0-0-3-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> phonePrefix</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">国家地区码，如86</span></td><td key=5></td></tr><tr key=0-0-3-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色uuid</span></td><td key=5></td></tr><tr key=0-0-3-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色名称</span></td><td key=5></td></tr><tr key=0-0-3-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> timezone</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">时区</span></td><td key=5></td></tr><tr key=0-0-3-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-3-14><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-3-15><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户uuid</span></td><td key=5></td></tr><tr key=0-0-3-16><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户名，唯一</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /users/getUserByName
<a id=/users/getUserByName> </a>
### 基本信息

**Path：** /v1/users/getUserByName

**Method：** GET

**接口描述：**
DescribeUserByName 根据用户名获取用户详情

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| userName | 是  |   |  用户名，唯一 |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> user</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> defaultProjectId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户默认项目uuid</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> defaultProjectName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户默认项目名称</span></td><td key=5></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">描述</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> email</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">邮箱</span></td><td key=5></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户id</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> language</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">默认语言（en_US/zh_CN）</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> phoneNumber</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">手机号</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> phonePrefix</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">国家地区码，如86</span></td><td key=5></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色uuid</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色名称</span></td><td key=5></td></tr><tr key=0-0-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> timezone</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">时区</span></td><td key=5></td></tr><tr key=0-0-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-14><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-0-15><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户uuid</span></td><td key=5></td></tr><tr key=0-0-0-16><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户名，唯一</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /users/verify
<a id=/users/verify> </a>
### 基本信息

**Path：** /v1/users/verify

**Method：** POST

**接口描述：**
VerifyUser 鉴定用户

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> password</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">密码</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> roleId</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色uuid</span></td><td key=5></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> userName</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户名，唯一</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /users/{user_id}
<a id=/users/{user_id}> </a>
### 基本信息

**Path：** /v1/users/{user_id}

**Method：** GET

**接口描述：**
DescribeUser 获取用户详情

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| user_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> error</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> code</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">错误码</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> message</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">错误信息</span></td><td key=5></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> status</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">错误状态</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> user</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-1-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-1-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> defaultProjectId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户默认项目uuid</span></td><td key=5></td></tr><tr key=0-1-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> defaultProjectName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户默认项目名称</span></td><td key=5></td></tr><tr key=0-1-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">描述</span></td><td key=5></td></tr><tr key=0-1-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> email</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">邮箱</span></td><td key=5></td></tr><tr key=0-1-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户id</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-1-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> language</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">默认语言（en_US/zh_CN）</span></td><td key=5></td></tr><tr key=0-1-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> phoneNumber</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">手机号</span></td><td key=5></td></tr><tr key=0-1-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> phonePrefix</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">国家地区码，如86</span></td><td key=5></td></tr><tr key=0-1-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色uuid</span></td><td key=5></td></tr><tr key=0-1-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色名称</span></td><td key=5></td></tr><tr key=0-1-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> timezone</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">时区</span></td><td key=5></td></tr><tr key=0-1-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-1-0-14><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-1-0-15><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户uuid</span></td><td key=5></td></tr><tr key=0-1-0-16><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户名，唯一</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /users/{user_id}
<a id=/users/{user_id}> </a>
### 基本信息

**Path：** /v1/users/{user_id}

**Method：** PUT

**接口描述：**
ModifyUser 修改用户信息

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| user_id |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">描述</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> email</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">邮箱</span></td><td key=5></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> phoneNumber</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">手机号</span></td><td key=5></td></tr><tr key=0-3><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> phonePrefix</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">国家地区码，如86</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
# os

## /oss
<a id=/oss> </a>
### 基本信息

**Path：** /v1/oss

**Method：** GET

**接口描述：**
DescribeOSs 获取os列表(暂不启用)

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| osName | 否  |   |  操作系统名称 |
| osType | 否  |   |  操作系统平台 |
| osVersion | 否  |   |  操作系统版本 |
| isAll | 否  |   |  是否显示所有 |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> oss</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作系统列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> architecture</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">架构:x86/x64/i386/</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> bits</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">指令宽度:64/32位</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deletedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">删除时间</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">ID</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>uint64</span></p></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> isDel</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">是否删除0未删除 1已删除</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int8</span></p></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作系统uuid</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作系统名称</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作系统分类:linux/windows</span></td><td key=5></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osVersion</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作系统版本</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> sysUser</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">管理员账户</span></td><td key=5></td></tr><tr key=0-0-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /oss/{os_id}
<a id=/oss/{os_id}> </a>
### 基本信息

**Path：** /v1/oss/{os_id}

**Method：** GET

**接口描述：**
DescribeOS 获取os系统详情(暂不启用)

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| os_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> os</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> architecture</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">架构:x86/x64/i386/</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> bits</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">指令宽度:64/32位</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deletedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">删除时间</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">ID</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>uint64</span></p></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> isDel</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">是否删除0未删除 1已删除</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int8</span></p></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作系统uuid</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作系统名称</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作系统分类:linux/windows</span></td><td key=5></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> osVersion</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作系统版本</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> sysUser</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">管理员账户</span></td><td key=5></td></tr><tr key=0-0-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
# instance

## /project/instances
<a id=/project/instances> </a>
### 基本信息

**Path：** /v1/project/instances

**Method：** GET

**接口描述：**
DescribeProjectInstances 获取实例列表

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| pageNumber | 否  |   |  页码 |
| pageSize | 否  |   |  每页数量 |
| instanceId | 否  |   |  实例ID |
| instanceName | 否  |   |  实例名称,模糊搜索 |
| deviceTypeId | 否  |   |  设备类型ID |
| status | 否  |   |  运行状态 |
| deviceId | 否  |   |  设备ID |
| sn | 否  |   |  SN |
| idcId | 否  |   |  机房ID |
| projectId | 否  |   |  项目ID |
| ilo_ip | 否  |   |  带外ip，精确搜索 |
| ipv4 | 否  |   |  ipv4，精确搜索 |
| ipv6 | 否  |   |  ipv6，精确搜索 |
| isAll | 否  |   |  是否显示全部，取值为1时表示全部 |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> instances</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">instance实体数组</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> DataVolumeRaidName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘raid名称</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuCores</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">单个cpu内核数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuFrequency</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu频率(GHz)</span></td><td key=5></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuManufacturer</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu厂商</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuModel</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu处理器型号</span></td><td key=5></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeInterfaceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘接口类型（SATA,SAS,NVME）</span></td><td key=5></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeRaidId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘raidId</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘单盘大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘类型（SSD，HDD）</span></td><td key=5></td></tr><tr key=0-0-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeUnit</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘单位</span></td><td key=5></td></tr><tr key=0-0-0-14><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例描述</span></td><td key=5></td></tr><tr key=0-0-0-15><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备uuid</span></td><td key=5></td></tr><tr key=0-0-0-16><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceSeries</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型类型，如computer</span></td><td key=5></td></tr><tr key=0-0-0-17><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceSeriesName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型类型，如计算型，存储型</span></td><td key=5></td></tr><tr key=0-0-0-18><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型规格, cps.c.normal</span></td><td key=5></td></tr><tr key=0-0-0-19><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceTypeId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型uuid</span></td><td key=5></td></tr><tr key=0-0-0-20><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceTypeName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型名称，如计算效能型,标准计算型</span></td><td key=5></td></tr><tr key=0-0-0-21><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-22><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuManufacturer</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu厂商</span></td><td key=5></td></tr><tr key=0-0-0-23><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuModel</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu处理器型号</span></td><td key=5></td></tr><tr key=0-0-0-24><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> hostname</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">主机名</span></td><td key=5></td></tr><tr key=0-0-0-25><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房uuid</span></td><td key=5></td></tr><tr key=0-0-0-26><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房名称</span></td><td key=5></td></tr><tr key=0-0-0-27><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> iloIp</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">带外管理IP</span></td><td key=5></td></tr><tr key=0-0-0-28><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> imageId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像uuid</span></td><td key=5></td></tr><tr key=0-0-0-29><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> imageName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像名称</span></td><td key=5></td></tr><tr key=0-0-0-30><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例uuid</span></td><td key=5></td></tr><tr key=0-0-0-31><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例名称</span></td><td key=5></td></tr><tr key=0-0-0-32><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> interfaceMode</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">【网口模式】【网络设置】: bond单网口,dual双网口</span></td><td key=5></td></tr><tr key=0-0-0-33><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> locked</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">是否锁定解锁锁定:locked,解锁unlocked</span></td><td key=5></td></tr><tr key=0-0-0-34><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> lockedName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">是否锁定解锁锁定:已解锁,已锁定</span></td><td key=5></td></tr><tr key=0-0-0-35><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-36><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memFrequency</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存主频（MHz)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-37><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">单个内存大小(GB)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-38><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存接口（如DDR3，DDR4）</span></td><td key=5></td></tr><tr key=0-0-0-39><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网卡数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-40><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicRate</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网卡传输速率(GE)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-41><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> privateIpv4</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内网IPV4</span></td><td key=5></td></tr><tr key=0-0-0-42><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> privateIpv6</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内网IPV6</span></td><td key=5></td></tr><tr key=0-0-0-43><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> projectId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例所属项目UUID</span></td><td key=5></td></tr><tr key=0-0-0-44><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> reason</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例错误状态时的错误原因</span></td><td key=5></td></tr><tr key=0-0-0-45><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> sn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备SN</span></td><td key=5></td></tr><tr key=0-0-0-46><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> status</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">运行状态</span></td><td key=5></td></tr><tr key=0-0-0-47><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> statusName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">运行状态中文名字</span></td><td key=5></td></tr><tr key=0-0-0-48><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-49><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeInterfaceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘接口类型（SATA,SAS,NVME）</span></td><td key=5></td></tr><tr key=0-0-0-50><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeRaidId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘raidId</span></td><td key=5></td></tr><tr key=0-0-0-51><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeRaidName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘raid名称</span></td><td key=5></td></tr><tr key=0-0-0-52><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘单盘大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-53><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘类型（SSD，HDD）</span></td><td key=5></td></tr><tr key=0-0-0-54><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeUnit</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘单位</span></td><td key=5></td></tr><tr key=0-0-0-55><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-56><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-0-57><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例所属用户UUID</span></td><td key=5></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageNumber</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> totalCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr>
               </tbody>
              </table>
            
## /project/instances
<a id=/project/instances> </a>
### 基本信息

**Path：** /v1/project/instances

**Method：** POST

**接口描述：**
CreateProjectInstance 创建实例

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> boot_mode</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> count</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">描述</span></td><td key=5></td></tr><tr key=0-3><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> deviceTypeId</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备类型ID</span></td><td key=5></td></tr><tr key=0-4><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> hostname</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">主机名称</span></td><td key=5></td></tr><tr key=0-5><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> idcId</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房ID</span></td><td key=5></td></tr><tr key=0-6><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> imageId</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像ID</span></td><td key=5></td></tr><tr key=0-7><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> instanceName</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例名称</span></td><td key=5></td></tr><tr key=0-8><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> password</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">密码</span></td><td key=5></td></tr><tr key=0-9><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> projectId</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目ID</span></td><td key=5></td></tr><tr key=0-10><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> sshKeyId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">秘钥ID</span></td><td key=5></td></tr><tr key=0-11><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> systemPartition</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘分区</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-11-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> format</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">format,如[swap, xfs]</span></td><td key=5></td></tr><tr key=0-11-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> point</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">point,如[swap, /, /var]</span></td><td key=5></td></tr><tr key=0-11-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> size</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">分区大小, MB为单位</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-12><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> systemVolumeRaidId</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘RAID ID</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> instanceIds</span></td><td key=1><span>string []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">instanceId 列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>string</span></p></td></tr><tr key=array-22><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> </span></td><td key=1><span></span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /project/instances/batch/instances:deleteInstances
<a id=/project/instances/batch/instances:deleteInstances> </a>
### 基本信息

**Path：** /v1/project/instances/batch/instances:deleteInstances

**Method：** DELETE

**接口描述：**
DeleteInstances 批量删除实例

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> instance_ids</span></td><td key=1><span>string []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例id列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>string</span></p></td></tr><tr key=array-23><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> </span></td><td key=1><span></span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /project/instances/batch/instances:modifyInstances
<a id=/project/instances/batch/instances:modifyInstances> </a>
### 基本信息

**Path：** /v1/project/instances/batch/instances:modifyInstances

**Method：** PUT

**接口描述：**
ModifyInstances 批量修改实例名称

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> instanceName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例名称</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> instance_ids</span></td><td key=1><span>string []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例id列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>string</span></p></td></tr><tr key=array-24><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> </span></td><td key=1><span></span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /project/instances/batch/instances:restartInstances
<a id=/project/instances/batch/instances:restartInstances> </a>
### 基本信息

**Path：** /v1/project/instances/batch/instances:restartInstances

**Method：** PUT

**接口描述：**
StartInstances 批量重启

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> instanceIds</span></td><td key=1><span>string []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>string</span></p></td></tr><tr key=array-25><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> </span></td><td key=1><span></span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /project/instances/batch/instances:startInstances
<a id=/project/instances/batch/instances:startInstances> </a>
### 基本信息

**Path：** /v1/project/instances/batch/instances:startInstances

**Method：** PUT

**接口描述：**
StartInstances 批量开机

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> instanceIds</span></td><td key=1><span>string []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>string</span></p></td></tr><tr key=array-26><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> </span></td><td key=1><span></span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /project/instances/batch/instances:stopInstances
<a id=/project/instances/batch/instances:stopInstances> </a>
### 基本信息

**Path：** /v1/project/instances/batch/instances:stopInstances

**Method：** PUT

**接口描述：**
StartInstances 批量关机

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> instanceIds</span></td><td key=1><span>string []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>string</span></p></td></tr><tr key=array-27><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> </span></td><td key=1><span></span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /project/instances/{instance_id}
<a id=/project/instances/{instance_id}> </a>
### 基本信息

**Path：** /v1/project/instances/{instance_id}

**Method：** GET

**接口描述：**
DescribeProjectInstance 获取实例详情

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| instance_id |   |  实例uuid |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> instance</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> DataVolumeRaidName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘raid名称</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuCores</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">单个cpu内核数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuFrequency</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu频率(GHz)</span></td><td key=5></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuManufacturer</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu厂商</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> cpuModel</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">cpu处理器型号</span></td><td key=5></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeInterfaceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘接口类型（SATA,SAS,NVME）</span></td><td key=5></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeRaidId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘raidId</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘单盘大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据盘类型（SSD，HDD）</span></td><td key=5></td></tr><tr key=0-0-0-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> dataVolumeUnit</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘单位</span></td><td key=5></td></tr><tr key=0-0-0-14><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例描述</span></td><td key=5></td></tr><tr key=0-0-0-15><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备uuid</span></td><td key=5></td></tr><tr key=0-0-0-16><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceSeries</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型类型，如computer</span></td><td key=5></td></tr><tr key=0-0-0-17><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceSeriesName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型类型，如计算型，存储型</span></td><td key=5></td></tr><tr key=0-0-0-18><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型规格, cps.c.normal</span></td><td key=5></td></tr><tr key=0-0-0-19><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceTypeId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型uuid</span></td><td key=5></td></tr><tr key=0-0-0-20><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceTypeName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机型名称，如计算效能型,标准计算型</span></td><td key=5></td></tr><tr key=0-0-0-21><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-22><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuManufacturer</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu厂商</span></td><td key=5></td></tr><tr key=0-0-0-23><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> gpuModel</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">gpu处理器型号</span></td><td key=5></td></tr><tr key=0-0-0-24><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> hostname</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">主机名</span></td><td key=5></td></tr><tr key=0-0-0-25><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房uuid</span></td><td key=5></td></tr><tr key=0-0-0-26><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> idcName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">机房名称</span></td><td key=5></td></tr><tr key=0-0-0-27><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> iloIp</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">带外管理IP</span></td><td key=5></td></tr><tr key=0-0-0-28><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> imageId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像uuid</span></td><td key=5></td></tr><tr key=0-0-0-29><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> imageName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">镜像名称</span></td><td key=5></td></tr><tr key=0-0-0-30><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例uuid</span></td><td key=5></td></tr><tr key=0-0-0-31><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例名称</span></td><td key=5></td></tr><tr key=0-0-0-32><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> interfaceMode</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">【网口模式】【网络设置】: bond单网口,dual双网口</span></td><td key=5></td></tr><tr key=0-0-0-33><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> locked</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">是否锁定解锁锁定:locked,解锁unlocked</span></td><td key=5></td></tr><tr key=0-0-0-34><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> lockedName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">是否锁定解锁锁定:已解锁,已锁定</span></td><td key=5></td></tr><tr key=0-0-0-35><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-36><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memFrequency</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存主频（MHz)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-37><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">单个内存大小(GB)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-38><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> memType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内存接口（如DDR3，DDR4）</span></td><td key=5></td></tr><tr key=0-0-0-39><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网卡数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-40><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> nicRate</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">网卡传输速率(GE)</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-41><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> privateIpv4</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内网IPV4</span></td><td key=5></td></tr><tr key=0-0-0-42><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> privateIpv6</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">内网IPV6</span></td><td key=5></td></tr><tr key=0-0-0-43><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> projectId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例所属项目UUID</span></td><td key=5></td></tr><tr key=0-0-0-44><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> reason</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例错误状态时的错误原因</span></td><td key=5></td></tr><tr key=0-0-0-45><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> sn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备SN</span></td><td key=5></td></tr><tr key=0-0-0-46><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> status</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">运行状态</span></td><td key=5></td></tr><tr key=0-0-0-47><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> statusName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">运行状态中文名字</span></td><td key=5></td></tr><tr key=0-0-0-48><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeAmount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-49><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeInterfaceType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘接口类型（SATA,SAS,NVME）</span></td><td key=5></td></tr><tr key=0-0-0-50><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeRaidId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘raidId</span></td><td key=5></td></tr><tr key=0-0-0-51><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeRaidName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘raid名称</span></td><td key=5></td></tr><tr key=0-0-0-52><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘单盘大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-53><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘类型（SSD，HDD）</span></td><td key=5></td></tr><tr key=0-0-0-54><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemVolumeUnit</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘单位</span></td><td key=5></td></tr><tr key=0-0-0-55><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-56><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-0-57><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例所属用户UUID</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /project/instances/{instance_id}
<a id=/project/instances/{instance_id}> </a>
### 基本信息

**Path：** /v1/project/instances/{instance_id}

**Method：** PUT

**接口描述：**
ModifyProjectInstance 修改实例信息

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| instance_id |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">描述</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> instanceName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例名称</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /project/instances/{instance_id}
<a id=/project/instances/{instance_id}> </a>
### 基本信息

**Path：** /v1/project/instances/{instance_id}

**Method：** DELETE

**接口描述：**
DeleteProjectInstance 删除实例

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| instance_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /project/instances/{instance_id}/lock
<a id=/project/instances/{instance_id}/lock> </a>
### 基本信息

**Path：** /v1/project/instances/{instance_id}/lock

**Method：** PUT

**接口描述：**
LockProjectInstance 锁定实例

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| instance_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /project/instances/{instance_id}/resetStatus
<a id=/project/instances/{instance_id}/resetStatus> </a>
### 基本信息

**Path：** /v1/project/instances/{instance_id}/resetStatus

**Method：** PUT

**接口描述：**
ResetInstanceStatus 重置实例状态

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| instance_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /project/instances/{instance_id}/restart
<a id=/project/instances/{instance_id}/restart> </a>
### 基本信息

**Path：** /v1/project/instances/{instance_id}/restart

**Method：** PUT

**接口描述：**
RestartProjectInstance 实例重启

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| instance_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /project/instances/{instance_id}/start
<a id=/project/instances/{instance_id}/start> </a>
### 基本信息

**Path：** /v1/project/instances/{instance_id}/start

**Method：** PUT

**接口描述：**
StartProjectInstance 实例开机

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| instance_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /project/instances/{instance_id}/stop
<a id=/project/instances/{instance_id}/stop> </a>
### 基本信息

**Path：** /v1/project/instances/{instance_id}/stop

**Method：** PUT

**接口描述：**
StopProjectInstance 实例关机

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| instance_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /project/instances/{instance_id}/unlock
<a id=/project/instances/{instance_id}/unlock> </a>
### 基本信息

**Path：** /v1/project/instances/{instance_id}/unlock

**Method：** PUT

**接口描述：**
UnLockProjectInstance 解锁实例

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| instance_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
# raid

## /raids
<a id=/raids> </a>
### 基本信息

**Path：** /v1/raids

**Method：** GET

**接口描述：**
DescribeRaids 获取raid列表

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> raids</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">raid实体列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> availableValue</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">可用分区值，单位GB</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> descriptionEn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">中文描述</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> descriptionZh</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">description</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deviceTypeId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备类型uuid</span></td><td key=5></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> diskType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">磁盘类型：SAS/SATA/SSD/NVME</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">raid 名称</span></td><td key=5></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> raidId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">raid uuid</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> systemPartitionCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘分区数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> volumeDetail</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">磁盘详细信息</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> volumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">磁盘类型</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /raids/{raid_id}
<a id=/raids/{raid_id}> </a>
### 基本信息

**Path：** /v1/raids/{raid_id}

**Method：** GET

**接口描述：**
DescribeRaid 获取raid详情

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| raid_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> availableValue</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">可用分区值，单位GB</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> descriptionEn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">中文描述</span></td><td key=5></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> descriptionZh</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">description</span></td><td key=5></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> deviceTypeId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">设备类型uuid</span></td><td key=5></td></tr><tr key=0-0-4><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> diskType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">磁盘类型：SAS/SATA/SSD/NVME</span></td><td key=5></td></tr><tr key=0-0-5><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">raid 名称</span></td><td key=5></td></tr><tr key=0-0-6><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> raidId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">raid uuid</span></td><td key=5></td></tr><tr key=0-0-7><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> systemPartitionCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">系统盘分区数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-8><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> volumeDetail</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">磁盘详细信息</span></td><td key=5></td></tr><tr key=0-0-9><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> volumeType</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">磁盘类型</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
# resource

## /resources
<a id=/resources> </a>
### 基本信息

**Path：** /v1/resources

**Method：** GET

**接口描述：**
DescribeResources 根据某个字段值精确查询对应的资源列表（如根据imageName查询镜像列表，根据idcName查询机房列表）

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| name | 否  |   |  机房名称<br/>IdcName string `json:"idcName"`<br/>机房英文名称<br/>IdcNameEn string `json:"idcNameEn"`*/<br/>机型名称 |
| deviceType | 否  |   |  机型规格 |
| imageName | 否  |   |  镜像名称 |
| userName | 否  |   |  用户名称 |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
# role

## /roles
<a id=/roles> </a>
### 基本信息

**Path：** /v1/roles

**Method：** GET

**接口描述：**
DescribeRoles 获取角色列表

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| pageNumber | 否  |   |  页码 |
| pageSize | 否  |   |  每页数量 |
| isAll | 否  |   |  是否显示所有 |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageNumber</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> roles</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色实体列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-2-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-2-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-2-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> descriptionCn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">description</span></td><td key=5></td></tr><tr key=0-0-2-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> descriptionEn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">权限描述</span></td><td key=5></td></tr><tr key=0-0-2-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> permission</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">权限</span></td><td key=5></td></tr><tr key=0-0-2-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色uuid</span></td><td key=5></td></tr><tr key=0-0-2-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色名称</span></td><td key=5></td></tr><tr key=0-0-2-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleNameCn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色名称，唯一</span></td><td key=5></td></tr><tr key=0-0-2-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleNameEn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色名称，唯一</span></td><td key=5></td></tr><tr key=0-0-2-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-2-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-2-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> totalCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr>
               </tbody>
              </table>
            
## /roles/roleInfo/current
<a id=/roles/roleInfo/current> </a>
### 基本信息

**Path：** /v1/roles/roleInfo/current

**Method：** GET

**接口描述：**
CurrentRole 获取当前登录用户的角色

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> role</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> descriptionCn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">description</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> descriptionEn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">权限描述</span></td><td key=5></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> permission</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">权限</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色uuid</span></td><td key=5></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色名称</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleNameCn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色名称，唯一</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleNameEn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色名称，唯一</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr>
               </tbody>
              </table>
            
## /roles/{role_id}
<a id=/roles/{role_id}> </a>
### 基本信息

**Path：** /v1/roles/{role_id}

**Method：** GET

**接口描述：**
DescribeRole 获取角色详情

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| role_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> role</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> descriptionCn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">description</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> descriptionEn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">权限描述</span></td><td key=5></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> permission</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">权限</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色uuid</span></td><td key=5></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色名称</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleNameCn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色名称，唯一</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> roleNameEn</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">角色名称，唯一</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-0-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr>
               </tbody>
              </table>
            
# apiKey

## /user/apikeys
<a id=/user/apikeys> </a>
### 基本信息

**Path：** /v1/user/apikeys

**Method：** GET

**接口描述：**
DescribeUserAPIKeys 获取APIKey列表

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| pageNumber | 否  |   |  页码 |
| pageSize | 否  |   |  每页数量 |
| name | 否  |   |  秘钥对名称 |
| type | 否  |   |  Token类型, [system/user] |
| isAll | 否  |   |  是否查询全部/导出 |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> apikeys</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> apiKeyId</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">apikey uuid</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">名称</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> readOnly</span></td><td key=1><span>integer</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">是否支持只读, [0/1], 1表示只读</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int8</span></p></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> token</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">32位字符令牌，使用token来独立访问openapi</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> type</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">apikey的类型，[system/user]</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userId</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">所属用户uuid</span></td><td key=5></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageNumber</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> totalCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr>
               </tbody>
              </table>
            
## /user/apikeys
<a id=/user/apikeys> </a>
### 基本信息

**Path：** /v1/user/apikeys

**Method：** POST

**接口描述：**
CreateUserApikey 创建apikey

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">秘钥对名称</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> readOnly</span></td><td key=1><span>integer</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">是否支持只读，[0/1], read_only=1 的时候说明这个key是只读key，不能访问写方法</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int8</span></p></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> type</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">Token类型, [system/user]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> apikeyId</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">apikey uuid</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /user/apikeys/{apikey_id}
<a id=/user/apikeys/{apikey_id}> </a>
### 基本信息

**Path：** /v1/user/apikeys/{apikey_id}

**Method：** GET

**接口描述：**
DescribeUserAPIKey 获取apikey详情

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| apikey_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> apikey</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> apiKeyId</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">apikey uuid</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">名称</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> readOnly</span></td><td key=1><span>integer</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">是否支持只读, [0/1], 1表示只读</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int8</span></p></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> token</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">32位字符令牌，使用token来独立访问openapi</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> type</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">apikey的类型，[system/user]</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-0-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userId</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">所属用户uuid</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /user/apikeys/{apikey_id}
<a id=/user/apikeys/{apikey_id}> </a>
### 基本信息

**Path：** /v1/user/apikeys/{apikey_id}

**Method：** PUT

**接口描述：**
ModifyUserApikey 修改apikey(暂不启用)

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| apikey_id |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">秘钥对名称</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> readOnly</span></td><td key=1><span>integer</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">是否支持只读，[0/1], read_only=1 的时候说明这个key是只读key，不能访问写方法</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int8</span></p></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /user/apikeys/{apikey_id}
<a id=/user/apikeys/{apikey_id}> </a>
### 基本信息

**Path：** /v1/user/apikeys/{apikey_id}

**Method：** DELETE

**接口描述：**
DeleteUserApikey 删除某个apikey

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| apikey_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
# project

## /user/projects
<a id=/user/projects> </a>
### 基本信息

**Path：** /v1/user/projects

**Method：** GET

**接口描述：**
DescribeUserProjects 获取项目列表

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| pageNumber | 否  |   |  页码 |
| pageSize | 否  |   |  每页数量 |
| projectName | 否  |   |  项目名称 1~64字符，只支持数字、大小写字母、英文下划线“_”及中划线“-” |
| owned | 否  |   |  0表示全部，1表示拥有者，2表示被共享者 |
| isAll | 否  |   |   |
| orderByCreatetime | 否  |   |  按创建时间asc/desc排列 |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageNumber</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">页数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">页大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> projects</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目实体列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-2-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-2-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-2-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目描述,新增</span></td><td key=5></td></tr><tr key=0-0-2-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目下实例数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-2-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> owned</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">此项目的共享标志位，1为拥有，2为共享</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-2-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> projectId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目uuid</span></td><td key=5></td></tr><tr key=0-0-2-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> projectName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目名称</span></td><td key=5></td></tr><tr key=0-0-2-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> shareProjects</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">共享</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-2-7-0><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-2-7-1><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> ownerUserId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目拥有者用户id</span></td><td key=5></td></tr><tr key=0-0-2-7-2><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> ownerUserName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目拥有者用户名</span></td><td key=5></td></tr><tr key=0-0-2-7-3><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> projectId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目uuid</span></td><td key=5></td></tr><tr key=0-0-2-7-4><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> projectName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目名称</span></td><td key=5></td></tr><tr key=0-0-2-7-5><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> shareProjects</span></td><td key=1><span> []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=array-28><td key=0><span style="padding-left: 80px"><span style="color: #8c8a8a">├─</span> </span></td><td key=1><span></span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-2-7-6><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> sharedUserId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目共享者用户id</span></td><td key=5></td></tr><tr key=0-0-2-7-7><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> sharedUserName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目拥有者用户名</span></td><td key=5></td></tr><tr key=0-0-2-7-8><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-2-7-9><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-2-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-2-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> totalCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">总条数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr>
               </tbody>
              </table>
            
## /user/projects
<a id=/user/projects> </a>
### 基本信息

**Path：** /v1/user/projects

**Method：** POST

**接口描述：**
CreateUserProject 创建项目

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> isDefault</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">是否作为该用户的默认项目, [0/1], 默认为0</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int8</span></p></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> isSystem</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">是否作为系统项目, [0/1], 默认为0</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int8</span></p></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> projectName</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">ProjectName 1~64字符，只支持数字、大小写字母、中英文下划线“_”及中划线“-”</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> projectId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目uuid</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /user/projects/{project_id}
<a id=/user/projects/{project_id}> </a>
### 基本信息

**Path：** /v1/user/projects/{project_id}

**Method：** GET

**接口描述：**
DescribeUserProject 获取项目详情

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| project_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> project</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> description</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目描述,新增</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> instanceCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目下实例数量</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> owned</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">此项目的共享标志位，1为拥有，2为共享</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> projectId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目uuid</span></td><td key=5></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> projectName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目名称</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> shareProjects</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">共享</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0-7-0><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-7-1><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> ownerUserId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目拥有者用户id</span></td><td key=5></td></tr><tr key=0-0-0-7-2><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> ownerUserName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目拥有者用户名</span></td><td key=5></td></tr><tr key=0-0-0-7-3><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> projectId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目uuid</span></td><td key=5></td></tr><tr key=0-0-0-7-4><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> projectName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目名称</span></td><td key=5></td></tr><tr key=0-0-0-7-5><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> shareProjects</span></td><td key=1><span> []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=array-29><td key=0><span style="padding-left: 80px"><span style="color: #8c8a8a">├─</span> </span></td><td key=1><span></span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-7-6><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> sharedUserId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目共享者用户id</span></td><td key=5></td></tr><tr key=0-0-0-7-7><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> sharedUserName</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目拥有者用户名</span></td><td key=5></td></tr><tr key=0-0-0-7-8><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-7-9><td key=0><span style="padding-left: 60px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /user/projects/{project_id}
<a id=/user/projects/{project_id}> </a>
### 基本信息

**Path：** /v1/user/projects/{project_id}

**Method：** PUT

**接口描述：**
ModifyUserProject 修改项目信息

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| project_id |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> projectName</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目名称 1~64字符，只支持数字、大小写字母、英文下划线“_”及中划线“-”</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /user/projects/{project_id}
<a id=/user/projects/{project_id}> </a>
### 基本信息

**Path：** /v1/user/projects/{project_id}

**Method：** DELETE

**接口描述：**
DeleteUserProject 删除项目

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| project_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /user/projects/{project_id}/cancelShare
<a id=/user/projects/{project_id}/cancelShare> </a>
### 基本信息

**Path：** /v1/user/projects/{project_id}/cancelShare

**Method：** PUT

**接口描述：**
CancelShareUserProject 取消共享项目

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| project_id |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> ownerID</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">from user_id</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> sharerID</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">to user_id</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /user/projects/{project_id}/describeSharProject
<a id=/user/projects/{project_id}/describeSharProject> </a>
### 基本信息

**Path：** /v1/user/projects/{project_id}/describeSharProject

**Method：** GET

**接口描述：**
DescribeShareUserProject 获取共享项目详情

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| project_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> projectId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目实体</span></td><td key=5></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> shares</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-1-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> created_by</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-1-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> created_time</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-1-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> deleted_time</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-1-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> is_default</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int8</span></p></td></tr><tr key=0-0-1-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> is_del</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int8</span></p></td></tr><tr key=0-0-1-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> owner_user_id</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-1-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> owner_user_name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-1-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> premission</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-1-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> project_id</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-1-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> project_name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-1-10><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> shared_user_id</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-1-11><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> shared_user_name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-1-12><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updated_by</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-1-13><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updated_time</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr>
               </tbody>
              </table>
            
## /user/projects/{project_id}/description
<a id=/user/projects/{project_id}/description> </a>
### 基本信息

**Path：** /v1/user/projects/{project_id}/description

**Method：** PUT

**接口描述：**
ModifyUserProject 修改项目信息

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| project_id |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> description</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">项目描述</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /user/projects/{project_id}/share
<a id=/user/projects/{project_id}/share> </a>
### 基本信息

**Path：** /v1/user/projects/{project_id}/share

**Method：** PUT

**接口描述：**
ShareUserProject 共享项目

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| project_id |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> ownerID</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">from user_id</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> sharerID</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">to user_id</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
# sshkey

## /user/sshkeys
<a id=/user/sshkeys> </a>
### 基本信息

**Path：** /v1/user/sshkeys

**Method：** GET

**接口描述：**
DescribeUserSshKeys 获取个人sshkey列表

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Query**

| 参数名称  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
| pageNumber | 否  |   |  页码 |
| pageSize | 否  |   |  每页数量 |
| name | 否  |   |  秘钥名称 |
| isAll | 否  |   |  是否显示全部 |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageNumber</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">页数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pageSize</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">页大小</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> sshkeys</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">sshkey实体列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-2-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-2-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-2-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> fingerPrint</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">公钥指纹</span></td><td key=5></td></tr><tr key=0-0-2-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">id</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-2-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> key</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">公钥，格式：ssh-rsa AAA</span></td><td key=5></td></tr><tr key=0-0-2-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">秘钥名称</span></td><td key=5></td></tr><tr key=0-0-2-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> sshkeyId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">秘钥uuid</span></td><td key=5></td></tr><tr key=0-0-2-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-2-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-2-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">所属用户uuid</span></td><td key=5></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> totalCount</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">总条数</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr>
               </tbody>
              </table>
            
## /user/sshkeys
<a id=/user/sshkeys> </a>
### 基本信息

**Path：** /v1/user/sshkeys

**Method：** POST

**接口描述：**
CreateUserSshkey 创建个人sshkey

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> key</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">公钥内容，格式：^ssh-rsa AAAAB3NzaC1yc2.*</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">秘钥名称</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> sshkeyId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">sshkey uuid</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /user/sshkeys/instances/{sshkey_id}
<a id=/user/sshkeys/instances/{sshkey_id}> </a>
### 基本信息

**Path：** /v1/user/sshkeys/instances/{sshkey_id}

**Method：** GET

**接口描述：**
GetInstancesBySshkey 根据sshkey获取实例列表

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| sshkey_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> instanceIds</span></td><td key=1><span>string []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">实例Id列表</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>string</span></p></td></tr><tr key=array-30><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> </span></td><td key=1><span></span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /user/sshkeys/{sshkey_id}
<a id=/user/sshkeys/{sshkey_id}> </a>
### 基本信息

**Path：** /v1/user/sshkeys/{sshkey_id}

**Method：** GET

**接口描述：**
DescribeUserSshKey 获取sshkey详情

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid, deprecated |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| sshkey_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> sshkey</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0-0><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建者</span></td><td key=5></td></tr><tr key=0-0-0-1><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> createdTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">创建时间</span></td><td key=5></td></tr><tr key=0-0-0-2><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> fingerPrint</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">公钥指纹</span></td><td key=5></td></tr><tr key=0-0-0-3><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>integer</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">id</span></td><td key=5><p key=2><span style="font-weight: '700'">format: </span><span>int64</span></p></td></tr><tr key=0-0-0-4><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> key</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">公钥，格式：ssh-rsa AAA</span></td><td key=5></td></tr><tr key=0-0-0-5><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">秘钥名称</span></td><td key=5></td></tr><tr key=0-0-0-6><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> sshkeyId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">秘钥uuid</span></td><td key=5></td></tr><tr key=0-0-0-7><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedBy</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新者</span></td><td key=5></td></tr><tr key=0-0-0-8><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> updatedTime</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">更新时间</span></td><td key=5></td></tr><tr key=0-0-0-9><td key=0><span style="padding-left: 40px"><span style="color: #8c8a8a">├─</span> userId</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">所属用户uuid</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /user/sshkeys/{sshkey_id}
<a id=/user/sshkeys/{sshkey_id}> </a>
### 基本信息

**Path：** /v1/user/sshkeys/{sshkey_id}

**Method：** PUT

**接口描述：**
ModifyUserSshkey 修改sshkey

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| sshkey_id |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">秘钥名称</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## /user/sshkeys/{sshkey_id}
<a id=/user/sshkeys/{sshkey_id}> </a>
### 基本信息

**Path：** /v1/user/sshkeys/{sshkey_id}

**Method：** DELETE

**接口描述：**
DeleteUserSshkey 删除sshkey

### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
| traceId  |   | 是  |   |  流量唯一id |
| authorization  |   | 是  |   |  demo: curl -X GET -H "Authorization:Bearer wwabmrpdxz04qa10frpuw7key9qkd9bu" http://127.0.0.1:8801/v1/idcs 请使用个人自己创建apikey时候的token |
| bmpUserId  |   | 否  |   |  用户uuid |
| bmpLanguage  |   | 否  |   |  用户语言 [zh_CN, en_US] |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| sshkey_id |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> result</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> success</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">操作是否成功 [true/false]</span></td><td key=5></td></tr>
               </tbody>
              </table>
            