#!/bin/bash
export PATH=$PATH:/usr/local/bin:/usr/bin:/usr/local/sbin:/usr/sbin:/export/server/php/sbin:/export/server/go/bin:/Doorgod/users/.local/bin:/Doorgod/users/bin
export PATH=/export/servers/go/bin/:$PATH

BIN_PATH=/export/servers/cps-agent-proxy
mkdir -p $BIN_PATH/log
APP_NAME=cps-agent-proxy
echo "Start $APP_NAME ..."

if [[ ! -d "$BIN_PATH" ]];then
    mkdir -p ${BIN_PATH}
fi

NUM=`ps aux | grep -E ${APP_NAME} | grep -v grep |wc -l`
if [[ "${NUM}" -ge "1" ]];then
    echo "cps-agent-proxy Process is running, shuting down first... "
    PID=`ps aux | grep -E ${APP_NAME} | grep -v grep | awk -F " " '{print $2}'`
    kill ${PID}
fi
echo "Start $APP_NAME again ..."
nohup ${BIN_PATH}/${APP_NAME} -c ${BIN_PATH}/conf/proxy.ini > ${BIN_PATH}/log/proxy.log 2>&1 &
