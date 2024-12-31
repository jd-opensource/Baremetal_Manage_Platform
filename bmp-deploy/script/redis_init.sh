#!/bin/sh
set -e
if [ ! -e /data/redis.conf ] && [ -e /redis.conf.tpl ]; then
    eval "echo \"$(cat /redis.conf.tpl)\"" > /data/redis.conf
fi
redis-server /data/redis.conf