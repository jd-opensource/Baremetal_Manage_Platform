#!/bin/bash
set -e
mkdir -p /var/lib/prometheus/conf/rules
mkdir -p /var/lib/prometheus/data
eval "echo \"$(cat /prometheus.tpl)\"" > /var/lib/prometheus/conf/prometheus.yml
exec /bin/prometheus --config.file=/var/lib/prometheus/conf/prometheus.yml \
        --storage.tsdb.path=/var/lib/prometheus/data \
        --web.console.libraries=/usr/share/prometheus/console_libraries \
        --web.console.templates=/usr/share/prometheus/consoles \
        --web.enable-lifecycle