#!/usr/bin/env bash
USAGE="Usage: $0 -i input_image_path -o output_image_path -p package_path 

Build vmlinuz and initramfs.gz from imaget.tar.gz

Options:
    -a    Architecture
    -d  distro    Operating system distro and version
    -i    Image path
    -o    Output build directory
    -p    Base directory of packages
    -h    Display this help and exit
    -v    Turn on verbose messages for debugging
    -V    Image version
"
while getopts "V:a:d:i:o:p:h:v" OPTION; do
    case $OPTION in
    V) VERSION=$OPTARG ;;
    a) ARCH=$OPTARG ;;
    d) DISTRO=$OPTARG ;;
    i) IMAGE=$OPTARG ;;
    o) OUTPUT_BUILD_DIR=$OPTARG ;;
    p) BASE_DIR=$OPTARG ;;
    h) echo "$USAGE" && exit 0 ;;
    v) set -x && export VERBOSE=1 ;;
    *) echo "$USAGE" && exit 1 ;;
    esac
done

set -e

IMAGE_TMP=$(mktemp -d -t image-XXXXXXX)
INITRAMFS_DIR=$(mktemp -d -t initramfs-XXXXXXX)
ROOTFS_DIR=$(mktemp -d -t rootfs-XXXXXXX)
trap "rm -rf $IMAGE_TMP && rm -rf $INITRAMFS_DIR && rm -rf $ROOTFS_DIR" EXIT
echo $IMAGE_TMP
xz -d -c -f $IMAGE | tar --warning=no-timestamp -x -f  - -C $IMAGE_TMP
ls $IMAGE_TMP

KERNEL=$(ls -al $IMAGE_TMP/boot/vmlinuz-* | awk {'print $9'} | sort -V | grep -v rescue | head -1)
KERNEL_VERSION=$(echo $KERNEL | sed  "s:$IMAGE_TMP/boot/vmlinuz-::g")
echo "Kernel version: $KERNEL_VERSION"

CURRENT_DATE=`date +%Y%m%d%H`
if [ -n "${VERSION}" ]; then
  VMLINUZ_NAME="${VERSION}-${DISTRO}-${CURRENT_DATE}-vmlinuz"
  INITRAMFS_NAME="${VERSION}-${DISTRO}-${CURRENT_DATE}-initramfs.gz"
else
  VMLINUZ_NAME="${DISTRO}-${CURRENT_DATE}-vmlinuz"
  INITRAMFS_NAME="${DISTRO}-${CURRENT_DATE}-initramfs.gz"
fi

# linux kernel
/bin/cp $IMAGE_TMP/boot/vmlinuz-$KERNEL_VERSION $OUTPUT_BUILD_DIR/$VMLINUZ_NAME
chown -R root:root $OUTPUT_BUILD_DIR/$VMLINUZ_NAME
chmod 755 $OUTPUT_BUILD_DIR/$VMLINUZ_NAME

# busybox
mkdir -p $INITRAMFS_DIR/bin
/bin/cp $BASE_DIR/$ARCH/initramfs/busybox $INITRAMFS_DIR/bin/
chmod +x $INITRAMFS_DIR/bin/busybox

# init
/bin/cp $BASE_DIR/$ARCH/initramfs/init $INITRAMFS_DIR/
chmod +x $INITRAMFS_DIR/init

# rootfs
/bin/cp $IMAGE $INITRAMFS_DIR/rootfs.tar.xz

# initramfs
cd $INITRAMFS_DIR
find . -print0 | cpio --null -ov --format=newc | gzip -9 > $OUTPUT_BUILD_DIR/$INITRAMFS_NAME
chown -R root:root $OUTPUT_BUILD_DIR/$INITRAMFS_NAME
chmod 644 $OUTPUT_BUILD_DIR/$INITRAMFS_NAME
