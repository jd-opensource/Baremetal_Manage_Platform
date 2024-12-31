# Concept

## Component Description
* bmp-console-web console front-end page. Use vue3 scaffolding to build. Contains login page, project management page, personal center page, instance management page and other page functions.
* bmp-console-api console api. Go language back-end business, call openapi interface to implement console related interfaces. Assemble all kinds of required data for bmp-console-web, need to be adapted to multiple languages.
* bmp-operation-web operation platform front-end page. Use vue3 scaffolding to build, including login page, idc management page, device type management page, image management page, device management page, role management page, user management page and other page functions.
* bmp-operation-api operation platform api. Go language back-end business, call openapi interface to implement operation platform related interfaces. Assemble all kinds of required data for bmp-operation-web, need to be adapted to multiple languages.
* bmp-openapi bmp-openapi is the core module of bmp. Implements the RESTful API format interface that meets the swagger2.0 specification. It provides all basic functions of bmp externally. It performs database operations internally and calls bmp-scheduler to complete related operations of instance lifecycle management.
* bmp-scheduler is an installation scheduling module. It accepts instance lifecycle management requests from bmp-openapi, converts upper-level requests into corresponding commands, and drives the execution of commands, cooperating with lower-level bmp-driver and bmp-agent to complete installation, reinstallation, startup, shutdown and other functions.
* bmp-driver is a single-idc application. In the case of multiple idcs, multiple sets of bmp-driver services need to be deployed, receiving mq, and starting, shutting down, restarting, setting pxe startup and other operations on the servers in this IDC.
* bmp-dhcp-agent is a single-idc application. Before installation, the dhcp configuration needs to be updated in advance, and the Mac-IP association relationship needs to be stored in the dhcp configuration. Only then can liveOS obtain the ip address from dhcp.
* bmp-db database
* bmp-redis redis cache
* bmp-mq message middleware
* bmp-tftp tftp server, which stores the relevant files required for pxe boot, including pxe boot program, pxe boot configuration, LiveOS kernel and initramfs.
* bmp-image http server, storing GuestOS image

## Installation process
![Architecture diagram](./bmp-deploy/picture/bm-deploy.png)
## Installation process description
1. The client (bmp-console-web) initiates the installation request, and bmp-console-api receives the request
2. bmp-console-api checks the request parameters, and forwards the request to bmp-openapi if it passes
3. bmp-openapi performs operations such as permission check, generates installation parameters, and sends them to bmp-scheduler
4. bmp-scheduler schedules the installation task, generates a series of installation instructions, and sends them to bmp-dhcp-agent, bmp-driver and bmp-agent through the bmp-mq service
5. bmp-dhcp-agent receives the instruction and sets the built-in dhcp server, so that the bm node can obtain the correct ip configuration and tftp address (bmp-tftp address) during the pxe startup phase
6. After receiving the instruction, bmp-driver sets the bm node to pxe startup and restarts
7. The bm node executes pxe startup, the PXEClient built into the network card starts, sends a dhcp request broadcast, and the dhcp server built into bmp-dhcp-agent responds to the corresponding ip configuration and tftp address after receiving the dhcp request
8. PXEClient configures its own ip, then downloads the pxe boot program from bmp-tftp and executes it. The pxe boot program continues to obtain other boot parameters from bmp-tftp, downloads the kernel and initramfs, starts the memory operating system, and the bmp-agent service built into the memory operating system starts to start
9. bmp-agent receives instructions and executes subsequent bm installation operations, such as: setting up raid, partitioning, etc.
10. bmp-agent downloads the client operating system image file from bmp-image, writes it to the bm node disk, and then initializes the client operating system
11. bmp-agent executes a restart to complete the operating system installation

## Glossary
* manager node: bmp management server, running all bmp components (except bmp-agent)
* bm node: bare metal server, physical server for normal operation, no operating system before installation, LiveOS during installation, and GuestOS after installation
* GuestOS: normal operating operating system
* LiveOS: memory operating system with bmp-agent pre-installed
* out-of-band network card: special network card on physical server for communication with bmc, also known as ipmi network card
* management network card: standard network card of physical server, in management network
* management network: layer 3 network, management network card of manager node communicates with management network card of bm node through management network