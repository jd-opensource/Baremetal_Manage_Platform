#!/bin/sh
set -e
eval "echo \"$(cat /alertmanager.tpl)\"" > /etc/alertmanager/alertmanager.yml
exec /bin/alertmanager --config.file=/etc/alertmanager/alertmanager.yml --storage.path=/alertmanager --log.level=debug