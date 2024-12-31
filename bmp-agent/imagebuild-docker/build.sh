#!/usr/bin/env bash
USAGE="Usage: $0

Build vmlinuz and initramfs.gz

Options:
    -a            Architecture
    -d  distro    Operating system distro and version
    -h            Display this help and exit
    -v            Turn on verbose messages for debugging
"

BUILD_SCRIPT_FILE=$0
BUILD_SCRIPT_ABSOLUTE_PATH=$(cd "$(dirname "$BUILD_SCRIPT_FILE")"; pwd)
BASE_DIR=$BUILD_SCRIPT_ABSOLUTE_PATH
VERBOSE_OPT=''
VERSION=$(cat "$BUILD_SCRIPT_ABSOLUTE_PATH/../bmpa/version" | tr -d '\n[:space:]')


while getopts "a:b:d:i:hv" OPTION; do
    case $OPTION in
    a) ARCH=$OPTARG ;;
    b) BASE_DIR=$OPTARG ;;
    d) DISTRO=$OPTARG ;;
    i) IMAGE=$OPTARG ;;
    h) echo "$USAGE" && exit 0 ;;
    v) set -x && export VERBOSE=1 && VERBOSE_OPT='-v';; 
    *) echo "$USAGE" && exit 1 ;;
    esac
done

set -e

DEST_BUILD_DIR=$BASE_DIR/build/$DISTRO
DOCKERFILE_PATH=$BASE_DIR/$DISTRO/Dockerfile
rm -rf $DEST_BUILD_DIR
mkdir -p $DEST_BUILD_DIR

$BASE_DIR/docker-build.sh $VERBOSE_OPT -t $IMAGE -f $DOCKERFILE_PATH -p ../

SAVE2_IMAGE_PLUGIN=$BASE_DIR/save2image_plugin.sh
if [[ -f "$BASE_DIR/$DISTRO/save2image_plugin.sh" ]];then
    SAVE2_IMAGE_PLUGIN=$BASE_DIR/$DISTRO/save2image_plugin.sh
fi

docker save $IMAGE | $BASE_DIR/save2image.sh $VERBOSE_OPT -p $SAVE2_IMAGE_PLUGIN -o ${DEST_BUILD_DIR}/image.tar.gz.tmp
mv "${DEST_BUILD_DIR}/image.tar.gz.tmp" "${DEST_BUILD_DIR}/image.tar.gz"
trap "rm -rf ${DEST_BUILD_DIR}/image.tar.gz" EXIT
$BASE_DIR/archive.sh $VERBOSE_OPT -i $DEST_BUILD_DIR/image.tar.gz  -p $BASE_DIR -o $DEST_BUILD_DIR -a $ARCH -V $VERSION -d $DISTRO

echo "Build image success"
