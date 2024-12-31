#!/usr/bin/env bash
USAGE="Usage: $0 -t image_tag -f dockerfile_path -p build_content_path

Build docker image

Options:
    -p             Build Content path   
    -t  tag list   Name and optionally a tag in the 'name:tag' format
    -f  file path  Name of the Dockerfile
    -h             Display this help and exit
    -v             Turn on verbose messages for debugging
"
while getopts "p:t:f:hv" OPTION; do
    case $OPTION in
    p) BUILD_CONTENT_PATH=$OPTARG ;;
    t) IMAGE=$OPTARG ;;
    f) DOCKERFILE_PATH=$OPTARG ;;
    h) echo "$USAGE" && exit 0 ;;
    v) set -x && export VERBOSE=1 ;;
    *) echo "$USAGE" && exit 1 ;;
    esac
done

set -e

docker build -t $IMAGE -f $DOCKERFILE_PATH $BUILD_CONTENT_PATH