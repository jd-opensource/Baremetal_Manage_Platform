#!/bin/bash

set -u
set -x
set -e

if [ $# -gt 0 ]; then
    EFI_BOOT_LABEL=$1
else
    EFI_BOOT_LABEL="bmp"
fi

if efibootmgr |grep -q "${EFI_BOOT_LABEL}";then
for num in `efibootmgr |grep "${EFI_BOOT_LABEL}"|awk -F'*' '{print $1}'|sed 's/Boot//g'`
do
    efibootmgr -B -b ${num}
done
fi

echo "Delete efi boot finished.!"
