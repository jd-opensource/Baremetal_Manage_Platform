# Service Management in BMP Installation Guide
## Usage of `bmp-deploy.sh`
~~~
# ./bmp-deploy.sh -h

Usage:  ./bmp-deploy.sh [-t|--tag TAG] COMMAND SERVICE...

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

#### Stopping Services
~~~
# Stop the bmp-console-web service
./bmp-deploy.sh stop bmp-console-web

# Stop both bmp-console-web and bmp-operation-web services
./bmp-deploy.sh stop bmp-console-web bmp-operation-web

# Stop all services
./bmp-deploy.sh stop all
~~~

## Starting Services
~~~
# Start the bmp-console-web service
./bmp-deploy.sh start bmp-console-web

# Start both bmp-console-web and bmp-operation-web services
./bmp-deploy.sh start bmp-console-web bmp-operation-web

# Start all services
./bmp-deploy.sh start all
~~~

## Restarting Services
~~~
# Restart the bmp-console-web service
./bmp-deploy.sh restart bmp-console-web
~~~

## Checking Service Running Status
~~~
# Check the running status of all services
./bmp-deploy.sh status all
~~~

## Viewing Container Images Associated with Services
~~~
# View the container image associated with the bmp-console-web service
./bmp-deploy.sh images bmp-console-web
~~~

## Uninstalling Services
~~~
# Uninstall a single service, such as bmp-console-web
# Note: Related data will also be deleted
./bmp-deploy.sh uninstall bmp-console-web

# Uninstall all services
./bmp-deploy.sh uninstall all
~~~

## Uninstalling BMP (Uninstall All Services)
~~~
cd /root/bmp-deploy/
./bmp-deploy.sh uninstall all
cd /root
rm -rf bmp-deploy bmp-deploy*.tar.gz
rm -rf /var/log/bmp
~~~

## reinstall BMP
./bmp-deploy.sh reinstall all | [service...]
~~~
#reinstall single service eg:bmp-console-web
./bmp-deploy.sh reinstall bmp-console-web
~~~

## get all configs of BMP
./bmp-deploy.sh config all | [service...]

## deploy bmp with specific tag version
./bmp-deploy.sh start all | [service...]    -t|--tag TAG
~~~
#deploy with tag 1.0
./bmp-deploy.sh start all --tag 1.0
~~~

