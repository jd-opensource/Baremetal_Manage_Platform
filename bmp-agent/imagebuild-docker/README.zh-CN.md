# 构建RamOS
[English](README.md) | 简体中文

此项目将会构建出pxe引导RamOS所需要的vmlinu和initramfs.gz文件，并将agent服务内置到RamOS中。

## 依赖
1. Docker 20.10.19及以上
2. JQ

## 构建过程
1. 进入到当前目录 imagebuild-docker
2. 构建RamOS

执行以下命令构建基于 CentOS 7.9 的 RamOS：
```
make install 
```
构建完成后，生成的 vmlinuz 和 initramfs.gz 文件将位于 ./build/centos_7_9 目录下。

3. 构建其他版本的 RamOS：
请参阅[Makefile](./Makefile)获取更多版本的构建选项
