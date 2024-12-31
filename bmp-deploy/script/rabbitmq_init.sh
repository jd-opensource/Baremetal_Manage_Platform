#!/bin/sh
(sleep 10
rabbitmqctl add_vhost $BMP_MQ_VHOST
rabbitmqctl add_user $BMP_MQ_USER "$BMP_MQ_PASSWORD"
rabbitmqctl set_user_tags $BMP_MQ_USER administrator
rabbitmqctl set_permissions -p "$BMP_MQ_VHOST" $BMP_MQ_USER '.*' '.*' '.*'
rabbitmqctl list_users
rabbitmqctl change_password $BMP_MQ_USER "$BMP_MQ_PASSWORD")&
rabbitmq-server