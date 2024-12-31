# [BMP Installation Guide](main.md) - Overview
## Glossary
* Manager node: The BMP management server, which runs all BMP components (except for BMP-agent)
* BM node: A bare metal server, a physical server used for normal operation, without an operating system before installation, running LiveOS during the installation process, and running GuestOS after installation
* GuestOS: The normal working operating system
* LiveOS: An in-memory operating system, which has BMP-agent pre-installed
* Out-of-band network card: A special network card on the physical server used for communication with BMC, also known as the IPMI network card
* Management network card: The standard network card on the physical server, located in the management network
* Management network: A 3-layer network, where the manager node's management network card communicates with the BM node's management network card through the management network
## Deployment Environment Requirements
### Manager Node
* Resource requirements
    * CPU: 4 cores
    * Memory: 8 GB
    * Disk: 40 GB
* Operating system: CentOS 7.3+
* Network requirements
    * Must be able to access the public network
    * Needs to open ports accessible by client browsers: 8080, 8081
    * Needs to open ports accessible by the management network: 10000, 5672, 67/udp, 69/udp
### BM Node
* Resource requirements
    * CPU: 1 core
    * Memory: 2 GB
    * Disk: 40 GB
    * Must have one out-of-band network card and one management network card


### Previous section [BMP Installation Guide](main.md)
### Next section [Deployment Architecture](deploy-architecture.md)