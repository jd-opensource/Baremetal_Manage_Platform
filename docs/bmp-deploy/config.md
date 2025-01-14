#  [bmp安装指导](main.md) - 配置管理
.env file configuration
~~~
BMP_HOST_IP=xxx.xxx.xxx.xxx
# bmp registry
BMP_REGISTRY=docker.io

# db middleware registry eg: mysql:5.7-debian
# if you cannot pull from docker.io please change docker.io to other Container image repository
BASE_REGISTRY=docker.m.daocloud.io

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

## Configuration Description:
| Configuration Name | Description |
|---|---|
| BMP_HOST_IP | IP of the manager node, please configure it to the management network card IP, BMP service components listen on this IP. |
| BMP_VERSION | Version of BMP to be installed, do not modify casually. |
| COMPOSE_PROJECT_NAME | Docker Compose project name, default value is `bmp`, do not modify. |
| BMP_DB_HOST | IP/hostname of bmp-db, do not modify. |
| BMP_DB_PORT | Port of bmp-db, do not modify. |
| BMP_DB_USER | Username for bmp-db. |
| BMP_DB_PASSWORD | Password for bmp-db. |
| BMP_DB_NAME | Database name for bmp-db, do not modify. |
| BMP_REDIS_HOST | IP/hostname of bmp-redis, do not modify. |
| BMP_REDIS_PORT | Port of bmp-redis, do not modify. |
| BMP_REDIS_PASSWORD | Password for bmp-redis. |
| BMP_MQ_HOST | IP/hostname of bmp-mq, do not modify. |
| BMP_MQ_PORT | Port of bmp-mq, do not modify. |
| BMP_MQ_USER | Username for bmp-mq. |
| BMP_MQ_PASSWORD | Password for bmp-mq. |
| BMP_MQ_VHOST | Vhost name for bmp-mq, do not modify. |
| BMP_OMAPI_HOST | IP/hostname of DHCP OMAPI service, do not modify. |
| BMP_OMAPI_PORT | Port of DHCP OMAPI service, do not modify. |
| BMP_OMAPI_TOKEN | Authentication token for DHCP OMAPI. |
| BMP_OPENAPI_HOST | IP/hostname of bmp-openapi, do not modify. |
| BMP_OPENAPI_PORT | Port of bmp-openapi, do not modify. |
| BMP_OPENAPI_CONSOLE_HOST | IP/hostname of bmp-openapi-console, do not modify. |
| BMP_OPENAPI_CONSOLE_PORT | Port of bmp-openapi-console, do not modify. |
| BMP_OOB_ALERT_HOST | IP/hostname of bmp-oob-alert, do not modify. |
| BMP_OOB_ALERT_PORT | Port of bmp-oob-alert, do not modify. |
| BMP_MQ_EXCHANGE_ROUTING_KEY | MQ routing key, this value should be equal to `idc_id`, do not modify. |
| BMP_IMAGE_PORT | Port of BMP image server, do not modify. |
| BMP_MONITOR_PROXY_PORT | Port of bmp-monitor-proxy, do not modify. |
| BMP_PROMETHEUS_HOST | IP/hostname of bmp_prometheus, do not modify. |
| BMP_PROMETHEUS_PORT | Port of bmp_prometheus, do not modify. |
| BMP_PROMETHEUS_DATA_DIR | Data storage path for bmp_prometheus, this path requires a large amount of storage space, it is recommended to modify according to actual conditions. |
| BMP_PUSHGATEWAY_HOST | IP/hostname of bmp_pushgateway, do not modify. |
| BMP_PUSHGATEWAY_PORT | Port of bmp_pushgateway, do not modify. |
| BMP_ALERTMANAGER_HOST | IP/hostname of bmp_alertmanager, do not modify. |
| BMP_ALERTMANAGER_PORT | Port of bmp_alertmanager, do not modify. |
| BMP_PRONOEA_HOST | IP/hostname of bmp_pronoea, do not modify. |
| BMP_PRONOEA_PORT | Port of bmp_pronoea, do not modify. |

## Notes
* The values of `BMP_DB_PASSWORD`, `BMP_REDIS_PASSWORD`, `BMP_MQ_PASSWORD`, and `BMP_OMAPI_TOKEN` can only consist of **uppercase and lowercase letters and numbers**. Special characters may cause errors.




