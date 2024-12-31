#!/bin/bash

set -e

log() {
    echo "`basename $0`: $@"
}

usage() {
    [[ -z "$1" ]] || echo -e "USAGE ERROR: $@\n"
    echo "`basename $0`: IMAGEFILE DEVICE"
    echo "  - This script images DEVICE with IMAGEFILE"
    exit 1
}

IMAGEFILE="$1"
DEVICE="$2"

[[ -f $IMAGEFILE ]] || usage "$IMAGEFILE (IMAGEFILE) is not a file"
[[ -b $DEVICE ]] || usage "$DEVICE (DEVICE) is not a block device"

# In production this will be replaced with secure erasing the drives
# For now we need to ensure there aren't any old (GPT) partitions on the drive
log "Erasing existing GPT and MBR data structures from ${DEVICE}"
sgdisk -Z $DEVICE || sgdisk -o $DEVICE

log "Imaging $IMAGEFILE to $DEVICE"

# limit the memory usage for qemu-img to 1 GiB
ulimit -v 1048576
qemu-img convert -t directsync -O host_device $IMAGEFILE $DEVICE
sync

log "${DEVICE} imaged successfully!"
