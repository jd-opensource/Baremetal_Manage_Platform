# pronoea接口文档
* [ChangeLog](#ChangeLog)

* [环境](#环境)

* [HTTP响应体定义](#HTTP响应体定义)

## ChangeLog

| 版本   | 变更内容   | 变更时间   | 变更人                  |
| ------ | ---------- | ---------- | ----------------------- |
| V1.0.0 | 初始化版本 | 2024/06/20 | 马赛迪(masaidi1@jd.com) |



## 环境


*  测试环境

```
 
```

*  预发环境

```
10.226.192.114:9999
```

*  正式环境

```
用户局域网，无正式环境
```

## HTTP响应体定义

```
{
    "Code": xxxxx,           //错误码 10000 表示成功，其他为失败
    "RequestId": "xxxxxx",   //请求id
    "Message": "success",    //错误信息
    "Result": {}
}
```

* [样本数据相关接口](#样本数据相关接口)

    * [时间范围查询样本数据](#时间范围查询样本数据)

* [告警规则接口](#告警规则接口)

    * [写告警规则接口](#写告警规则接口)
    * [删除告警规则接口](#删除告警规则接口)

## 样本数据相关接口

### 时间范围查询样本数据
```bash
POST /api/query
```

#### 请求参数
| 参数名             | 类型   | 是否必须 | 默认值 | 描述      |
| ------------------ | ------ | -------- | ------ | --------- |
| metricName | String | 是       | ---    | 样本名称，目前只有bmp_monitor |
| labels | map[string]string | 否       | ---    | 标签查询条件 |
| startTime | int64 | 是       | ---    | 11位起始时间戳 |
| endTime | int64 | 是       | ---    | 11位结束时间戳 |
| step | int | 是       | ---    | 步长 |      

#### 响应示例

```json
{
  "Code": 10000,
  "RequestId": "5badf577-ed9c-47c6-ad48-b201cd2159e0",
  "Message": "success",
  "Result": {
    "data": [
      {
          "timestamp": 1718313728,     // 时间点
          "value": "35774971904"       // 时间点对应的值
      }
    ],
    "query": {                         //接收到的查询条件
      "metricName": "bmp_monitor",
      "labels": null,
      "startTime": 1718313728,
      "endTime": 1718313788,
      "step": 60
    }
  }
}
```

## 告警规则接口

### 写告警规则接口
```bash
POST /api/rules/write
```

#### 请求参数

#### 带内监控告警规则入参
| 参数名             | 类型   | 是否必须 | 默认值 | 描述      |
| ------------------ | ------ | -------- | ------ | --------- |
| ruleId | String | 是       | ---    | 规则id |
| source | String | 是       | ---    | 规则来源|
| dimension | String | 否      | ---    | 维度 |
| dimensionName | String | 否       | ---    | 维度名称 |
| resource | String | 否       | ---    | 资源类型 | 
| resourceName | String | 否       | ---    | 资源类型名称 | 
| triggerOption | []*RuleTrigger | 是       | ---    | 触发条件 | 
| noticeOption | *RuleNotice | 是       | ---    | 通知策略 |
| instanceIds | []String | 是       | ---    | 实例ID列表 |
| deviceTag| String | 否       | ---    | 盘符、挂载点、网口 | 

##### RuleTrigger结构定义
| 参数名             | 类型   | 是否必须 | 默认值 | 描述      |
| ------------------ | ------ | -------- | ------ | --------- |
| metric | String | 是       | ---    | 监控指标 |
| metricName | String | 是       | ---    | 监控指标名称|
| period | int | 是      | ---    | 周期 分钟位单位 |
| calculation | String | 是       | ---    | 计算方式[min,max,avg,sum] |
| calculationUnit | String | 否       | ---    | 单位 | 
| operation | String | 是       | ---    | 比较方式  [> >= < <= == !=]或者[gt gte lt lte eq neq] | 
| threshold | Float64 | 是       | ---    | 阈值 | 
| times | int | 是       | ---    | 持续周期数  [1, 2, 3, 5, 10, 15, 30, 60] |
| instanceIds | []String | 是       | ---    | 实例ID列表 |
| noticeLevel| int | 是      | ---    | 告警级别 [1表示一般，2表示严重，3表示紧急] | 

##### RuleNotice结构定义
| 参数名             | 类型   | 是否必须 | 默认值 | 描述      |
| ------------------ | ------ | -------- | ------ | --------- |
| noticePeriod | String | 是       | ---    | 通知周期(分钟) [5 10 15 30 60 180 360 720 1440] |

#### 响应示例

```json
{
  "Code": 10000,
  "RequestId": "341ae5b5-a55b-4e88-8e5d-81adfade81c2",
  "Message": "success",
  "Result": null
}
```

### 删除告警规则接口
```bash
POST /api/rules/delete
```

#### 请求参数

#### 带内监控告警规则入参
| 参数名             | 类型   | 是否必须 | 默认值 | 描述      |
| ------------------ | ------ | -------- | ------ | --------- |
| ruleId | String | 是       | ---    | 规则id |
| source | String | 是       | ---    | 规则来源|

#### 响应示例

```json
{
  "Code": 10000,
  "RequestId": "01343512-c291-48f1-82f3-c6589edb6028",
  "Message": "success",
  "Result": {
    "ruleId": "19f449d9-0574-4b69-8878-64a2e380c3ab",
    "source": "bmpInBand"
  }
}
```
