# One-click deployment of BMP system
- [One-click deployment overview](#1)
- [Term explanation](#1.1)
- [Overall deployment steps](#1.2)
- [Environment preparation](#2)
- [Management node deployment](#3)
- [Environment preparation](#3.1)
- [Configure environment parameters](#3.2)
- [One-click deployment of BMP](#3.3)
- [Access BMP](#3.4)
- [BMP platform manages physical servers](#3.5)
- [Virtual machine simulation deployment](#4)
## One-click deployment overview<a id="1"></a>
### Term explanation<a id="1.1"></a>
- BMP
Bare Metal Manage Platform, referred to as BMP.
- manager node
BMP management server (management node), running all BMP components (except bmp-agent)
- bm node
Installation node, a physical server for normal work, without operating system before installation, running LiveOS during installation, and running GuestOS after installation.
- GuestOS
A normal operating system
- LiveOS
A memory operating system with bmp-agent pre-installed
- Out-of-band network card
A special network card on a physical server used to communicate with bmc, also known as an ipmi network card
- Management network card
A standard network card on a physical server, located in the management network

### Overall deployment steps<a id="1.2"></a>
- The manager node and the bm node are in a large layer 2 network (if you only experience the BMP platform and do not intend to experience the BMP bare metal installation process, this step can be skipped)
- Install the operating system CentOS 7.9 on the manager node
- Configure the dhcp-relay address of the bm node tor to the ip address of the manager node (if you only experience the BMP platform and do not intend to experience the BMP bare metal installation process, this step can be skipped)
- Deploy the BMP platform through docker on the manager node. Refer to [Management node deployment](#3)

## Environment preparation<a id="2"></a>
**Environment preparation includes *network environment preparation* and *Server preparation* Two parts, the overall architecture is as follows**

![bmp-framework_en.png](bmp-deploy/picture/bmp-framework_en.png)

### 1. Network preparation
*Please note: If you only experience the BMP management component, the following network environment configuration can be skipped; if you are preparing for real installation, the following network configuration must be performed*
The manager node and the bm node are in a large Layer 2 network, and the bm node needs to obtain its IP address based on the DHCP service.
Before deploying the BMP platform, you need to configure the dhcp-relay of the upstream switch of the installation node to the manager node.
The following shows an example network configuration

![network-deploy.png](bmp-deploy/picture/network-deploy.png)
In this network structure, TOR1 represents the upstream switch of the installation node eth0 network card.
- manager node
eth0 ip: 192.168.14.80/27
This address is usually a private network address, and the DHCP service of BMP will be started on this private network address. This private network address will also be notified to the BMP platform through the configuration file, see [Configure environment parameters](#3.2) for details
- bm node
ipmi ip: 192.168.11.130/27
- TOR1
ip: 192.168.12.129/27
vlan id: 200
- Log in to TOR1, check the dhcp relay of vlan 200, and check whether the relay address is the manager node management network card ip (in this example, 192.168.14.80)
Log in to TOR1: ssh switchuser@192.168.12.129, enter the password.
H3C switch command: dis cu int Vlan-interface 200
~~~
Here are the steps to reconfigure DHCP relay using H3C switch:

#Enter system view:
system-view
#Enter interface view, for example VLAN interface 200:
interface vlan 200
#Configure the server address of DHCP relay, i.e. management node ip:
dhcp relay server 192.168.14.80
#Save configuration:
save f
#Return to system view:
quit
#Exit configuration mode:
quit

Please note that the specific commands and parameters may vary depending on the switch model, software version, and configuration. Make sure you understand the specific configuration requirements of your switch.
~~~
#### The BMP standard network communication process is shown in the figure:
![network.png](bmp-deploy/picture/network.png)
* The out-of-band network card ipmi of the physical machine, the management network card eth0 and the management network card eth0 of the manager node are all in the same Layer 2 management network.
* Each BMP component in the manager node runs as a container, and the service port is exposed to the management network access through docker-proxy and the eth0 network card.
* The built-in BMC management system of the bm node is exposed to the management network access through the ipmi network card
* Other services running on the bm node are exposed to the management network access through the eth0 network card
* In order to ensure that the system running on the bm node can obtain the IP configuration from the built-in DHCP server of the bmp-dhcp-agent container of the manager node, the upstream switch of the eth0 network card of the bm node may need to be configured with a DHCP relay, and the relay address is the management IP of the manager node.
* At the beginning of the installation, the bmp-driver accesses the out-of-band IP of the bm node through the management network to set the boot mode of the physical machine and control the power on and off.
* During the pxe boot phase, the PXEClient built into the eth0 network card of the bm node accesses the dhcp server built into the bmp-dhcp-agent through the management network, obtains the ip configuration and the pxe boot program download address, sets the eth0 ip, and downloads the LiveOS kernel and initramfs from bmp-tftp, and then starts LiveOS.
* When LiveOS starts, it also obtains the ip configuration from the dhcp server built into the bmp-dhcp-agent and sets the eth0 ip.
* During the GuesOS installation phase, bmp-agent can obtain the required installation information from bmp-mq, or download the GuestOS image from bmp-tftp
* Regardless of whether PXEClient, LiveOS or GustOS is running, the ip of the management network card eth0 of the bm node is the same, and the specific ip address is set in <u>Operation Platform-Device Management</u>. Please assign the correct ip address to the bm node to ensure that the system running on the bm node at each stage can work properly.

### 2. Server preparation
manager node:
~~~
Resource requirements
cpu: 4 cores
Memory: 8 GB
Disk: 40 GB
Operating system: CentOS 7.9
Network requirements
Access to the public network
Ports that need to be opened to client browsers: 8080, 8081
Ports that need to be opened to management network access: 10000, 5672, 67/udp, 69/udp
~~~
Minimum configuration of the installed node
~~~
cpu: 1 core
Memory: 2 GB
Disk: 40 GB
Has an out-of-band network card and a management network card
~~~

## Management node deployment<a id="3"></a>
~~~
Step 1: Environment preparation

Step 2: Configure environment parameters

Step 3: One-click deployment of BMP

Step 4: Access BMP

Step 5: BMP platform manages physical servers
~~~

### Step 1: Environment preparation<a id="3"></a> id="3.1"></a>

#### 1. Install docker and docker compose
- yum installation method (CentOS 7.9 system)
~~~
#You can use domestic mirror sources, such as Tsinghua mirror sources
yum install -y yum-utils
yum-config-manager \
--add-repo \
https://mirrors.tuna.tsinghua.edu.cn/docker-ce/linux/centos/docker-ce.repo
sed -i 's+https://download.docker.com+https://mirrors.tuna.tsinghua.edu.cn/docker-ce+' /etc/yum.repos.d/docker-ce.repo

yum install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

systemctl start docker
~~~

For more installation methods, please refer to the docker official website
- Docker official document address: https://docs.docker.com/
- Docker Compose official document address: https://docs.docker.com/compose/

### 2. Download BMP source code to the manager node server
Before downloading, first clarify the deployment directory of BMP. This instance is installed in the user's home directory by default. First
~~~
# Jump to the home directory, and then execute the subsequent download command
cd ~
~~~
Find the [Code] button in the upper right corner of the homepage and click the drop-down menu, as shown in the figure
![github.png](bmp-deploy/picture/github.png)
- Method 1: Use the git clone command
~~~
# Execute on the management node, clone the HTTPS link copied above, for example:
git clone https://github.com/jd-opensource/Baremetal_Manage_Platform.git
~~~
- Method 2: Click Download ZIP, download the zip package directly, and manually transfer it to the management server
For more github operations, refer to the official github operation document:
https://docs.github.com/en/repositories/creating-and-managing-repositories/cloning-a-repository
- Method 3: You can download the installation package directly from JD Cloud
~~~
wget https://bmp.s3.cn-north-1.jdcloud-oss.com/opensource/Baremetal_Manage_Platform-33040a7.zip
unzip Baremetal_Manage_Platform-33040a7.zip -d Baremetal_Manage_Platform
~~~

### Step 2: Configure environment parameters <a id="3.2"></a>
- Deployment script directory
The deployment script is located in the bmp-deploy directory
File description in the directory: [bmp-deploy file and directory description] (bmp-deploy/bmp-deploy.md)

- Modify configuration
Edit .env file
- Modify BMP_HOST_IP to the manager node management network card IP. This example is 192.168.14.80. If the server has multiple network cards, please configure one of the IPs. If bond is configured, use the bond network card. Otherwise, try to choose a network card with a private network IP address.
ifconfig command-bond0 example image
![ifconfig-bond0.png](bmp-deploy/picture/ifconfig-bond0.png)
ifconfig command-eth0 example image
![ifconfig-eth0.png](bmp-deploy/picture/ifconfig-eth0.png)
- Modify BASE_REGISTRY. Deployment requires the use of basic container images such as mysql/redis/mq/nginx. This example pulls from quay.io by default. If the network is restricted, please modify it to other public or private image repositories

~~~
cd ~/Baremetal_Manage_Platform/bmp-deploy/
# Check the server IP address, first check the IP address of bond0, and then check the IP address of eth0. If you are not sure, please consult the operation and maintenance personnel.
ifconfig

# Edit environment variables
# Must be consistent with the manager node server IP
# Set BMP_HOST_IP=192.168.14.80
vi .env
~~~

For more configurations, please refer to [Configuration Management](bmp-deploy/config.md)

### Step 3: One-click deployment of BMP<a id="3.3"></a>
- Install BMP
~~~
cd ~/Baremetal_Manage_Platform/bmp-deploy/
bash bmp-deploy-precheck.sh -e public
bash bmp-deploy.sh start all
~~~

- Check the status of each service
~~~
bash bmp-deploy.sh status all
~~~
Under normal circumstances, the status of all services is running. If some services are not running, refer to [Troubleshooting](bmp-deploy/troubleshoot.md)
During the installation process, the relevant operating system image will be automatically downloaded from JD Object Storage. The download speed is related to the server network speed. Please wait patiently.
For more service management operations, refer to [Service Management](bmp-deploy/service-manage.md)

### Step 4: Access BMP<a id="3.4"></a>
Note: Default account: admin, password: df9Ydfd$c

Console: http://manager_ip:8080. This example is http://192.168.14.80:8080.

Operation platform: http://manager_ip:8081. This example is http://192.168.14.80:8081.

### Step 5: Manage bare metal servers on the management platform<a id="3.5"></a>
Log in to the <u>BMP operation platform</u>, enter the <u>device management</u> interface, and import the bm node information

## Virtual machine simulation deployment<a id="4"></a>
* Intranet IPv4, network subnet mask and gateway Please fill in the correct address so that the PXEClient, LiveOS, and GuestOS running on the bm node can access the manager network normally.
If you want to use a virtual machine to simulate a bm node, see [Virtual machine simulation bm node](bmp-deploy/vm-example.md)