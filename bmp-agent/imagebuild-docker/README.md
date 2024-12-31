# Building RamOS
English | [简体中文](README.zh-CN.md) 

This project will build the vmlinu and initramfs.gz files required for pxe to boot RamOS, and build the agent service into RamOS.

## Dependency
1. Docker 20.10.19+
2. JQ

## Build Process
1. Enter the current directory imagebuild-docker
2. Building RamOS

Execute the following command to build RamOS based on CentOS 7.9:
```
make install 
```
After the build is complete, the generated vmlinuz and initramfs.gz files will be located in the ./build/centos_7_9 directory.

3. Building other versions of RamOS:
See [Makefile](./Makefile) for build options for more versions.
