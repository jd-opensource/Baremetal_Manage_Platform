#!/bin/bash
set -e
SCRIPTS_DIR=$(dirname $(realpath $0))
APP_NAME=$1

mkdir -p output/bin
cp $SCRIPTS_DIR/control output/bin/
#mkdir output/bmp-deploy
#find . -maxdepth 1 ! -regex '\.' ! -regex '\./output' ! -regex '\./\.git' ! -regex '\./tmp_build.*\.sh' | xargs -I {} cp -r {} output/bmp-deploy/
echo APP_NAME=$APP_NAME >> output/bin/app.properties
echo COMPILER_VERSION=$COMPILER_VERSION >> output/bin/app.properties
