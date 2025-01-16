# Local Development and Deployment
This chapter covers:
- How to deploy and debug individual components locally (suitable for local development and debugging)
- How to deploy and integrate components (suitable for non-containerized deployment of the BMP platform)

## Development and Deployment Steps
- [Component Introduction](#1)
- [Single Component Deployment and Debugging](#2)
    - [Code Introduction and Port Planning](#2.1)
        - [Code Introduction](#2.1.1)
        - [Setting Environment Variables](#2.1.2)
    - [Base Components](#2.2)
        - [bmp-image](#2.2.1)
        - [bmp-tftp](#2.2.2)
        - [bmp-db](#2.2.3)
        - [bmp-redis](#2.2.4)
        - [bmp-mq](#2.2.5)
        - [bmp-rsyslog](#2.2.6)
    - [Frontend Components](#2.3)
    - [Backend Components](#2.4)
        - [Backend Go Components](#2.4.1)
        - [Backend Python Components](#2.4.2)
    - [Open-Source Components](#2.5)
        - [bmp-pushgateway](#2.5.1)
        - [bmp-prometheus](#2.5.2)
        - [bmp-alertmanager](#2.5.3)
- [Component Integration Deployment](#3)
    - [Environment Preparation](#3.1)
    - [Setting Environment Variables](#3.2)
    - [Starting Components in Order](#3.3)
    - [Accessing BMP](#3.4)
    - [Integrating Bare Metal Servers with the Management Platform](#3.5)


## Component Introduction<a id="1"></a>
* bmp-console-web: Control panel frontend page. Built using the vue3 scaffold. Includes login page, project management page, personal center page, instance management page, etc.
* bmp-console-api: Control panel API. Go language backend business, calls openapi interface to implement control panel related interfaces. Assembles various required data for bmp-console-web, requires multi-language adaptation.
* bmp-operation-web: Operation platform frontend page. Built using the vue3 scaffold, includes login page, data center management page, machine type management page, image management page, device management page, role management page, user management page, etc.
* bmp-operation-api: Operation platform API. Go language backend business, calls openapi interface to implement operation platform related interfaces. Assembles various required data for bmp-operation-web, requires multi-language adaptation.
* bmp-openapi: bmp-openapi is the core module of BMP. It implements restful API format interfaces that meet the swagger2.0 specification. Provides all basic functions of BMP. Performs database operations internally and calls bmp-scheduler to complete instance lifecycle management related operations.
* bmp-scheduler: Provisioning scheduler module. Accepts instance lifecycle management requests from bmp-openapi, converts upper-layer requests into corresponding commands, and drives commands to execute, coordinating with lower-level bmp-driver and bmp-agent to complete tasks such as installation, reinstallation, power on, and power off.
* bmp-driver: Single data center application, needs to deploy multiple sets of bmp-driver services in the case of multiple data centers, receives mq, and performs operations such as power on, power off, restart, and setting PXE boot for servers in the data center.
* bmp-dhcp-agent: Single data center application, needs to update DHCP configuration before installation, stores the MAC-IP association relationship in the DHCP configuration. LiveOS can then obtain the IP address from DHCP.
* bmp-db: Database
* bmp-redis: Redis cache
* bmp-mq: Message middleware
* bmp-tftp: TFTP server, stores files required for PXE boot, including PXE boot program, PXE boot configuration, LiveOS kernel, and initramfs.
* bmp-image: HTTP server, stores GuestOS images
* bmp-rsyslog: rsyslog log component
* bmp-oob-alert: Out-of-band alert component
* bmp-oob-agent: Out-of-band monitoring information collection component
* bmp-monitor-proxy: In-band monitoring forwarding component
* bmp-prometheus: Monitoring data collection component
* bmp-pushgateway: Collects monitoring data from bmp-monitor-proxy and pushes it to Prometheus
* bmp-alertmanager: Alert component
* bmp-pronoea: Receives alert information from bmp-alertmanager, converts the format, and passes it to bmp-openapi

### Installation Process
![Architecture Diagram](./bmp-deploy/picture/bm-deploy.png)
### Installation Process Explanation
1. The client (bmp-console-web) initiates an installation request, which is received by bmp-console-api.
2. bmp-console-api checks the request parameters, and if they pass, forwards the request to bmp-openapi.
3. bmp-openapi performs permission checks and other operations, generates installation parameters, and sends them to bmp-scheduler.
4. bmp-scheduler schedules the installation task, generates a series of installation instructions, and sends them to bmp-dhcp-agent, bmp-driver, and bmp-agent through the bmp-mq service.
5. bmp-dhcp-agent receives the instructions, sets up the built-in DHCP server, so that the bm node can obtain the correct IP configuration and TFTP address (bmp-tftp address) during the PXE boot phase.
6. bmp-driver receives the instructions, sets the bm node to PXE boot, and restarts it.
7. The bm node performs PXE boot, the PXEClient built into the network card starts, sends a DHCP request broadcast, and the DHCP server built into bmp-dhcp-agent receives the DHCP request and responds with the corresponding IP configuration and TFTP address.
8. PXEClient configures its own IP address, then downloads the PXE boot program and executes it from bmp-tftp, the PXE boot program continues to obtain other boot parameters from bmp-tftp, downloads the kernel and initramfs, starts the memory operating system, and the bmp-agent service built into the memory operating system starts.
9. bmp-agent receives the instructions and performs subsequent bm installation operations, such as setting up RAID and partitioning.
10. bmp-agent downloads the customer operating system image file from bmp-image, writes it to the bm node's disk, and then initializes the customer operating system.
11. bmp-agent executes a restart, completing the operating system installation.

### Monitoring Process
![Monitoring Diagram](./bmp-deploy/picture/monitor.png)
1. Monitoring probes collect host monitoring information and summarize it to the bmp-monitor-proxy component.
2. bmp-monitor-proxy passes the monitoring information to bmp-pushgateway.
3. bmp-prometheus periodically pulls monitoring data from bmp-pushgateway.
4. When bmp-prometheus triggers an alert rule, the alert information is passed to bmp-alertmanager.
5. bmp-alertmanager passes the alert information to bmp-pronoea for format conversion.
6. bmp-pronoea passes the alert information to bmp-api for alert display and alerting.

### Glossary
* manager node: BMP management server, runs all BMP components (except bmp-agent).
* bm node: Bare metal server, a physical server used for normal work, has no operating system before installation, runs LiveOS during the installation phase, and runs GuestOS after installation is complete.
* GuestOS: Normal working operating system
* LiveOS: Memory operating system, with bmp-agent pre-installed.
* Out-of-band network card: A special network card on the physical server used for communication with BMC, also known as the IPMI network card.
* Management network card: The standard network card on the physical server, located in the management network.
* Management network: A 3-layer network, where the management network card of the manager node and the management network card of the bm node communicate through the management network.

## Single Component Deployment and Debugging<a id="2"></a>
### BMP Code<a id="2.1"></a>
#### Code Introduction<a id="2.1.1"></a>

| Component            | Category   | Language/Component |
|-------------------|-----------|-------------------|
| bmp-image         | Base Component | nginx |
| bmp-tftp          | Base Component | tftp |
| bmp-db            | Base Component | mysql |
| bmp-redis         | Base Component | redis |
| bmp-mq            | Base Component | rabbitmq |
| bmp-console-web   | Frontend   | vue |
| bmp-operation-web | Frontend   | vue |
| bmp-console-api   | Backend   | go |
| bmp-operation-api | Backend   | go |
| bmp-openapi       | Backend   | go |
| bmp-scheduler     | Backend   | go |
| bmp-openapi-console | Backend   | go |
| bmp-driver        | Backend   | go |
| bmp-oob-alert     | Backend   | go |
| bmp-oob-agent     | Backend   | go |
| bmp-pronoea       | Backend   | go |
| bmp-monitor-proxy | Backend   | go |
| bmp-dhcp-agent    | Backend   | python |
| bmp-rsyslog       | Open-Source Component | rsyslog |
| bmp-pushgateway   | Open-Source Component | pushgateway |
| bmp-prometheus    | Open-Source Component | prometheus |
| bmp-alertmanager  | Open-Source Component | alertmanager |


### 2.1.2 Setting Environment Variables
Backend components will call each other. Plan the ports of each component and write them into the configuration file `.env`.
For example: This environment is deployed on IP 192.168.14.80, and the configuration file `.env` is as follows. Please change the IP address to the actual one.
```bash
set -a
BMP_DB_HOST=192.168.14.80
BMP_DB_PORT=3306
BMP_DB_USER=bmp_rw
BMP_DB_PASSWORD='LpK9Jq12Zf!'
BMP_DB_NAME=bmp
BMP_REDIS_HOST=192.168.14.80
BMP_REDIS_PORT=6379
BMP_REDIS_PASSWORD='LpK9Jq12Zf'
BMP_MQ_HOST=192.168.14.80
BMP_MQ_PORT=5672
BMP_MQ_USER=bmp
BMP_MQ_PASSWORD='LpK9Jq12Zf'
BMP_MQ_VHOST=/bmp
BMP_OMAPI_HOST=192.168.14.80
BMP_OMAPI_PORT=7911
BMP_OMAPI_TOKEN="LpK9Jq12Zf"
BMP_OPENAPI_HOST=192.168.14.80
BMP_OPENAPI_PORT=8801
BMP_OPENAPI_CONSOLE_HOST=192.168.14.80
BMP_OPENAPI_CONSOLE_PORT=8802
BMP_OOB_ALERT_HOST=192.168.14.80
BMP_OOB_ALERT_PORT=8804
BMP_MQ_EXCHANGE_ROUTING_KEY=idc-vm4xsulx1k2d9z4xkctrttig02zl
BMP_IMAGE_PORT=10000
BMP_MONITOR_PROXY_PORT=8805
BMP_PROMETHEUS_HOST=192.168.14.80
BMP_PROMETHEUS_PORT=9090
BMP_PROMETHEUS_DATA_DIR=/data/prometheus
BMP_PUSHGATEWAY_HOST=192.168.14.80
BMP_PUSHGATEWAY_PORT=9091
BMP_ALERTMANAGER_HOST=192.168.14.80
BMP_ALERTMANAGER_PORT=9093
BMP_PRONOEA_HOST=192.168.14.80
BMP_PRONOEA_PORT=9999
BMP_RSYSLOG_PORT=1514
set +a
```
To make the environment variables take effect:
```bash
source.env
```
### 2.2 Basic Components

#### 2.2.1 bmp-image
It mainly stores GuestOS image files, bmp-agent (the agent of the installation node), and device_import template (Excel file), and provides HTTP services externally through Nginx.
```bash
# Step 1: Install Nginx, taking CentOS as an example
yum install nginx
systemctl start nginx

# Step 2: Download the following files from JD Cloud Object Storage and save them to /home/bmp/image
mkdir -p /home/bmp/bmp-image
chmod -R 755 /home/bmp/bmp-image
cd /home/bmp/bmp-image
# GuestOS image
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/GuestOS/v1.7.0-centos-7.9-2022070716.tar.xz 
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/GuestOS/v1.7.0-ubuntu-18.04-2022062709.tar.xz
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/GuestOS/v1.7.0-centos-7.9-arm-2023080716.tar.xz
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/GuestOS/v1.7.0-ubuntu-18.04-arm-2023081111.tar.xz
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/GuestOS/v1.7.0-loongnix-8.4-2023110218.tar.xz

# bmp-agent
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/others/bmp-agent.bin
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/others/bmp-agent.bin.arm
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/others/bmp-agent-windows.tgz
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/others/device_import.xlsx

# Step 3: Add Nginx configuration file
cat << EOF > /etc/nginx/conf.d/bmp-image.conf
server {
    listen       10000;
    listen  [::]:10000;
    server_name  localhost;
    location / {
        root   /home/bmp/bmp-image;
        index  index.html index.htm;
    }
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /home/bmp/bmp-image;
    }
}
EOF


# Step 4: Start Nginx
nginx -s reload

# For more Nginx configurations, please refer to the official website:
https://nginx.org/en/docs/
```

### 2.2.2 Setting the tftp.env Environment Variable
```bash
set -a
BMP_DB_HOST=192.168.14.80
BMP_DB_PORT=3306
BMP_DB_USER=bmp_rw
BMP_DB_PASSWORD='LpK9Jq12Zf!'
BMP_DB_NAME=bmp
BMP_REDIS_HOST=192.168.14.80
BMP_REDIS_PORT=6379
BMP_REDIS_PASSWORD='LpK9Jq12Zf'
BMP_MQ_HOST=192.168.14.80
BMP_MQ_PORT=5672
BMP_MQ_USER=bmp
BMP_MQ_PASSWORD='LpK9Jq12Zf'
BMP_MQ_VHOST=/bmp
BMP_OMAPI_HOST=192.168.14.80
BMP_OMAPI_PORT=7911
BMP_OMAPI_TOKEN="LpK9Jq12Zf"
BMP_OPENAPI_HOST=192.168.14.80
BMP_OPENAPI_PORT=8801
BMP_OPENAPI_CONSOLE_HOST=192.168.14.80
BMP_OPENAPI_CONSOLE_PORT=8802
BMP_OOB_ALERT_HOST=192.168.14.80
BMP_OOB_ALERT_PORT=8804
BMP_MQ_EXCHANGE_ROUTING_KEY=idc-vm4xsulx1k2d9z4xkctrttig02zl
BMP_IMAGE_PORT=10000
BMP_MONITOR_PROXY_PORT=8805
BMP_PROMETHEUS_HOST=192.168.14.80
BMP_PROMETHEUS_PORT=9090
BMP_PROMETHEUS_DATA_DIR=/data/prometheus
BMP_PUSHGATEWAY_HOST=192.168.14.80
BMP_PUSHGATEWAY_PORT=9091
BMP_ALERTMANAGER_HOST=192.168.14.80
BMP_ALERTMANAGER_PORT=9093
BMP_PRONOEA_HOST=192.168.14.80
BMP_PRONOEA_PORT=9999
BMP_RSYSLOG_PORT=1514

# bootloader
bmp_bootloader_prefix=bootloader
bmp_bootloader_images=(
    BOOTLOONGARCH64.EFI
    grubaa64.efi
    grubx64.efi
    pxelinux.0
)

# LiveOS image
bmp_oss_liveos_prefix=LiveOS

bmp_kernel_name_x86=v2.0.7-centos_7_9-2024082914-vmlinuz
bmp_initramfs_name_x86=v2.0.7-centos_7_9-2024082914-initramfs.gz

bmp_kernel_name_arm64=v2.0.7-centos_7_9_arm64-2024082914-vmlinuz
bmp_initramfs_name_arm64=v2.0.7-centos_7_9_arm64-2024082914-initramfs.gz

bmp_kernel_name_loonarch64=vmlinuz-loongarch
bmp_initramfs_name_loonarch64=initramfs-loongarch.gz

bmp_liveos_images=(
    ${bmp_kernel_name_x86}
    ${bmp_initramfs_name_x86}
    ${bmp_kernel_name_arm64}
    ${bmp_initramfs_name_arm64}
    ${bmp_kernel_name_loonarch64}
    ${bmp_initramfs_name_loonarch64}
)
BMP_KERNEL_PATH_x86=/${bmp_kernel_name_x86}
BMP_INITRAMFS_PATH_x86=/${bmp_initramfs_name_x86}
BMP_KERNEL_PATH_arm64=/images/arm64/${bmp_kernel_name_arm64}
BMP_INITRAMFS_PATH_arm64=/images/arm64/${bmp_initramfs_name_arm64}
BMP_KERNEL_PATH_loonarch64=/images/loongarch64/${bmp_kernel_name_loonarch64}
BMP_INITRAMFS_PATH_loonarch64=/images/loongarch64/${bmp_initramfs_name_loonarch64}
set +a
```
To make it effective:
```bash
source tftp.env
```
### 2.2.2 Downloading Kernel and LiveOS Files from JD Cloud Object Storage
```bash
mkdir -p /var/lib/tftpboot/images/{arm64,loongarch64}
mkdir -p /var/lib/tftpboot/uefi/{arm64,loongarch64,x86_64}
mkdir -p /var/lib/tftpboot/pxelinux.cfg

# arm64-image
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/LiveOS/v2.0.7-centos_7_9_arm64-2024082914-initramfs.gz -P /var/lib/tftpboot/images/arm64
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/LiveOS/v2.0.7-centos_7_9_arm64-2024082914-vmlinuz -P /var/lib/tftpboot/images/arm64
# loongarch64-image
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/LiveOS/initramfs-loongarch.gz -P /var/lib/tftpboot/images/loongarch64
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/LiveOS/initramfs-loongarch.gz -P /var/lib/tftpboot/images/loongarch64
# x86_64-image
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/LiveOS/v2.0.7-centos_7_9-2024082914-initramfs.gz -P /var/lib/tftpboot/
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/LiveOS/v2.0.7-centos_7_9-2024082914-vmlinuz -P /var/lib/tftpboot/
# bootloader
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/bootloader/BOOTLOONGARCH64.EFI -P /var/lib/tftpboot/uefi/loongarch64
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/bootloader/grubaa64.efi -P /var/lib/tftpboot/uefi/x86_64
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/bootloader/grubx64.efi -P /var/lib/tftpboot/uefi/x86_64
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/bootloader/pxelinux.0 -P /var/lib/tftpboot/
```
### 2.2.2 Setting Configuration Files
```bash
cat << EOF > /var/lib/tftpboot/pxelinux.cfg/default
default bmp agent
    prompt 0
    timeout 1
    label bmp agent
    kernel ${BMP_KERNEL_PATH_x86}
    append initrd=${BMP_INITRAMFS_PATH_x86} net.ifnames=0 biosdevname=0 BMP_MQ_HOST=${BMP_HOST_IP} BMP_MQ_PORT=${BMP_MQ_PORT} BMP_MQ_USER=${BMP_MQ_USER} BMP_MQ_PASSWORD=${BMP_MQ_PASSWORD} BMP_MQ_VHOST=${BMP_MQ_VHOST} BMP_MQ_EXCHANGE_ROUTING_KEY=${BMP_MQ_EXCHANGE_ROUTING_KEY} BMP_IMAGE_HOST=${BMP_HOST_IP} BMP_IMAGE_PORT=${BMP_IMAGE_PORT} BMP_RSYSLOG_HSOT=${BMP_HOST_IP} BMP_RSYSLOG_PORT=${BMP_RSYSLOG_PORT}
EOF

cat << EOF > /var/lib/tftpboot/uefi/arm64/grub.cfg
set timeout=1
set default=1
menuentry "bmp agent arm64 uefi" {
    linux ${BMP_KERNEL_PATH_arm64} console=ttyAMA0 console=tty0 net.ifnames=0 biosdevname=0 ksdevice=bootif kssendmac text BMP_MQ_HOST=${BMP_HOST_IP} BMP_MQ_PORT=${BMP_MQ_PORT} BMP_MQ_USER=${BMP_MQ_USER} BMP_MQ_PASSWORD='${BMP_MQ_PASSWORD}' BMP_MQ_VHOST=${BMP_MQ_VHOST} BMP_MQ_EXCHANGE_ROUTING_KEY=${BMP_MQ_EXCHANGE_ROUTING_KEY} BMP_IMAGE_HOST=${BMP_HOST_IP} BMP_IMAGE_PORT=${BMP_IMAGE_PORT}
    initrd ${BMP_INITRAMFS_PATH_arm64}
}
EOF

cat << EOF > /var/lib/tftpboot/uefi/loongarch64/grub.cfg
set timeout=1
set default=1
menuentry 'bmp agent loongarch64 uefi' {
    linux ${BMP_KERNEL_PATH_loonarch64} console=ttyS0,115200 console=tty0 net.ifnames=0 biosdevname=0 ksdevice=bootif kssendmac text BMP_MQ_HOST=${BMP_HOST_IP} BMP_MQ_PORT=${BMP_MQ_PORT} BMP_MQ_USER=${BMP_MQ_USER} BMP_MQ_PASSWORD=${BMP_MQ_PASSWORD} BMP_MQ_VHOST=${BMP_MQ_VHOST} BMP_MQ_EXCHANGE_ROUTING_KEY=${BMP_MQ_EXCHANGE_ROUTING_KEY} BMP_IMAGE_HOST=${BMP_HOST_IP} BMP_IMAGE_PORT=${BMP_IMAGE_PORT}
    initrd ${BMP_INITRAMFS_PATH_loonarch64}
}
EOF

cat << EOF > /var/lib/tftpboot/uefi/x86_64/grub.cfg
set timeout=1
set default=1
menuentry 'bmp agent x86_64 uefi' {
    linuxefi ${BMP_KERNEL_PATH_x86}  net.ifnames=0 biosdevname=0 ksdevice=bootif kssendmac text BMP_MQ_HOST=${BMP_HOST_IP} BMP_MQ_PORT=${BMP_MQ_PORT} BMP_MQ_USER=${BMP_MQ_USER} BMP_MQ_PASSWORD=${BMP_MQ_PASSWORD} BMP_MQ_VHOST=${BMP_MQ_VHOST} BMP_MQ_EXCHANGE_ROUTING_KEY=${BMP_MQ_EXCHANGE_ROUTING_KEY} BMP_IMAGE_HOST=${BMP_HOST_IP} BMP_IMAGE_PORT=${BMP_IMAGE_PORT}
    initrdefi ${BMP_INITRAMFS_PATH_x86}
}
EOF

# The final directory structure is as follows:
.
├── images
│   ├── arm64
│   │   ├── v2.0.7-centos_7_9_arm64-2024082914-initramfs.gz
│   │   └── v2.0.7-centos_7_9_arm64-2024082914-vmlinuz
│   └── loongarch64
│       ├── initramfs-loongarch.gz
│       └── vmlinuz-loongarch
├── pxelinux.0
├── pxelinux.cfg
│   └── default
├── uefi
│   ├── arm64
│   │   ├── grubaa64.efi
│   │   └── grub.cfg
│   ├── loongarch64
│   │   ├── BOOTLOONGARCH64.EFI
│   │   └── grub.cfg
│   └── x86_64
│       ├── grub.cfg
│       └── grubx64.efi
├── v2.0.7-centos_7_9-2024082914-initramfs.gz
└── v2.0.7-centos_7_9-2024082914-vmlinuz
```
### 2.2.2 Starting the TFTP Service
```bash
systemctl restart tftp
# Checking the status
systemctl status tftp
```
#### 2.2.3 bmp - mysql
| Application Component | Category | Language/Component |
|-------------------|------| --- |
| bmp - db | Basic Component | mysql |

The following are the basic steps to install and deploy MySQL on a Linux system:
```bash
# Step 1: Install MySQL
sudo yum install mysql-community-server

# Step 2: Start the MySQL service
sudo systemctl start mysqld

# Step 3: After installation, MySQL generates a temporary password. You can obtain this temporary password by viewing the log file
grep 'temporary password' /var/log/mysqld.log

# Step 4: Log in to MySQL
mysql -u root -p
Enter the temporary password you found in the previous step.

# Step 5: Create a new user and database
ALTER USER 'root'@'localhost' IDENTIFIED BY 'LpK9Jq12Zf!';
CREATE USER 'bmp_rw'@'%' IDENTIFIED BY 'LpK9Jq12Zf!';
CREATE DATABASE bmp;
GRANT ALL PRIVILEGES ON bmp.* TO 'bmp_rw'@'%';
FLUSH PRIVILEGES;

# Step 6: Import initialization data
The data is located in bmp - deploy/sql/bmp.sql. The import command is
cd bmp - deploy
mysql -uroot -p bmp <./sql/bmp.sql
Enter the set root password.
```
#### 2.2.4 bmp - redis
| Application Component | Category | Language/Component |
|-------------------|------| --- |
| bmp - redis | Basic Component | redis |

```bash
The following steps can be used to install Redis on CentOS 7:

# Step 1: Install Redis
sudo yum install redis -y

# Step 2: Configure the Redis password
The Redis configuration file is located at /etc/redis.conf
vi /etc/redis.conf
Append a line to set the password
requirepass LpK9Jq12Zf
Append a line to set the bind address, that is, which addresses are allowed to access this Redis. This example is for testing, set to 0.0.0.0, that is, all addresses are accessible. Please set according to the actual situation in the production environment
bind 0.0.0.0

# Step 3: Start the Redis service
sudo systemctl start redis

# Step 4: Test the password
redis-cli -a LpK9Jq12Zf ping
```
#### 2.2.5 bmp - mq
| Application Component | Category | Language/Component |
|-------------------|------| --- |
| bmp - mq | Basic Component | rabbitmq |

```bash
The following steps can be used to install RabbitMQ on CentOS 7:

# Step 1: Install Erlang
# RabbitMQ is written in Erlang, so we need to install Erlang first:
sudo yum install erlang -y

# Step 2: Install RabbitMQ
sudo yum install rabbitmq-server -y

# Step 3: Start the RabbitMQ service
sudo systemctl start rabbitmq-server

# Step 4: Create an administrator account
# By default, the RabbitMQ management plugin is not enabled. We need to create an administrator account and enable the management plugin:
rabbitmqctl add_vhost /bmp
rabbitmqctl add_user bmp "LpK9Jq12Zf"
rabbitmqctl set_user_tags bmp administrator
rabbitmqctl set_permissions -p "/bmp" bmp '.*' '.*' '.*'
rabbitmqctl list_users

### Step 5: Access the RabbitMQ management interface
Now, you can access the RabbitMQ management interface through a browser at `http://your-server-ip:15672`. Log in with the administrator account you just created.
```
### bmp-rsyslog <a id="2.2.6"></a>

| Application Component | Category | Language/Component |
|-------------------|------| --- |
| bmp-rsyslog | Open-source Component | rsyslog |


```bash
The following steps can be used to install rsyslog on CentOS 7:

# Step 1: Install rsyslog
sudo yum install rsyslog -y

# Step 2: Configure rsyslog
cat << EOF > /etc/rsyslog.d/bmp-rsyslog.conf
module(load="imudp") # needs to be done just once
input(type="imudp" port="514")
\$template RemoteLogs,"/var/log/bmp/bmp-rsyslog/%fromhost-ip%/%PROGRAMNAME%-%\$YEAR%-%\$MONTH%-%\$DAY%.log"
\$template DynamicDir,"/var/log/bmp/bmp-rsyslog/%fromhost-ip%"
:syslogtag, startswith, "ip"?DynamicDir
if $fromhost-ip!= "127.0.0.1" then?RemoteLogs
& ~
EOF

# Step 3: Start the rsyslog service
sudo systemctl start rsyslog
```

### Front-end Component Vue <a id="2.3"></a>

| Application Component | Category | Language/Component |
|-------------------|------| --- |
| bmp-console-web | Front-end | vue |
| bmp-operation-web | Front-end | vue |

```bash
# Node version requirement
Version 16.19.0

# Install Node.js
wget https://nodejs.org/dist/v16.19.0/node-v16.19.0-linux-x64.tar.xz
tar -xvf node-v16.19.0-linux-x64.tar.xz 
mv node-v16.19.0-linux-x64 /usr/local/nodejs
cd /usr/bin
ln -s /usr/local/nodejs/bin/node node
ln -s /usr/local/nodejs/bin/npm npm
ln -s /usr/local/nodejs/bin/npx npx

# Configure the NPM mirror source
npm config set registry https://registry.npmjs.org 
# You can configure a domestic mirror source according to your network situation
npm config set registry https://registry.npmmirror.com

# Install dependencies
npm install --legacy-peer-deps 

# Package
npm run build-pre

# Local development and testing
npm run dev
```

The above compilation and packaging commands will generate a `dist` directory. Copy the `dist` directory to the main directory of Nginx respectively:

- `/home/bmp/bmp-console-web`
- `/home/bmp/bmp-operation-web`

```bash
mkdir -p /home/bmp/bmp-console-web
mkdir -p /home/bmp/bmp-operation-web
cp -r bmp-console-web/dist/* /home/bmp/bmp-console-web
cp -r bmp-operation-web/dist/* /home/bmp/bmp-operation-web
```

- Then configure Nginx. If Nginx has been installed, skip the installation step.

```bash
# Install Nginx
yum install nginx
systemctl start nginx
```

- Configure `bmp-console-web`

Copy the following content to `/etc/nginx/conf.d/bmp-console-web.conf`:

```bash
log_format  main1  '$time_local | $remote_addr | $upstream_addr | $http_host | $status | $body_bytes_sent | $upstream_response_time | $request_time | $http_host | $request | $http_referer';

server {
    listen 8080;
    server_name bmp-console.bmp.local;
    charset utf-8;
    access_log  /var/log/nginx/access.log  main1;

    error_page 500 502 503 504 /50x.html;
    location = /50x.html {
        root html;
    }

    root /home/bmp/bmp-console-web;
    index index.html;

    location /console-web/ {
        set $bmp_console_api "127.0.0.1:8800";
        rewrite ^/console-web/(.*) /$1 break;
        proxy_pass http://$bmp_console_api;
    }
    
    location /oob-alert/ {
        client_max_body_size 51200M;
        set $bmp_oob_alert "127.0.0.1:8804";
        rewrite ^/oob-alert/(.*) /$1 break;
        proxy_pass http://$bmp_oob_alert;      
    }   

    location / {
        try_files $uri $uri/ @router;
        index index.html;
    }

    location @router {
        rewrite ^.*$ /index.html last;
    }
    location ~.*\.(html)$ {
        add_header Cache-Control no-cache;
    }
}
```

- Configure `bmp-operation-web`

Copy the following content to `/etc/nginx/conf.d/bmp-operation-web.conf`:

```bash
server {
    listen 8081;
    server_name bmp-operation.bmp.local;
    charset utf-8;
    access_log  /var/log/nginx/access.log  main1;

    error_page 500 502 503 504 /50x.html;
    location = /50x.html {
        root html;
    }

    root /home/bmp/bmp-operation-web;
    index index.html;

    # compression-webpack-plugin configuration
    #gzip on;
    #gzip_min_length 1k;
    #gzip_comp_level 9;
    #gzip_types text/plain application/javascript application/x-javascript text/css application/xml text/javascript applicationx-httpd-php image/jpeg image/gif image/png;
    #gzip_vary on;
    # Configure the condition to disable gzip, support regular expressions. Here, it means that gzip is not enabled for IE6 and below (because low versions of IE do not support it).
    #gzip_disable "MSIE [1-6]\.";

    location /operation-web/ {
        client_max_body_size 51200M;
        set $bmp_operation_api "127.0.0.1:8799";
        rewrite ^/operation-web/(.*) /$1 break;
        proxy_pass http://$bmp_operation_api;      
    }
    
    location /oob-alert/ {
        client_max_body_size 51200M;
        set $bmp_oob_alert "127.0.0.1:8804";
        rewrite ^/oob-alert/(.*) /$1 break;
        proxy_pass http://$bmp_oob_alert;      
    }
    
    location / {
        try_files $uri $uri/ @router;
        index index.html;
    }

    location @router {
        rewrite ^.*$ /index.html last;
    }
    location ~.*\.(html)$ {
        add_header Cache-Control no-cache;
    }
}
```

Reload the Nginx configuration:

```bash
nginx -s reload
``` 
### Back-end Components <a id="2.4"></a>
#### Back-end Go Components <a id="2.4.1"></a>

| Application Component | Category | Language/Component |
|-------------------|------| --- |
| bmp-console-api | Back-end | go |
| bmp-operation-api | Back-end | go |
| bmp-openapi | Back-end | go |
| bmp-scheduler | Back-end | go |
| bmp-openapi-console | Back-end | go |
| bmp-driver | Back-end | go |
| bmp-oob-alert | Back-end | go |
| bmp-oob-agent | Back-end | go |
| bmp-pronoea | Back-end | go |
| bmp-monitor-proxy | Back-end | go |

1. The back-end components rely on the following basic components. Before debugging the back-end code locally, start the following components first:
    - [bmp-db](#2.2.3)
    - [bmp-redis](#2.2.4)
    - [bmp-mq](#2.2.5)

2. Create the log directories for each component first.
```bash
# For example
mkdir -p /var/log/bmp/bmp-openapi
```
3. [Set environment variables](#2.1.2)
4. Compile and run.
```bash
# Version requirement: golang:1.17

# Install golang:1.17
wget https://golang.google.cn/dl/go1.17.linux-amd64.tar.gz
tar -zxf go1.17.linux-amd64.tar.gz -C /usr/local

# Add the following configuration at the end of the /etc/profile file. Press 'esc' and then input ':wq' to save.
# Golang config
export GOROOT=/usr/local/go 
export GOPATH=/local/gopath
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

# Create /local/gopath
mkdir -p /local/gopath

# Make the environment variables take effect
source /etc/profile

# Check the Go version
go version

# Set the proxy. You can choose different proxies according to your own network.
export GOPROXY=https://goproxy.cn

# Local compilation and debugging. For example, bmp-openapi
cd bmp-openapi

# Compile
go build -o bmp-openapi

# Run
./bmp-openapi
# During the joint debugging phase, each component needs to run in the background.
nohup./bmp-openapi &
```

Please note that the back-end services rely on middleware and other components. If you conduct local joint debugging, you need to deploy the *dependent components* in advance. For specific dependencies, see the `conf` directory files in each Go code. For example, the `conf` file `bmp-openapi.ini` of `bmp-openapi` is as follows:

```bash
# mysql
jdbc.url=
bmp_db_host=${BMP_DB_HOST}
bmp_db_port=${BMP_DB_PORT}
bmp_db_user=${BMP_DB_USER}
bmp_db_password=${BMP_DB_PASSWORD}
bmp_db_name=${BMP_DB_NAME}

bmp_redis_host=${BMP_REDIS_HOST}
bmp_redis_port=${BMP_REDIS_PORT}
bmp_redis_password=${BMP_REDIS_PASSWORD}


# RabbitMQ for Ironic
bmp_mq_host=${BMP_MQ_HOST}
bmp_mq_port=${BMP_MQ_PORT}
bmp_mq_user=${BMP_MQ_USER}
bmp_mq_password=${BMP_MQ_PASSWORD}
bmp_mq_vhost=${BMP_MQ_VHOST}
bmp_mq_exchange=CPS_IRONIC_SCHEDULER
bmp_mq_receive_routing_key=/test

bmp_monitor_proxy_host=${BMP_MONITOR_PROXY_HOST||192.168.12.72}
bmp_monitor_proxy_port=${BMP_MONITOR_PROXY_PORT||8805}

bmp_pronoea_host=${BMP_PRONOEA_HOST||192.168.12.75}
bmp_pronoea_port=${BMP_PRONOEA_PORT||9999}

bmp_image_host=${BMP_IMAGE_HOST||192.168.12.72}
bmp_image_port=${BMP_IMAGE_PORT||10000}
```

### Back-end Python Components <a id="2.4.2"></a>

| Application Component | Category | Language/Component |
|-------------------|------| --- |
| bmp-dhcp-agent | Back-end | python |

This component relies on environment variables. Please complete the [Port Planning](#2.1.2) section first (this section is **mandatory**). And set the following environment variables.

```bash
cat bmp-dhcp-agent.env
set -a
BMP_MQ_PORT=5672
BMP_MQ_USER=bmp
BMP_MQ_PASSWORD='LpK9Jq12Zf'
BMP_MQ_VHOST=/bmp
BMP_OMAPI_HOST=${BMP_HOST_IP}
BMP_OMAPI_PORT=7911
BMP_OMAPI_TOKEN="LpK9Jq12Zf"
BMP_OMAPI_KEY=$(echo -n ${BMP_OMAPI_TOKEN} | base64)

BMP_MQ_HOST=${BMP_HOST_IP}
BMP_OMAPI_HOST=${BMP_HOST_IP}
BMP_DHCP_CONFIG_DIR=/home/dhcp/data
BMP_DHCP_CONTROL_BIN=/home/dhcp/dhcpd_control.sh
# BMP_TFTP_CONFIG_DIR needs to be consistent with the tftp document path in the tftp installation steps
BMP_TFTP_CONFIG_DIR=/var/lib/tftpboot
BMP_TFTP_MQ_HOST=${BMP_HOST_IP}
BMP_TFTP_MQ_PORT=${BMP_MQ_PORT}
BMP_TFTP_MQ_USER=${BMP_MQ_USER}
BMP_TFTP_MQ_PASSWORD=${BMP_MQ_PASSWORD}
BMP_TFTP_MQ_VHOST=${BMP_MQ_VHOST}
BMP_TFTP_MQ_EXCHANGE_ROUTING_KEY=${BMP_MQ_EXCHANGE_ROUTING_KEY}
BMP_TFTP_IMAGE_HOST=${BMP_HOST_IP}
BMP_TFTP_IMAGE_PORT=${BMP_IMAGE_PORT}
BMP_TFTP_RSYSLOG_HSOT=${BMP_HOST_IP}
BMP_TFTP_RSYSLOG_PORT=${BMP_RSYSLOG_PORT}
set +a
```

```bash
source  bmp-dhcp-agent.env
```

Create the `dhcp` directory.
```bash
mkdir -p /home/dhcp/data
```

Edit the configuration file `/home/dhcp/data/dhcpd.conf`.
```bash
cat << EOF > /home/dhcp/data/dhcpd.conf
ddns-update-style none;
ignore client-updates;
ddns-update-style none;
ignore client-updates;

next-server ${BMP_HOST_IP};
filename "pxelinux.0";

default-lease-time         600;
max-lease-time             1200;
option domain-name "localhost";
option domain-name-servers 114.114.114.114,8.8.8.8;

#pxe
option client-system-architecture code 93 = unsigned integer 16;
class "pxe-clients" {
    match if (substring (option vendor-class-identifier, 0, 9) = "PXEClient") or
             (substring (option vendor-class-identifier, 0, 9) = "HW-Client");
    if option client-system-architecture = 00:00 {
        filename "pxelinux.0";
    }else if option client-system-architecture = 00:07 {
        filename "/uefi/x86_64/grubx64.efi";
    }else if option client-system-architecture = 00:0b {
        filename "/uefi/arm64/grubaa64.efi";
    }else if option client-system-architecture = 00:0c {
        filename "/uefi/loongarch64/BOOTLOONGARCH64.EFI";
    }else if option client-system-architecture = 00:27 {
        filename "/uefi/loongarch64/BOOTLOONGARCH64.EFI";
    }
}

key omapi_key {
    algorithm hmac-md5;
    secret $(echo -n ${BMP_OMAPI_TOKEN} | base64);
};
omapi-key omapi_key;
omapi-port ${BMP_OMAPI_PORT};

subnet ${BMP_HOST_IP} netmask 255.255.255.255 {
    
}
EOF
```

Edit the startup file `/home/dhcp/dhcpd_control.sh`.
```bash
#!/bin/sh
set -ue

[ -z "$BMP_DHCP_CONFIG_DIR" ] && echo "error! BMP_DHCP_CONFIG_DIR is unset!" && exit 1

if [! -e $BMP_DHCP_CONFIG_DIR/dhcpd.conf ] && [ -e /dhcpd.conf.tpl ]; then
    eval "echo "$(cat /dhcpd.conf.tpl)"" > $BMP_DHCP_CONFIG_DIR/dhcpd.conf
fi

LOG_FILE=/var/log/bmp/bmp-dhcp-agent/dhcpd.log
exec_file=/usr/sbin/dhcpd
data_dir=$BMP_DHCP_CONFIG_DIR
dhcpd_conf="$data_dir/dhcpd.conf"
PID=""
if [! -r "$dhcpd_conf" ]; then
    echo "Please ensure '$dhcpd_conf' exists and is readable."
    echo "Run the container with arguments 'man dhcpd.conf' if you need help with creating the configuration."
    exit 1
fi

[ -e "$data_dir/dhcpd.leases" ] || touch "$data_dir/dhcpd.leases"


help(){
    echo "${0} <start|stop|restart|status>"
    exit 1
}

get_pid()
{
    local pid=""
    pid=$(ps -eo pid,comm,args | grep "$exec_file" | grep -v 'grep' | awk '{print $1}')
    PID=$pid
}

checkhealth(){
    get_pid
    [ X"$PID"!= X"" ] && echo "dhcpd is running" && return 0
    echo "dhcpd is not running" && return 1
}

start(){
    /usr/sbin/dhcpd -f -d --no-pid -cf "$data_dir/dhcpd.conf" -lf "$data_dir/dhcpd.leases" >> $LOG_FILE 2>&1 &
}

stop(){
    checkhealth || return 0
    [ X"$PID"!= X"" ] && kill $PID
    sleep 2
    checkhealth || return 0
    return 1
}

case "${1}" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    status|health|checkhealth)
        checkhealth
        ;;
    restart)
        stop && start
        ;;
    *)
        help
        ;;
esac
```

```bash
# A python3.6+ environment is required.

# Install the dhcp service locally. For example, on a CentOS system
sudo yum install dhcp
sudo yum install python3
# Install Python dependencies
pip3 install -i https://pypi.tuna.tsinghua.edu.cn/simple/ -r requirements.txt

# Run
cd bmp-dhcp-agent
export PYTHONPATH=$(pwd)
python3./bmpda/cmd/server.py
``` 

### Open-source Components <a id="2.5"></a>

| Application Component | Category | Language/Component |
|-------------------|------| --- |
| bmp-prometheus | Open-source Component | prometheus |
| bmp-pushgateway | Open-source Component | pushgateway |
| bmp-alertmanager | Open-source Component | alertmanager |

The above three monitoring and alerting components are open-source Prometheus solutions and can be installed and used together.

The working principle of the three is as follows:
1. Applications or services push metric data to the Pushgateway.
2. The Prometheus server periodically pulls these metric data from the Pushgateway and stores them in its own time series database.
3. The Alertmanager component sends alert messages for the information that triggers the alert conditions in Prometheus.

#### bmp - pushgateway <a id="2.5.1"></a>
```bash
The following steps can be used to install Pushgateway on CentOS 7:

# Step 1: Download the Pushgateway binary file
First, download the Pushgateway binary file. You can obtain it from the official Prometheus website or GitHub repository. For example, use the following command to download Pushgateway v1.4.0:
wget https://github.com/prometheus/pushgateway/releases/download/v1.4.0/pushgateway-1.4.0.linux-amd64.tar.gz
If the network is restricted, you can use other methods to download it:
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/opensource/pushgateway-1.4.0.linux-amd64.tar.gz

# Step 2: Extract the Pushgateway file
tar -xvf pushgateway-1.4.0.linux-amd64.tar.gz

# Step 3: Create the Pushgateway user and group
sudo groupadd --system pushgateway
sudo useradd --system -s /bin/false -g pushgateway pushgateway

# Step 4: Move the Pushgateway file to an appropriate location
sudo mv pushgateway-1.4.0.linux-amd64 /usr/local/pushgateway

# Step 5: Set file permissions
sudo chown -R pushgateway:pushgateway /usr/local/pushgateway
sudo chmod -R 755 /usr/local/pushgateway

# Step 6: Create the Pushgateway configuration file
sudo mkdir /etc/pushgateway
sudo touch /etc/pushgateway/pushgateway.yml

# Step 7: Create the Pushgateway service file
sudo cat << EOF > /etc/systemd/system/pushgateway.service
[Unit]
Description=Pushgateway
Wants=network-online.target
[Service]
User=pushgateway
Group=pushgateway
ExecStart=/usr/local/pushgateway/pushgateway \
    --web.config.file /etc/pushgateway/pushgateway.yml \
    --web.listen-address :9091 \
    --web.enable-admin-api
Restart=always
[Install]
WantedBy=multi-user.target
EOF

# Step 8: Start the Pushgateway service
sudo systemctl daemon-reload
sudo systemctl start pushgateway
# Step 9: Check the service status
sudo systemctl status pushgateway
```

#### bmp - prometheus <a id="2.5.2"></a>
```bash
The following steps can be used to install Prometheus on CentOS:

# Step 1: Download the Prometheus binary file
wget https://github.com/prometheus/prometheus/releases/download/v2.36.1/prometheus-2.36.1.linux-amd64.tar.gz
If the network is restricted, you can use other methods to download it:
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/opensource/prometheus-2.36.1.linux-amd64.tar.gz

# Step 2: Extract the Prometheus file
tar -xzvf prometheus-2.36.1.linux-amd64.tar.gz

# Step 3: Create the Prometheus user and group
sudo groupadd prometheus
sudo useradd -g prometheus  -d /home/prometheus prometheus

# Step 4: Move the Prometheus file to an appropriate location
sudo mv prometheus-2.36.1.linux-amd64 /usr/local/prometheus

# Step 5: Set file permissions
sudo chown -R prometheus:prometheus /usr/local/prometheus
sudo chmod -R 755 /usr/local/prometheus

# Step 6: Create the Prometheus configuration file
sudo mkdir /etc/prometheus
sudo cat << EOF > /etc/prometheus/prometheus.yml
scrape_configs:
  - job_name: 'pushgateway'
    scrape_interval: 60s
    honor_labels: true
    static_configs:
      - targets: ['\${BMP_PUSHGATEWAY_HOST}:\${BMP_PUSHGATEWAY_PORT}']
alerting:
  alertmanagers:
    - static_configs:
       - targets: ['\${BMP_ALERTMANAGER_HOST}:\${BMP_ALERTMANAGER_PORT}']
rule_files:
  - /home/prometheus/conf/rules/*.yml
EOF
# Please note that you need to set the values of the following four variables:
# BMP_ALERTMANAGER_HOST / BMP_PUSHGATEWAY_PORT / BMP_ALERTMANAGER_HOST / BMP_ALERTMANAGER_PORT

# Step 7: Create the Prometheus service file
sudo cat << EOF > /etc/systemd/system/prometheus.service
[Unit]
Description=Prometheus
Wants=network-online.target

[Service]
User=prometheus
Group=prometheus
ExecStart=/usr/local/prometheus/prometheus \
    --config.file /etc/prometheus/prometheus.yml \
    --storage.tsdb.path /home/prometheus \
    --web.console.libraries /usr/local/prometheus/console-libraries \
    --web.console.templates /usr/local/prometheus/consoles \
    --web.enable-lifecycle
Restart=always
[Install]
WantedBy=multi-user.target
EOF

# Step 8: Start the Prometheus service
sudo systemctl daemon-reload
sudo systemctl start prometheus

## Step 9: Check the Prometheus service status
sudo systemctl status prometheus
```

#### bmp - alertmanager <a id="2.5.3"></a>
```bash
# Step 1: Download the Alertmanager binary file
wget https://github.com/prometheus/alertmanager/releases/download/v0.24.0/alertmanager-0.24.0.linux-amd64.tar.gz
If the network is restricted, you can use other methods to download it:
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/opensource/alertmanager-0.24.0.linux-amd64.tar.gz
# Step 2: Extract the Alertmanager file
tar -xzvf alertmanager-0.24.0.linux-amd64.tar.gz

# Step 3: Create the Alertmanager user and group
sudo groupadd alertmanager
sudo useradd -g alertmanager -d /home/alertmanager alertmanager

# Step 4: Move the Alertmanager file to an appropriate location
sudo mv alertmanager-0.24.0.linux-amd64 /usr/local/alertmanager

# Step 5: Set file permissions
sudo chown -R alertmanager:alertmanager /usr/local/alertmanager
sudo chmod -R 755 /usr/local/alertmanager

# Step 6: Create the Alertmanager configuration file
sudo mkdir /etc/alertmanager
sudo cat << EOF > /etc/alertmanager/alertmanager.yml
global:
  resolve_timeout: 5s
route:
  group_by: ['alertname']
  group_wait: 0s
  group_interval: 60m
  repeat_interval: 10s
  receiver: 'bmpAlertReceiver'
  routes:
  - receiver: 'bmpAlertReceiver'
    group_interval: 5m
    match:
      noticePeriodLabel: NoticePeriod-5m
  - receiver: 'bmpAlertReceiver'
    group_interval: 10m
    match:
      noticePeriodLabel: NoticePeriod-10m
  - receiver: 'bmpAlertReceiver'
    group_interval: 15m
    match:
      noticePeriodLabel: NoticePeriod-15m
  - receiver: 'bmpAlertReceiver'
    group_interval: 30m
    match:
      noticePeriodLabel: NoticePeriod-30m
  - receiver: 'bmpAlertReceiver'
    group_interval: 60m
    match:
      noticePeriodLabel: NoticePeriod-60m
  - receiver: 'bmpAlertReceiver'
    group_interval: 180m
    match:
      noticePeriodLabel: NoticePeriod-180m
  - receiver: 'bmpAlertReceiver'
    group_interval: 720m
    match:
      noticePeriodLabel: NoticePeriod-720m
  - receiver: 'bmpAlertReceiver'
    group_interval: 1440m
    match:
      noticePeriodLabel: NoticePeriod-1440m
receivers:
  - name: 'bmpAlertReceiver'
    webhook_configs:
      - url: 'http://127.0.0.1:9999/api/alert/receiver'
EOF
# Please note that 127.0.0.1:9999 is ${BMP_PRONOEA_HOST}:${BMP_PRONOEA_PORT}, and it can be modified according to the actual situation.

### Step 7: Create the Alertmanager service file
sudo cat << EOF > /etc/systemd/system/alertmanager.service
[Unit]
Description=Alertmanager
Wants=network-online.target

[Service]
User=alertmanager
Group=alertmanager
ExecStart=/usr/local/alertmanager/alertmanager \
    --config.file /etc/alertmanager/alertmanager.yml \
    --web.listen-address :9093
    --log.level=debug
Restart=always

[Install]
WantedBy=multi-user.target
EOF

# Step 8: Start the Alertmanager service
sudo systemctl daemon-reload
sudo systemctl start alertmanager

## Step 9: Check the Alertmanager service status
sudo systemctl status alertmanager

```
## Component Joint Debugging and Deployment <a id="3"></a>

In the previous steps, the non-container deployment commands for individual components were introduced respectively. This section describes how to start each component in sequence for joint debugging.

### Step 1: Environment Preparation <a id="3.1"></a>
**Environment preparation includes two parts: *network environment preparation* and *server preparation*. For specific steps, please refer to the [Environment Preparation](env-prepare.md) section.**

### Step 2: [Set Environment Variables](#2.1.2) <a id="3.2"></a>

### Step 3: Start Components in Sequence <a id="3.3"></a>
Please start each component in the following order:

1. Start the underlying middleware and initialize it.
    - [bmp-db](#2.2.3)
    - [bmp-redis](#2.2.4)
    - [bmp-mq](#2.2.5)

2. Start [bmp-scheduler](#2.4.1).

3. Start [bmp-openapi and bmp-openapi-console](#2.4.1).

4. Start the back-end related components.
    - [bmp-image](#2.2.1)
    - [bmp-tftp](#2.2.2)
    - [bmp-rsyslog](#2.2.6)
    - [Other back-end components](#2.4.1)

| Application Component | Category | Language/Component |
|-------------------|------| --- |
| bmp - console - api | Back-end | go |
| bmp - operation - api | Back-end | go |
| bmp - driver | Back-end | go |
| bmp - oob - alert | Back-end | go |
| bmp - oob - agent | Back-end | go |
| bmp - pronoea | Back-end | go |
| bmp - monitor - proxy | Back-end | go |
| bmp - dhcp - agent | Back-end | python |

5. Start the monitoring components.
   **Please note that the order cannot be changed.**
    - [bmp-pushgateway](#2.5.1)
    - [bmp-alertmanager](#2.5.3)
    - [bmp-prometheus](#2.5.2)
    - [bmp-pronoea](#2.4.1)

6. Start the front-end components.

| Application Component | Category | Language/Component |
|-------------------|------| --- |
| bmp - console - web | Front-end | vue |
| bmp - operation - web | Front-end | vue |

### Step 4: Access BMP <a id="3.4"></a>
**Note**: The default account is "admin" and the password is "df9Ydfd$c".
The console can be accessed at http://manager_ip:8080. In this example, it is http://192.168.14.80:8080.
The operation platform can be accessed at http://manager_ip:8081. In this example, it is http://192.168.14.80:8081.

### Step 5: Manage Bare Metal Servers in the Management Platform <a id="3.5"></a>
Log in to the <u>bmp operation platform</u>, enter the <u>device management</u> interface, and import the bmp node information. 