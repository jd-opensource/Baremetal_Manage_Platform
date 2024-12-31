# 概念

## 组件说明
* bmp-console-web 控制台前端页面。使用vue3脚手架搭建。包含登录页、项目管理页、个人中心页、实例管理页等页面功能。
* bmp-console-api 控制台api。go语言后端业务，调用openapi接口实现控制台相关接口。为bmp-console-web拼装各类需要的数据，需要进行多语言适配。
* bmp-operation-web 运营平台前端页面。使用vue3脚手架搭建，包含登录页、机房管理页、机型管理页、镜像管理页、设备管理页、角色管理页、用户管理页等页面功能。
* bmp-operation-api 运营平台api。go语言后端业务，调用openapi接口实现运营平台相关接口。为bmp-operation-web拼装各类需要的数据，需要进行多语言适配。
* bmp-openapi bmp-openapi是bmp的核心模块。实现了满足swagger2.0规范的restful api格式接口。对外提供了bmp所有基础功能。对内进行数据库操作、调用bmp-scheduler完成实例生命周期管理的相关操作。
* bmp-scheduler 装机调度模块。承接bmp-openapi的实例生命周期管理请求，将上层请求转换为对应的commands，并且驱动commands执行，协同下层的bmp-driver，bmp-agent共同完成装机、重装、开机、关机等功能。
* bmp-driver 单机房应用，在有多个机房情况下需要部署多套bmp-driver服务，接收mq，对本机房的服务器进行开机、关机、重启、设置pxe启动等操作。
* bmp-dhcp-agent 单机房应用，装机之前需要提前更新dhcp配置，将Mac-IP关联关系存入dhcp配置中。之后liveOS才能从dhcp中获取到ip地址。
* bmp-db 数据库
* bmp-redis redis缓存
* bmp-mq 消息中间件
* bmp-tftp tftp服务器，存储着pxe引导需要的相关文件，包括pxe引导程序，pxe引导配置,LiveOS的内核和initramfs。
* bmp-image http服务器，存储着GuestOS镜像

## 装机流程
![架构图](./bmp-deploy/picture/bm-deploy.png)
## 装机流程说明
1. 客户端（bmp-console-web）发起装机请求，bmp-console-api接收请求  
2. bmp-console-api检查请求参数，通过后将请求转发给bmp-openapi  
3. bmp-openapi执行权限检查等操作，生成装机参数，发送给bmp-scheduler  
4. bmp-scheduler调度装机任务，生成一系列装机指令，通过bmp-mq服务发送给bmp-dhcp-agent、bmp-driver和bmp-agent  
5. bmp-dhcp-agent收到指令，设置内置的dhcp服务器，使得bm节点在pxe启动阶段能够获取正确的ip配置和tftp地址(bmp-tftp地址)  
6. bmp-driver收到指令后，将bm节点设置为pxe启动并重启  
7. bm节点执行pxe启动，内置在网卡内的PXEClient启动，发送dhcp请求广播，bmp-dhcp-agent内置的dhcp服务器收到dhcp请求后，响应相应的ip配置和tftp地址  
8. PXEClient配置自身ip，之后从bmp-tftp下载pxe引导程序并执行，pxe引导程序继续从bmp-tftp获取其它引导参数，下载内核和initramfs，启动内存操作系统，内置在内存操作系统中bmp-agent服务开始启动  
9. bmp-agent接收指令，执行后续bm装机操作，如：设置raid，分区等  
10. bmp-agent从bmp-image下载客户操作系统镜像文件，写到bm节点磁盘中，然后初始化客户操作系统  
11. bmp-agent执行重启，完成操作系统安装  

## 名词解释
* manager节点: bmp管理服务器，运行着所有bmp组件（除了bmp-agent）
* bm节点：裸金属服务器，正常工作用的物理服务器，未装机前无操作系统，装机阶段运行LiveOS，装机完成后运行GuestOS
* GuestOS: 正常工作的操作系统
* LiveOS: 内存操作系统，里面预装了bmp-agent
* 带外网卡：物理服务器上用于与bmc通信的特殊网卡，又称ipmi网卡
* 管理网卡：物理服务器的标准的网卡，处在管理网络中
* 管理网络: 3层网络，manager节点的管理网卡与bm节点的管理网卡通过管理网络通信

