#!/usr/bin/env bash

USAGE="Usage: $0 packdir

Before image archive to add operation

Options:
    -p    Plugin file
    -h    Display this help and exit
    -v    Turn on verbose messages for debugging
"
while getopts "p:hv" OPTION; do
    case $OPTION in
    p) PACKDIR=$OPTARG ;;
    h) echo "$USAGE" && exit 0 ;;
    v) set -x && export VERBOSE=1 ;;
    *) echo "$USAGE" && exit 1 ;;
    esac
done

set -e

# Check packdir is / or empty, avoid operate host directory
PACKDIR=`echo $PACKDIR | sed 's/ //g'`
if [[ -f "$PACKDIR" || "$PACKDIR" =~ ^/+$ ]];then
    echo "packdir can not use / or empty"
    exit 1
fi

# Dockerfile can not set image hostname persistent
echo "bmp_agent" > $PACKDIR/etc/hostname