# Service startup failed
Restart a service
~~~
./bmp-deploy.sh status all #Check which services are not running properly
./bmp-deploy.sh start ServiceName #Start the service
~~~

# Reason for pxe boot failure
pxe boot failure
Is the machine's boot mode BIOS
Is the machine's out-of-band network correct
Is the machine's BMC normal
Is the machine's network port connection correct

# Docker is not started
Error: Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
After restarting the manager node, you need to restart the docker service
~~~
systemctl start docker.service
~~~