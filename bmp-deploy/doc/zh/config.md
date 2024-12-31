#  [bmp安装指导](main.md) - 配置管理
.env文件包含bmp各种配置参数
~~~
BMP_HOST_IP=xxx.xxx.xxx.xxx
BMP_VERSION=dev
COMPOSE_PROJECT_NAME=bmp
BMP_DB_HOST=bmp-db
BMP_DB_PORT=3306
BMP_DB_USER=bmp_rw
BMP_DB_PASSWORD='LpK9Jq12Zf'
BMP_DB_NAME=bmp
BMP_REDIS_HOST=bmp-redis
BMP_REDIS_PORT=6379
BMP_REDIS_PASSWORD='LpK9Jq12Zf'
BMP_MQ_HOST=bmp-mq
BMP_MQ_PORT=5672
BMP_MQ_USER=bmp
BMP_MQ_PASSWORD='LpK9Jq12Zf'
BMP_MQ_VHOST=/bmp
BMP_OMAPI_HOST=bmp-dhcp-agent
BMP_OMAPI_PORT=7911
BMP_OMAPI_TOKEN="LpK9Jq12Zf"
BMP_OPENAPI_HOST=bmp-openapi
BMP_OPENAPI_PORT=8801
BMP_OPENAPI_CONSOLE_HOST=bmp-openapi-console
BMP_OPENAPI_CONSOLE_PORT=8802
BMP_OOB_ALERT_HOST=bmp-oob-alert
BMP_OOB_ALERT_PORT=8804
BMP_MQ_EXCHANGE_ROUTING_KEY=idc-vm4xsulx1k2d9z4xkctrttig02zl
BMP_IMAGE_PORT=10000
BMP_MONITOR_PROXY_PORT=8805
BMP_PROMETHEUS_HOST=bmp-prometheus
BMP_PROMETHEUS_PORT=9090
BMP_PROMETHEUS_DATA_DIR=/data/prometheus
BMP_PUSHGATEWAY_HOST=bmp-pushgateway
BMP_PUSHGATEWAY_PORT=9091
BMP_ALERTMANAGER_HOST=bmp-alertmanager
BMP_ALERTMANAGER_PORT=9093
BMP_PRONOEA_HOST=bmp-pronoea
BMP_PRONOEA_PORT=9999
~~~

## 配置说明：
| 配置名                         | 说明                                            |
|-----------------------------|-----------------------------------------------|
| BMP_HOST_IP                 | manager节点ip, 请配置为管理网卡ip，bmp各服务组件监听此ip         |
| BMP_VERSION                 | 待安装的bmp版本，请勿随便修改                              |
| COMPOSE_PROJECT_NAME        | docker compose 项目名，默认值是bmp，请勿修改               |
| BMP_DB_HOST                 | bmp-db ip/域名，请勿修改                             |
| BMP_DB_PORT                 | bmp-db 端口，请勿修改                                |
| BMP_DB_USER                 | bmp-db 用户名                                    |
| BMP_DB_PASSWORD             | bmp-db 用户密码                                   |
| BMP_DB_NAME                 | bmp-db 数据库名，请勿修改                              |
| BMP_REDIS_HOST              | bmp-redis ip/域名，请勿修改                          |
| BMP_REDIS_PORT              | bmp-redis 端口，请勿修改                             |
| BMP_REDIS_PASSWORD          | bmp-redis 密码                                  |
| BMP_MQ_HOST                 | bmp-mq ip/域名，请勿修改                             |
| BMP_MQ_PORT                 | bmp-mq 端口，请勿修改                                |
| BMP_MQ_USER                 | bmp-mq 用户名                                    |
| BMP_MQ_PASSWORD             | bmp-mq 密码                                     |
| BMP_MQ_VHOST                | bmp-mq vhost名，请勿修改                            |
| BMP_OMAPI_HOST              | dhcp omapi服务 ip/域名，请勿修改                       |
| BMP_OMAPI_PORT              | dhcp omapi服务端口，请勿修改                           |
| BMP_OMAPI_TOKEN             | dhcp omapi认证token                             |
| BMP_OPENAPI_HOST            | bmp-openapi ip/域名，请勿修改                        |
| BMP_OPENAPI_PORT            | bmp-openapi 端口，请勿修改                           |
| BMP_OPENAPI_CONSOLE_HOST    | bmp-openapi-console ip/域名，请勿修改                |
| BMP_OPENAPI_CONSOLE_PORT    | bmp-openapi-console 端口，请勿修改                   |
| BMP_OOB_ALERT_HOST          | bmp-oob-alert ip/域名，请勿修改                      |
| BMP_OOB_ALERT_PORT          | bmp-oob-alert 端口，请勿修改                         |
| BMP_MQ_EXCHANGE_ROUTING_KEY | mq routing key, 此值应该等于idc_id，请勿修改             |
| BMP_IMAGE_PORT              | bmp 镜像服务器端口, 请勿修改                             |
| BMP_MONITOR_PROXY_PORT      | bmp-monitor-proxy 端口，请勿修改                     |
| BMP_PROMETHEUS_HOST         | bmp_prometheus ip/域名，请勿修改                     |
| BMP_PROMETHEUS_PORT         | bmp_prometheus 端口，请勿修改                        |
| BMP_PROMETHEUS_DATA_DIR     | bmp_prometheus 数据存储路径，此路径需要的存储空间较大，建议根据实际情况修改 |
| BMP_PUSHGATEWAY_HOST        | bmp_pushgateway ip/域名，请勿修改                    |
| BMP_PUSHGATEWAY_PORT        | bmp_pushgateway 端口，请勿修改                       |
| BMP_ALERTMANAGER_HOST       | bmp_alertmanager ip/域名，请勿修改                   |
| BMP_ALERTMANAGER_PORT       | bmp_alertmanager 端口，请勿修改                      |
| BMP_PRONOEA_HOST            | bmp_pronoea ip/域名，请勿修改                        |
| BMP_PRONOEA_PORT            | bmp_pronoea 端口，请勿修改                           |

## 注意
* BMP_DB_PASSWORD, BMP_REDIS_PASSWORD, BMP_MQ_PASSWORD, BMP_OMAPI_TOKEN的值只能由 **大小字母和数字** 组成，包含特殊字符可能会导致错误。

# 上一节 [实用工具](tool.md)

