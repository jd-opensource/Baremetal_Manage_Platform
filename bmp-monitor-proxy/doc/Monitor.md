# CPS监控接入——接口定义

Table of Contents
=================
* [指标上报](#指标上报)
* [查询指标](#查询指标)


变更日期 | 变更内容
---- | ----
2019-05-09 | 初始化
2019-05-16 | 接口调整，按照运维平台监控适配
2019-05-21 | 接口调整，根据前端组件，优化接口格式
2019-05-23 | 接口调整，去掉Tags，转移到MetricName中


## 指标上报

### url

```http
/put
```

### method

```http
POST
```

### 请求参数

参数名称 | 类型 | 说明
---- | ---- | -----
source | String | 必传参数，云物理机：cps
region | String | 必选参数，如：cn-north-1
instanceId | String | 必选参数，实例ID
sn | String | 可选参数，云物理机必选参数
dataPoints | List\<DataPoint> | 必选参数，监控数据

**DataPoint**

参数名称 | 类型 | 说明
---- | ---- | -----
metric | String | 写入的指标名称
timestamp | Long | 秒级时间戳。建议的上报时间戳：上报时间间隔的整数倍，如上报间隔为5s，则建议上报的时间戳为 timestamp = current timestamp - (current timestamp % time interval) = 1487647187 - （1487647187 % 5） = 1487647187 -2 = 1487647185
valueType | String | int，float
value | Int,float | 监控的值。最大值为long.MAX_VALUE=9223372036854775807=263-1。累计类型的指标，累计到最大值后需要翻转为0，重新开始计数。翻转后不影响速率的计算。
tags | Map<String,String> | 标签，0个或多个。最多5个。

#### Request

```json
{
    "source": "cps",
    "region": "cn-north-1",
    "instanceId": "cps-97fzedyq3fhk7dilhgqgtmiservg",
    "sn": "J33WRF1",
    "dataPoints": [
        {
            "metric": "cps.cpu.util",
            "timestamp": 1558060300,
            "valueType": "float",
            "value": 5.51,
            "tags": {}
        },
        {
            "metric": "cps.disk.used",
            "timestamp": 1558060300,
            "valueType": "int",
            "value": 3221225472,
            "tags": {
                "device": "sda"
            }
        },
        {
            "metric": "cps.disk.used",
            "timestamp": 1558060300,
            "valueType": "int",
            "value": 64424509440,
            "tags": {
                "device": "sdb"
            }
        }
    ]
}
```

### 响应参数


参数名称 | 类型 | 说明
---- | ---- | -----
requestId | String | 请求ID
result | Result | 结果

**Result**

参数名称 | 类型 | 说明
---- | ---- | -----
code | Int | 结果码
message | String | 结果描述

#### Response

```json
{
    "requestId": "bjehoueq3tiu7c5ie6bha2pufsw16a7d",
    "result": {
        "code": 0,
        "message": "success"
    }
}
```

## 查询接口

### url

```http
/describeMetricData
```

### method

```http
POST
```

### 请求参数


参数名称 | 类型 | 说明
---- | ---- | ----
source | String | 云物理机：cps
region | String | 地域：可用地域列表。注：该字段将覆盖dataCenter字段
instanceId | String | 实例ID
metrics | List<String> | 希望获取的指标列表，如为空默认为全部
startTime | String | 查询时间范围的开始时间， UTC时间，格式：2016-12-11T00:00:00+0800（早于30d时，将被重置为30d）（注意在url中+要转译为%2B故url中为2016-12-11T00:00:00%2B0800）
endTime | String | 查询时间范围的结束时间， UTC时间，格式：2016-12-11T00:00:00+0800（为空时，将由startTime与timeInterval计算得出）（注意在url中+要转译为%2B故url中为2016-12-11T00:00:00%2B0800）
timeInterval | String | 时间间隔：1h，6h，12h，1d，3d，7d，14d固定时间间隔使用固定的聚合粒度


#### Request

**固定时间间隔**
```json
{
    "source": "cps",
    "region": "cn-north-1",
    "instanceId": "cps-97fzedyq3fhk7dilhgqgtmiservg",
    "metrics": [
        "cps.disk.bytes.read",
        "cps.disk.bytes.write"
    ],
    "timeInterval": "1h"
}
```
**自定义时间间隔**
```json
{
    "source": "cps",
    "region": "cn-north-1",
    "instanceId": "cps-97fzedyq3fhk7dilhgqgtmiservg",
    "metrics": [
        "cps.disk.bytes.read",
        "cps.disk.bytes.write"
    ],
    "startTime": "2019-05-01T00:00:00+0800",
    "endTime": "2019-05-18T00:00:00+0800"
}
```

### 响应参数

参数名称 | 类型 | 说明
---- | ---- | ----
requestId | String | 请求ID
result | Map<String, List\<MetricData> | 结果

**MetricData**

参数名称 | 类型 | 说明
---- | ---- | ----
metric | Metric | Metric信息
data | List<Dps> | 监控数据组成的列表

**Metric**

参数名称 | 类型 | 说明
---- | ---- | ----
aggregator | String | 聚合类型：last、avg
calculateUnit | String | 单位
metric | String | metric名称
metricName | String | tag+metric显示名称
period | String | 聚合周期

**Dps**

参数名称 | 类型 | 说明
---- | ---- | ----
timestamp | Long | 豪秒级时间戳。
value | float | 监控的值。


#### Response

**无tags的指标**
```json
{
    "requestId": "bjehoueq3tiu7c5ie6bha2pufsw16a7d",
    "result": {
        "metricDatas": [
            {
                "data": [
                    {
                        "timestamp": 1557990900000,
                        "value": "0.0000"
                    },
                    {
                        "timestamp": 1557990960000,
                        "value": "0.0000"
                    },
                    {
                        "timestamp": 1557991020000,
                        "value": "0.0000"
                    },
                    …………
                ],
                "metric": {
                    "calculateUnit": "%",
                    "metric": "cps.cpu.util",
                    "metricName": "CPU Usage",
                    "aggregator": "avg",
                    "period": "1min"
                }
            }
        ]
    }
}

```

**带tags的指标**
```json
{
    "requestId": "bjehoueq3tiu7c5ie6bha2pufsw16a7d",
    "result": {
        "metricDatas": [
            {
                "data": [
                    {
                        "timestamp": 1557990900000,
                        "value": "0.0000"
                    },
                    {
                        "timestamp": 1557990960000,
                        "value": "0.0000"
                    },
                    {
                        "timestamp": 1557991020000,
                        "value": "0.0000"
                    },
					…………
                ],
                "tags": [
                    {
                        "tagKey": "device",
                        "tagValue": "sda"
                    }
                ],
                "metric": {
                    "calculateUnit": "Bps",
                    "metric": "cps.disk.bytes.read",
                    "metricName": "sda Disk Read Traffic",
                    "aggregator": "last",
                    "period": "1min"
                }
            },
            {
                "data": [
                    {
                        "timestamp": 1557990900000,
                        "value": "4505.6000"
                    },
                    {
                        "timestamp": 1557990960000,
                        "value": "10513.0667"
                    },
                    {
                        "timestamp": 1557991020000,
                        "value": "4642.1333"
                    },
					…………
                ],
                "metric": {
                    "calculateUnit": "Bps",
                    "metric": "cps.disk.bytes.read",
                    "metricName": "sdb Disk Read Traffic",
                    "aggregator": "last",
                    "period": "1min"
                }
            },
            {
                "data": [
                    {
                        "timestamp": 1557990900000,
                        "value": "0.0000"
                    },
                    {
                        "timestamp": 1557990960000,
                        "value": "0.0000"
                    },
                    {
                        "timestamp": 1557991020000,
                        "value": "0.0000"
                    },
					…………
                ], 
                "metric": {
                    "calculateUnit": "Bps",
                    "metric": "cps.disk.bytes.write",
                    "metricName": "sda Disk Write Traffic",
                    "aggregator": "last",
                    "period": "1min"
                }
            },
            {
                "data": [
                    {
                        "timestamp": 1557990900000,
                        "value": "4505.6000"
                    },
                    {
                        "timestamp": 1557990960000,
                        "value": "10513.0667"
                    },
                    {
                        "timestamp": 1557991020000,
                        "value": "4642.1333"
                    },
					…………
                ], 
                "metric": {
                    "calculateUnit": "Bps",
                    "metric": "cps.disk.bytes.write",
                    "metricName": "sdb Disk Write Traffic",
                    "aggregator": "last",
                    "period": "1min"
                }
            }
        ]
    }
}
```
