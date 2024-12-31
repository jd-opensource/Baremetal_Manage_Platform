# [BMP Installation Guide](main.md) - BM Node Deployment Process
The BMP BM node deployment process is shown in the following diagram:


![Deployment Process](picture/bm-deploy.png)
BMP BM Node Deployment Process Diagram

## Deployment Process Explanation
1. The client (bmp-console-web) initiates a deployment request, which is received by bmp-console-api.
2. bmp-console-api checks the request parameters, and if they pass, forwards the request to bmp-openapi.
3. bmp-openapi performs permission checks and other operations, generates deployment parameters, and sends them to bmp-scheduler.
4. bmp-scheduler schedules the deployment task, generates a series of deployment instructions, and sends them to bmp-dhcp-agent, bmp-driver, and bmp-agent through the bmp-mq service.
5. bmp-dhcp-agent receives the instructions, sets up its built-in DHCP server so that the BM node can obtain the correct IP configuration and TFTP address (bmp-tftp address) during the PXE boot phase.
6. bmp-driver receives the instructions, sets the BM node to PXE boot, and restarts it.
7. The BM node performs a PXE boot, and the PXEClient built into the network card starts up, sending a DHCP request broadcast. The built-in DHCP server in bmp-dhcp-agent receives the DHCP request and responds with the corresponding IP configuration and TFTP address.
8. The PXEClient configures its own IP, and then downloads the PXE boot program and executes it. The PXE boot program continues to download other boot parameters from bmp-tftp, downloads the kernel and initramfs, starts the memory operating system, and the bmp-agent service built into the memory operating system begins to start up.
9. bmp-agent receives the instructions and performs subsequent BM node deployment operations, such as setting up RAID and partitioning.
10. bmp-agent downloads the customer operating system image file from bmp-image, writes it to the BM node's disk, and then initializes the customer operating system.
11. bmp-agent executes a restart, completing the operating system installation.

### Previous section [Service Management](service-manage.md)
### Next section [Troubleshooting](troubleshoot.md)