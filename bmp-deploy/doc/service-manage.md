# [BMP Installation Guide](main.md) - Service Management
bmp-deploy.sh Usage Instructions
~~~
#./bmp-deploy.sh -h

Usage: ./bmp-deploy.sh [-t|--tag TAG] COMMAND SERVICE...

COMMAND:
    start       Install and start SERVICE
    stop        Stop SERVICE
    restart     Restart SERVICE
    status      Check SERVICE status
    images      Display SERVICE image
    uninstall   Uninstall SERVICE
    reinstall   Reinstall SERVICE
    config      Get SERVICE config

SERVICE:
    all                 All SERVICES
    bmp-db
    bmp-redis
    bmp-mq
    bmp-image
    bmp-pushgateway
    bmp-alertmanager
    bmp-prometheus
    bmp-tftp
    bmp-rsyslog
    bmp-console-web
    bmp-operation-web
    bmp-console-api
    bmp-operation-api
    bmp-openapi
    bmp-openapi-console
    bmp-scheduler
    bmp-driver
    bmp-dhcp-agent
    bmp-oob-alert
    bmp-oob-agent
    bmp-monitor-proxy
    bmp-pronoea
~~~
## Stop Services
./bmp-deploy.sh stop all | [service...]
~~~
# Stop the bmp-console-web service
./bmp-deploy.sh stop bmp-console-web

# Stop the bmp-console-web and bmp-operation-web services
./bmp-deploy.sh stop bmp-console-web bmp-operation-web

# Stop all services
./bmp-deploy.sh stop all
~~~

## Start Services
./bmp-deploy.sh start all | [service...]
~~~
# Start the bmp-console-web service
./bmp-deploy.sh start bmp-console-web

# Start the bmp-console-web and bmp-operation-web services
./bmp-deploy.sh start bmp-console-web bmp-operation-web

# Start all services
./bmp-deploy.sh start all
~~~

## Restart Services
./bmp-deploy.sh restart all | [service...]
~~~
# Restart the bmp-console-web service
./bmp-deploy.sh restart bmp-console-web
~~~

## Check Service Status
./bmp-deploy.sh status all | [service...]
~~~
# Check the status of all services
./bmp-deploy.sh status all
~~~

## Check Service-Related Container Images
./bmp-deploy.sh images all | [service...]
~~~
# Check the container image related to the bmp-console-web service
./bmp-deploy.sh images bmp-console-web
~~~

## Uninstall Services
./bmp-deploy.sh uninstall all | [service...]

Note: Related data will also be deleted
~~~
# Uninstall a single service, such as bmp-console-web
./bmp-deploy.sh uninstall bmp-console-web

# Uninstall all services
./bmp-deploy.sh uninstall all
~~~

## Uninstall BMP (Uninstall All Services)
~~~
cd /root/bmp-deploy/
./bmp-deploy.sh uninstall all
cd /root
rm -rf bmp-deploy bmp-deploy*.tar.gz
rm -rf /var/log/bmp
~~~

## Reinstall Services (Uninstall + Install)
./bmp-deploy.sh reinstall all | [service...]
~~~
# Reinstall a single service, such as bmp-console-web
./bmp-deploy.sh reinstall bmp-console-web
~~~

## Get Service Configuration
./bmp-deploy.sh config all | [service...]

## Deploy BMP with a Specific Version
./bmp-deploy.sh start all | [service...]    -t|--tag TAG
~~~
# Deploy version 1.0
./bmp-deploy.sh start all --tag 1.0
~~~

### Previous section [Deploying BMP](deploy.md)
### Next section [BM Node Deployment Process](bm-deploy.md)