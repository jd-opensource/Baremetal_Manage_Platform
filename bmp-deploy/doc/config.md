# [BMP Installation Guide](main.md) - Configuration Management
The `.env` file contains various configuration parameters for BMP.
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

## Configuration Explanation:
| Configuration Name                         | Description                                            |
|-------------------------------------------|--------------------------------------------------------|
| BMP_HOST_IP                 | Manager node IP, please configure it as the management network card IP, BMP various service components listen on this IP         |
| BMP_VERSION                 | The version of BMP to be installed, please do not modify it randomly                              |
| COMPOSE_PROJECT_NAME        | Docker Compose project name, default value is bmp, please do not modify               |
| BMP_DB_HOST                 | BMP-DB IP/hostname, please do not modify                             |
| BMP_DB_PORT                 | BMP-DB port, please do not modify                                |
| BMP_DB_USER                 | BMP-DB username                                    |
| BMP_DB_PASSWORD             | BMP-DB user password                                   |
| BMP_DB_NAME                 | BMP-DB database name, please do not modify                              |
| BMP_REDIS_HOST              | BMP-Redis IP/hostname, please do not modify                          |
| BMP_REDIS_PORT              | BMP-Redis port, please do not modify                             |
| BMP_REDIS_PASSWORD          | BMP-Redis password                                  |
| BMP_MQ_HOST                 | BMP-MQ IP/hostname, please do not modify                             |
| BMP_MQ_PORT                 | BMP-MQ port, please do not modify                                |
| BMP_MQ_USER                 | BMP-MQ username                                    |
| BMP_MQ_PASSWORD             | BMP-MQ password                                     |
| BMP_MQ_VHOST                | BMP-MQ vhost name, please do not modify                            |
| BMP_OMAPI_HOST              | DHCP OMAPI service IP/hostname, please do not modify                       |
| BMP_OMAPI_PORT              | DHCP OMAPI service port, please do not modify                           |
| BMP_OMAPI_TOKEN             | DHCP OMAPI authentication token                             |
| BMP_OPENAPI_HOST            | BMP-OpenAPI IP/hostname, please do not modify                        |
| BMP_OPENAPI_PORT            | BMP-OpenAPI port, please do not modify                           |
| BMP_OPENAPI_CONSOLE_HOST    | BMP-OpenAPI-Console IP/hostname, please do not modify                |
| BMP_OPENAPI_CONSOLE_PORT    | BMP-OpenAPI-Console port, please do not modify                   |
| BMP_OOB_ALERT_HOST          | BMP-OOB-Alert IP/hostname, please do not modify                      |
| BMP_OOB_ALERT_PORT          | BMP-OOB-Alert port, please do not modify                         |
| BMP_MQ_EXCHANGE_ROUTING_KEY | MQ routing key, this value should be equal to idc_id, please do not modify             |
| BMP_IMAGE_PORT              | BMP image server port, please do not modify                             |
| BMP_MONITOR_PROXY_PORT      | BMP-monitor-proxy port, please do not modify                     |
| BMP_PROMETHEUS_HOST         | BMP-Prometheus IP/hostname, please do not modify                     |
| BMP_PROMETHEUS_PORT         | BMP-Prometheus port, please do not modify                        |
| BMP_PROMETHEUS_DATA_DIR     | BMP-Prometheus data storage path, this path requires a large amount of storage space, please modify according to actual needs |
| BMP_PUSHGATEWAY_HOST        | BMP-Pushgateway IP/hostname, please do not modify                    |
| BMP_PUSHGATEWAY_PORT        | BMP-Pushgateway port, please do not modify                       |
| BMP_ALERTMANAGER_HOST       | BMP-Alertmanager IP/hostname, please do not modify                   |
| BMP_ALERTMANAGER_PORT       | BMP-Alertmanager port, please do not modify                      |
| BMP_PRONOEA_HOST            | BMP-Pronoea IP/hostname, please do not modify                        |
| BMP_PRONOEA_PORT            | BMP-Pronoea port, please do not modify                           |

## Note
* The values of BMP_DB_PASSWORD, BMP_REDIS_PASSWORD, BMP_MQ_PASSWORD, and BMP_OMAPI_TOKEN can only consist of **uppercase and lowercase letters and digits**. Including special characters may cause errors.