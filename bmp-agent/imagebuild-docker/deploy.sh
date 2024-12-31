#!/bin/bash

arch=$(uname -m)
DISTRO=${1:-centos_7_9}
SOURCE_DIR=./build/${DISTRO}
bmp_kernel="bmp_kernel_name_x86"
bmp_initramfs="bmp_initramfs_name_x86"
if [ "$arch" == "x86_64" ]; then
    :
elif [ "$arch" == "aarch64" ]; then
    SOURCE_DIR="${SOURCE_DIR}_arm64"
    bmp_kernel="bmp_kernel_name_arm64"
    bmp_initramfs="bmp_initramfs_name_arm64"
else
    echo "error"
    exit
fi
# file to 72
scp ${SOURCE_DIR}/* 10.208.12.72:/home/bmp/bmp-deploy/data/
scp ${SOURCE_DIR}/* 10.226.193.42:/home/bmp/bmp-deploy/data/

# config name
vmlinuz_name=$(find "${SOURCE_DIR}/" -type f -name "*-vmlinuz" -exec basename {} \;)
initramfs_name=$(find "${SOURCE_DIR}/" -type f -name "*-initramfs.gz" -exec basename {} \;)
echo "${bmp_kernel}=${vmlinuz_name}"
echo "${bmp_initramfs}=${initramfs_name}"