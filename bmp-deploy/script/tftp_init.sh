#!/bin/bash
set -e

# legacy for x86_64
mkdir -p /var/tftpboot/pxelinux.cfg
eval "echo \"$(cat /pxelinux.tpl)\"" > /var/tftpboot/pxelinux.cfg/default

#uefi for x86_64
mkdir -p /var/tftpboot/uefi/x86_64
eval "echo \"$(cat /grub_x86_64.cfg.tpl)\"" > /var/tftpboot/uefi/x86_64/grub.cfg

#uefi for arm64
mkdir -p /var/tftpboot/uefi/arm64
eval "echo \"$(cat /grub_arm64.cfg.tpl)\"" > /var/tftpboot/uefi/arm64/grub.cfg

#uefi for loongarch64
mkdir -p /var/tftpboot/uefi/loongarch64
eval "echo \"$(cat /grub_loongarch64.cfg.tpl)\"" > /var/tftpboot/uefi/loongarch64/grub.cfg

exec in.tftpd -L --secure /var/tftpboot