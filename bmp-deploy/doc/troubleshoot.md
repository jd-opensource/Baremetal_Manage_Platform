# [BMP Installation Guide](main.md) - Troubleshooting
* PXE boot failure
1. Is the boot mode of the machine set to BIOS?
2. Is the out-of-band network of the machine correctly configured?
3. Is the BMC (Baseboard Management Controller) of the machine functioning normally?
4. Are the network cables of the machine connected correctly?
* A service is not running normally
~~~
./bmp-deploy.sh status all          # Check which services are not running normally
./bmp-deploy.sh start ServiceName   # Start the service
~~~

### Previous section [BM Node Deployment Process](bm-deploy.md)
### Next section [Practical Tools](tool.md)
