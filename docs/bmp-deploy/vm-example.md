[BMP Installation Guide]-Deploying a Physical BM Node with Virtual Machines

## Deployment Description

• This section describes how to simulate a physical server with a virtual machine and use it as a BM node接入bmp management.

• In this example,KVM and VirtualBMC are deployed on the same server.You can choose the BMP manager node as the deployment server,which means one server can act as both the manager node and the BM node.With just one server,you can complete the deployment and installation testing of BMP.

## Network Architecture


Steps to Create a Virtualized Physical Server

• Virtual machine network address planning
CIDR:192.168.100.0/24
Gateway:192.168.100.1
Virtual machine IP:192.168.100.2
DHCP relay server IP:192.168.100.254


• Install the virtual machine suite

```
  yum install -y qemu-kvm libvirt virt-install
  systemctl start libvirtd
  ```



• Create a virtual bridge and configure IP,this IP will act as the gateway for the virtual BM node,see network address planning

```
  ip link add br-vbm type bridge
  ip link set br-vbm up
  ip address add 192.168.100.1/24 dev br-vbm
  ```



• Firewall settings

```
  iptables -I FORWARD -i br-vbm  -j ACCEPT
  iptables -I FORWARD -o br-vbm  -j ACCEPT
  iptables -t nat -I POSTROUTING -s 192.168.100.0/24 ! -o br-vbm -j MASQUERADE
  ```



• Create a virtual machine and connect it to the virtual bridge

```
  #Create a virtual machine with configuration: 2c4G100G, SN: vmsn123456, connected to virtual bridge: br-vbm
  virt-install \
      --name vbm-1 \
      --memory 4096 \
      --vcpus=2 \
      --hvm \
      --boot hd \
      --network bridge:br-vbm \
      --disk /var/lib/libvirt/images/vbm-1.qcow2,format=qcow2,size=15 \
      --noreboot \
      --graphics vnc,listen=0.0.0.0 \
      --sysinfo type=smbios,system_serial=vmsn123456
  ```



• Install VirtualBMC.VirtualBMC official documentation:[VirtualBMC Documentation]()

```
  yum install -y python3 python3-devel libvirt-devel rust cargo                 #virtualbmc dependencies include python3
  python3 -m pip -U pip
  python3 -m pip install setuptools_rust
  python3 -m pip install virtualbmc
  vbmcd
  ```



• Create a VirtualBMC instance and associate it with the previously created virtual machine

```
  vbmc add vbm-1 --port 623      #vbm-1 represents the virtual machine name, 623 represents the port that the virtual BMC listens on
  vbmc start vbm-1
  vbmc list                   #Check VirtualBMC instances
  ```



• Test if the VirtualBMC instance is effective

```
  #The default user for VirtualBMC instances is admin, and the password is password
  ipmitool -I lanplus -H 127.0.0.1 -U admin -P password -p 623 power status
  ```



• Configure DHCP relay.Relay the DHCP requests sent by the virtual machine to the manager node.Since both the manager node and the libvirt service listen to the DHCP service port(udp/67),to prevent conflicts,run the DHCP relay program in an independent network namespace

```
  #Install DHCP relay
  yum install -y dhcp

  #Add network namespace and virtual network card
  ip link add veth0 type veth peer name veth1
  ip link set veth0 up
  ip link set veth1 up
  ip link set dev veth0 master br-vbm
  ip netns add ns1
  ip link set veth1 netns ns1

  #Set IP address for DHCP relay server, must be in the same subnet as the virtual machine
  ip netns exec ns1 ip address add 192.168.100.254/24 dev veth1 
  ip netns exec ns1 ip route add default via 192.168.100.1

  #Start the relay server, relay address is the manager node's management IP
  ip netns exec ns1 dhcrelay 10.208.12.72
  ```



• If the manager node and the virtual BM node are the same server,you need to enable the TFTP NAT extension function,otherwise,the virtual BM node may not be able to access bmp-tftp normally.

```
  echo 0 > /proc/sys/net/netfilter/nf_conntrack_helper
  iptables -t raw -A PREROUTING -p udp --dport 69 -j CT --helper tftp
  modprobe nf_nat_tftp
  ```



## Importing Devices into BMP Operation Platform

```
Out-of-band IP: VirtualBMC server IP  
Out-of-band port: VirtualBMC instance port, in this example, it is 6230  
Out-of-band login account: VirtualBMC instance user, in this example, it is admin  
Out-of-band login password: VirtualBMC instance password, in this example, it is password  
Subnet mask: Virtual machine network mask, in this example, it is 255.255.255.0  
Gateway: Virtual machine network gateway, in this example, it is 192.168.100.1  
Internal IPv4: Virtual machine IP, choose any available address from the CIDR, in this example, it is 192.168.100.2  
MAC1: Virtual machine network card MAC address, can be obtained through virsh domiflist VM_NAME  
SN: Virtual machine SN, in this example, it is vmsn123456
```