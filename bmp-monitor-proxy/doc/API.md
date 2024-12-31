
### 查询Agent状态
> schema
> 
>     HTTP
>  
> url
> 
>     /api/v1/describeAgentStatus
>
> method
>
>     POST
>
> request
>
>     instance: 实例ID, 多个ID以逗号分隔
> 
> response
> 
> 成功: 
> ```json
> {
>    "code": 0,
>    "message": "Success",
>    "data": [
>       {
>          "instance_id": "cps-h1mtnemztijtpgbmjfexpj0dovm4",
>          "sn": "xxxx",
>          "version": "v1.0.0-xxx",
>          "timestamp": 1553151660,
>          "status": 0
>       },
>       {
>          "instance_id": "cps-h1mtnemztijtpgbmjfexpj0dovm3",
>          "sn": "",
>          "version": "",
>          "timestamp": 0,
>          "status": 2
>       }
>       ......
>     ]
> }
> ```
> 失败：
> ```json
> {
>    "code": 1,
>    "message": "No Parameter",
>    "data": null
> }
> ```


### 代理上报接口 Put
> schema
> 
>     Mutual TLS
>  
> url
> 
>     /api/v1/put
>
> method
>
>     POST
>
> request
> ```json
> {
>     "sn": "PF11ZXQG",
>     "dataPoints": [
>         {
>             "metric": "cps.network.packets.egress",
>             "timestamp": 1553151660,
>             "value": 1281955,
>             "tags": {}
>         },
>         {
>             "metric": "cps.process.total",
>             "timestamp": 1553151660,
>             "value": 319,
>             "tags": {}
>         },
>         ......
>     ]
> }
> ```
> 
> response
> 
> 成功: 
> ```json
> {
>    "code": 0,
>    "message": "Success",
>    "data": {}
> }
> ```
> 失败：
> ```json
> {
>    "code": 1,
>    "message": "No Parameter",
>    "data": {}
> }
> ```
>


### Agent心跳上报 Heartbeat
> schema
> 
>     Mutual TLS
>  
> url
> 
>     /api/v1/heartbeat
>
> method
>
>     POST
>
> request
> ```json
> {
>     "sn": "PF11ZXQG",
>     "agent_version": "0.0.1-v0.0.2-5-g202dea7-dirty-20190321145703",
>     "timestamp": 1553151692,
>     "status": 0
> }
> ```
> 
> response
> 
> 成功: 
> ```json
> {
>    "code": 0,
>    "message": "Success",
>    "data": {}
> }
> ```
> 失败：
> ```json
> {
>    "code": 1,
>    "message": "No Parameter",
>    "data": {}
> }
> ```
>

### 边缘节点代理上报接口 ProxyPut (边缘节点Proxy调用，暂未使用)
> schema
> 
>     Mutual TLS
>  
> url
> 
>     /api/v1/proxyPut
>
> method
>
>     POST
>
> request
> ```json
> {
>     "appCode": "jcloud",
>     "serviceCode": "cps",
>     "dataCenter": null,
>     "resourceId": "cps-h1mtnemztijtpgbmjfexpj0dovm4",
>     "region": "cn-north-1",
>     "dataPoints": [
>         {
>             "metric": "cps.network.packets.egress",
>             "timestamp": 1553151660,
>             "value": 1281955,
>             "tags": {}
>         },
>         {
>             "metric": "cps.process.total",
>             "timestamp": 1553151660,
>             "value": 319,
>             "tags": {}
>         },
>         ......
>     ]
> }
> ```
> 
> response
> 
> 成功: 
> ```json
> {
>    "code": 0,
>    "message": "Success",
>    "data": {}
> }
> ```
> 失败：
> ```json
> {
>    "code": 1,
>    "message": "No Parameter",
>    "data": {}
> }
> ```


### response code
> 
> |Code|含义|
> |-----|----|
> |0|成功|
> |1|缺少参数|
> |2|参数错误|
> |3|无结果|
> |4|远程接口错误|
> |99|未知错误|