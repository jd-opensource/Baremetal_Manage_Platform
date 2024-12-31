#!/bin/sh
set -ue

[ -z "$BMP_DHCP_CONFIG_DIR" ] && echo "error! BMP_DHCP_CONFIG_DIR is unset!" && exit 1

if [ ! -e $BMP_DHCP_CONFIG_DIR/dhcpd.conf ] && [ -e /dhcpd.conf.tpl ]; then
    eval "echo \"$(cat /dhcpd.conf.tpl)\"" > $BMP_DHCP_CONFIG_DIR/dhcpd.conf
fi

LOG_FILE=/var/log/bmp/bmpda/dhcpd.log
exec_file=/usr/sbin/dhcpd
data_dir=$BMP_DHCP_CONFIG_DIR
dhcpd_conf="$data_dir/dhcpd.conf"
PID=""
if [ ! -r "$dhcpd_conf" ]; then
    echo "Please ensure '$dhcpd_conf' exists and is readable."
    echo "Run the container with arguments 'man dhcpd.conf' if you need help with creating the configuration."
    exit 1
fi

[ -e "$data_dir/dhcpd.leases" ] || touch "$data_dir/dhcpd.leases"


help(){
    echo "${0} <start|stop|restart|status>"
    exit 1
}

get_pid()
{
    local pid=""
    pid=$(ps -ef | grep "$exec_file" | grep -v 'grep' | awk '{print $1}')
    PID=$pid
}

checkhealth(){
    get_pid
    [ X"$PID" != X"" ] && echo "dhcpd is running" && return 0
    echo "dhcpd is not running" && return 1
}

start(){
    /usr/sbin/dhcpd -f -d --no-pid -cf "$data_dir/dhcpd.conf" -lf "$data_dir/dhcpd.leases" >> $LOG_FILE 2>&1 &
}

stop(){
    checkhealth || return 0
    [ X"$PID" != X"" ] && kill $PID
    sleep 2
    checkhealth || return 0
    return 1
}

case "${1}" in
    start)
        stop && start
        ;;
    stop)
        stop
        ;;
    status|health|checkhealth)
        checkhealth
        ;;
    restart)
        stop && start
        ;;
    *)
        help
        ;;
esac

