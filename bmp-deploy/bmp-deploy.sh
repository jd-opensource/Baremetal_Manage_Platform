#!/bin/bash
set -ue
WORKSPACE=$(cd $(dirname $0); pwd)
set -a
. $WORKSPACE/config/service.cfg
set +a

#
## third-party service
 #bmp_service_thirdparty=(
 #      bmp-db
 #      bmp-redis
 #      bmp-mq
 #      bmp-image
 #      bmp-pushgateway
 #      bmp-alertmanager
 #      bmp-prometheus
 #)
 #
 ## app serveice
 #bmp_service_app=(
 #      bmp-console-web
 #      bmp-operation-web
 #      bmp-console-api
 #      bmp-operation-api
 #      bmp-openapi
 #      bmp-openapi-console
 #      bmp-scheduler
 #      bmp-driver
 #      bmp-dhcp-agent
 #      bmp-oob-alert
 #      bmp-oob-agent
 #      bmp-monitor-proxy
 #      bmp-pronoea
 #)
 #
 ## other service
 #bmp_service_other=(
 #      bmp-tftp
 #      bmp-rsyslog
 #)
 #
 #bmp_service_all=(
 #      ${bmp_service_thirdparty[@]}
 #      ${bmp_service_other[@]}
 #      ${bmp_service_app[@]}
 #)
#
#


SERVICES_ALL=(${bmp_service_all[@]})
COMMANDS_ALL=(
    start
    stop
    restart
    status
    images
    uninstall
    reinstall
    config
)
ALL_SERVICE=false
COMMAND=""
SERVICES=()
export TAG=""

# x86-64 arm
ARCH=$(uname -i)

# set docker compose cmdline
if docker compose version &> /dev/null;then
    COMPOSE="docker compose"
elif docker-compose version &> /dev/null;then
    COMPOSE="docker-compose"
else
    echo "docker compose is not find"
    exit 1
fi


COMPOSE_OPTS="--project-directory $WORKSPACE"
COMPOSE_FILE=$WORKSPACE/docker-compose.yml
COMPOSE_OPTS="$COMPOSE_OPTS -f $COMPOSE_FILE"

# env-file
compose_version=$($COMPOSE version | head -n 1 | awk -F ' |,|-' '{print $4}')
compose_version=${compose_version#*v}
compose_v1=$(echo $compose_version | awk -F '.' '{print $1}')
compose_v2=$(echo $compose_version | awk -F '.' '{print $2}')
if [ $compose_v1 -lt 2 ];then
    echo "docker compose version $compose_v1 is not supported"
    exit
else
    COMPOSE_OPTS="$COMPOSE_OPTS --env-file $WORKSPACE/.env"
fi

#load env
set -a
. $WORKSPACE/.env
set +a

export BMP_OMAPI_KEY=$(echo -n "$BMP_OMAPI_TOKEN" | base64)

usage(){
    help='
Usage:  ./bmp-deploy.sh [-t|--tag TAG] COMMAND SERVICE...

COMMAND:
    start       Install and start SERVICE
    stop        Stop SERVICE
    restart     restart SERVICE
    status      check SERVICE status
    images      display SERVICE image
    uninstall   uninstall SERVICE
    reinstall   reinstall SERVICE
    config      get SERVICE config

SERVICE:
    all                 All SERVICES
'
    for s in ${SERVICES_ALL[@]}; do
        help="${help}    $s\n"
    done

    printf "$help"
}

parse()
{
    POSITIONAL_ARGS=()
    DEFAULT=NO
    while [[ $# -gt 0 ]]; do
      case $1 in
        -h|--help)
          usage
          exit 0
          ;;
        -t|--tag)
          TAG="$2"
          shift
          shift
          ;;
        -*|--*)
          echo "Unknown option $1"
          exit 1
          ;;
        *)
          POSITIONAL_ARGS+=("$1") # save positional arg
          shift # past argument
          ;;
      esac
    done
}

init(){
    #parse parameters
    parse $@
    if [ ${#POSITIONAL_ARGS[@]} -lt 2 ]; then
        usage
    fi

    #check command
    COMMAND=${POSITIONAL_ARGS[0]}
    if [[ ! " ${COMMANDS_ALL[@]} " =~ " $COMMAND " ]]; then
        echo "COMMAND $COMMAND is not support!"
        usage
        exit 1
    fi

    #check services
    if [ "${POSITIONAL_ARGS[1]}" = "all" ]; then
        ALL_SERVICE=true
        SERVICES=(${SERVICES_ALL[@]})
    else
        SERVICES=(${POSITIONAL_ARGS[@]:1})
    fi

    for s in ${SERVICES[@]}; do
        if [[ ! " ${SERVICES_ALL[@]} " =~ " $s " ]]; then
            echo "service $s is not support!"
            usage
            exit 1
        fi
    done

    #set tag
    # [ -n "$TAG" ] || TAG=${BMP_VERSION:-latest}
    [ -n "$TAG" ] || TAG=${BMP_VERSION}
    [ -n "$TAG" ] || {
        echo "TAG must be set"
        usage
        exit 1
    }

    # arch tag
    #if [ "$ARCH" != "x86_64" ] && [[ ! "$TAG" =~ "$ARCH" ]];then
    #    TAG=${TAG}-${ARCH}
    #fi

    #check BMP_HOST_IP
    if [ ! -n "$BMP_HOST_IP" ] || ! ip addr show | grep "inet $BMP_HOST_IP" > /dev/null; then
        echo "BMP_HOST_IP=$BMP_HOST_IP" is not a local ip
        exit
    fi

    # load image config
    set -a
    . $WORKSPACE/config/image.cfg
    set +a

        # set umask
        umask 0022

        # create cache dir
        [ -d $WORKSPACE/$bmp_cache_dir ] || mkdir $WORKSPACE/$bmp_cache_dir
}

pre_start(){
    local service=$1

    # download bootloader images
    if [ "$service" == "bmp-tftp" ]; then
        for image in ${bmp_bootloader_images[@]};do
            if [ ! -e $WORKSPACE/${bmp_cache_dir}/${image} ];then
                echo "downloading ${bmp_oss_url}/${bmp_bootloader_prefix}/${image} ..."
                curl -sf ${bmp_oss_url}/${bmp_bootloader_prefix}/${image} -o $WORKSPACE/${bmp_cache_dir}/${image} || {
                    echo download ${bmp_oss_url}/${bmp_bootloader_prefix}/${image} fail!
                    exit 1
                }
            fi
        done
    fi

    # download LiveOS images
    if [ "$service" == "bmp-tftp" ]; then
        for image in ${bmp_liveos_images[@]};do
            if [ ! -e $WORKSPACE/${bmp_cache_dir}/${image} ];then
                echo "downloading ${bmp_oss_url}/${bmp_oss_liveos_prefix}/${image} ..."
                curl -sf ${bmp_oss_url}/${bmp_oss_liveos_prefix}/${image} -o $WORKSPACE/${bmp_cache_dir}/${image} || {
                    echo download ${bmp_oss_url}/${bmp_oss_liveos_prefix}/${image} fail!
                    exit 1
                }
            fi
        done
    fi

    # download GuestOS images
    if [ "$service" == "bmp-image" ]; then
        for image in ${bmp_guestos_images[@]};do
            if [ ! -e $WORKSPACE/${bmp_cache_dir}/${image} ];then
                echo "downloading ${bmp_oss_url}/${bmp_oss_guestos_prefix}/${image} ..."
                curl -sf ${bmp_oss_url}/${bmp_oss_guestos_prefix}/${image} -o $WORKSPACE/${bmp_cache_dir}/${image} || {
                    echo download ${bmp_oss_url}/${bmp_oss_guestos_prefix}/${image} fail!
                    exit 1
                }
            fi
        done
    fi

        # download bmp-agent
    if [ "$service" == "bmp-image" ]; then
        for image in ${bmp_agent_images[@]};do
            if [ ! -e $WORKSPACE/${bmp_cache_dir}/${image} ];then
                echo "downloading ${bmp_oss_url}/${bmp_oss_agent_prefix}/${image} ..."
                curl -sf ${bmp_oss_url}/${bmp_oss_agent_prefix}/${image} -o $WORKSPACE/${bmp_cache_dir}/${image} || {
                    echo download ${bmp_oss_url}/${bmp_oss_agent_prefix}/${image} fail!
                    exit 1
                }
            fi
        done
    fi

    return 0
}

post_start(){
    local service=$1

    # init bootloader
    if [ "$service" == "bmp-tftp" ]; then
        #legacy x86_64
        [ -e $WORKSPACE/${bmp_cache_dir}/pxelinux.0 ] && cp -f $WORKSPACE/${bmp_cache_dir}/pxelinux.0 /var/lib/bmp/bmp-tftp/

        #uefi x86_64
        mkdir -p /var/lib/bmp/bmp-tftp/uefi/x86_64
        [ -e $WORKSPACE/${bmp_cache_dir}/grubx64.efi ] && cp -f $WORKSPACE/${bmp_cache_dir}/grubx64.efi /var/lib/bmp/bmp-tftp/uefi/x86_64/grubx64.efi

        #arm64
        mkdir -p /var/lib/bmp/bmp-tftp/uefi/arm64
        [ -e $WORKSPACE/${bmp_cache_dir}/grubaa64.efi ] && cp -f $WORKSPACE/${bmp_cache_dir}/grubaa64.efi /var/lib/bmp/bmp-tftp/uefi/arm64/grubaa64.efi

        #loongarch64
        mkdir -p /var/lib/bmp/bmp-tftp/uefi/loongarch64
        [ -e $WORKSPACE/${bmp_cache_dir}/BOOTLOONGARCH64.EFI ] && cp -f $WORKSPACE/${bmp_cache_dir}/BOOTLOONGARCH64.EFI /var/lib/bmp/bmp-tftp/uefi/loongarch64/BOOTLOONGARCH64.EFI
    fi


    # init LiveOS images
    if [ "$service" == "bmp-tftp" ]; then
        # x86_64
        [ -e $WORKSPACE/${bmp_cache_dir}/${bmp_kernel_name_x86} ] && cp -f $WORKSPACE/${bmp_cache_dir}/${bmp_kernel_name_x86} /var/lib/bmp/bmp-tftp/
        [ -e $WORKSPACE/${bmp_cache_dir}/${bmp_initramfs_name_x86} ] && cp -f $WORKSPACE/${bmp_cache_dir}/${bmp_initramfs_name_x86} /var/lib/bmp/bmp-tftp/

        #arm64
        mkdir -p /var/lib/bmp/bmp-tftp/images/arm64
        [ -e $WORKSPACE/${bmp_cache_dir}/${bmp_kernel_name_arm64} ] && cp -f $WORKSPACE/${bmp_cache_dir}/${bmp_kernel_name_arm64} /var/lib/bmp/bmp-tftp/images/arm64/${bmp_kernel_name_arm64}
        [ -e $WORKSPACE/${bmp_cache_dir}/${bmp_initramfs_name_arm64} ] && cp -f $WORKSPACE/${bmp_cache_dir}/${bmp_initramfs_name_arm64} /var/lib/bmp/bmp-tftp/images/arm64/${bmp_initramfs_name_arm64}

        #loongarch64
        mkdir -p /var/lib/bmp/bmp-tftp/images/loongarch64
        [ -e $WORKSPACE/${bmp_cache_dir}/${bmp_kernel_name_loonarch64} ] && cp -f $WORKSPACE/${bmp_cache_dir}/${bmp_kernel_name_loonarch64} /var/lib/bmp/bmp-tftp/images/loongarch64/${bmp_kernel_name_loonarch64}
        [ -e $WORKSPACE/${bmp_cache_dir}/${bmp_initramfs_name_loonarch64} ] && cp -f $WORKSPACE/${bmp_cache_dir}/${bmp_initramfs_name_loonarch64} /var/lib/bmp/bmp-tftp/images/loongarch64/${bmp_initramfs_name_loonarch64}
    fi

    # init GuestOS image && device_import template && bmp-agent
    if [ "$service" == "bmp-image" ]; then
                # copy GuestOS
                for image in ${bmp_guestos_images[@]};do
                        [ -e /var/lib/bmp/bmp-image/${image} ] || cp $WORKSPACE/${bmp_cache_dir}/${image} /var/lib/bmp/bmp-image/
                done
        images=$(find /var/lib/bmp/bmp-image/ -name "*.tar.xz")
        [ -z "$images" ] && echo "error image not find" && exit 1

                # copy device_import template
        [ -e $WORKSPACE/${bmp_data_dir}/device_import.xlsx ] && cp -f $WORKSPACE/${bmp_data_dir}/device_import.xlsx /var/lib/bmp/bmp-image/

                # copy bmp-agent
                for image in ${bmp_agent_images[@]}; do
                        [ -e /var/lib/bmp/bmp-image/${image} ] || cp $WORKSPACE/${bmp_cache_dir}/${image} /var/lib/bmp/bmp-image/
                done
    fi

    #init db
    if [ "$service" == "bmp-db" ]; then
        if [ -e $WORKSPACE/sql/bmp.sql ];then
            #wait for db to be ready
            container_name="bmp-bmp-db-1"
            [ "$compose_v1" -lt 2 ] &&  container_name="bmp_bmp-db_1"
            retry=60
            while [ $retry -gt 0 ];do
                docker exec -i $container_name sh -c 'exec mysql -uroot -p"$MYSQL_ROOT_PASSWORD" $MYSQL_DATABASE -e  "show tables"' &> /dev/null && break
                retry=$(($retry-1))
                sleep 2;
            done
            if [ $retry -le 0 ]; then
                echo "error! bmp-db is not ready"
                return 1
            fi

            # 检测数据库是否存在，防止数据被覆盖
            n=$(docker exec -i $container_name sh -c 'exec mysql -uroot -p"$MYSQL_ROOT_PASSWORD" $MYSQL_DATABASE -e  "show tables"' 2>/dev/null | wc -l)
            if [ $n -lt 2 ]; then
                echo "exec sql ..."
                docker exec -i $container_name sh -c 'exec mysql -uroot -p"$MYSQL_ROOT_PASSWORD" $MYSQL_DATABASE' < $WORKSPACE/sql/bmp.sql || true
            fi
        fi
    fi

    # check status
    n=$($COMPOSE $COMPOSE_OPTS ps $service | grep -P 'running|Up' | wc -l)
    if [ $n -eq 0 ]; then
        $COMPOSE $COMPOSE_OPTS up -d --quiet-pull $service
        echo "$service start again"
    fi

    return
}

do_start(){
    for service in ${SERVICES[@]};do
        pre_start $service
    done
    echo ${bmp_db_image}
    $COMPOSE $COMPOSE_OPTS up -d --quiet-pull ${SERVICES[@]}
    for service in ${SERVICES[@]};do
        post_start $service
    done
}

do_stop(){
    $COMPOSE $COMPOSE_OPTS stop ${SERVICES[@]}
}

do_restart(){
    do_stop
    do_start
}

do_clean(){
    for service in ${SERVICES[@]}; do
        [ -e /var/lib/bmp/$service ] && rm -rf /var/lib/bmp/$service
        [ -e /var/log/bmp/$service ] && rm -rf /var/log/bmp/$service
    done
    echo do_clean done
}

do_uninstall(){
    do_stop
    for s in ${SERVICES[@]};do
        container_id=$($COMPOSE $COMPOSE_OPTS ps -a -q $s)
        image=$($COMPOSE $COMPOSE_OPTS images $s | awk 'NR==2{print $2":"$3}')
        [ -n "$container_id" ] && docker rm $container_id
        [ -n "$image" ] && docker rmi $image
    done
    do_clean
}

do_reinstall(){
    do_uninstall
    do_start
}

do_status(){
    is_ok=true
    for s in ${SERVICES[@]}; do
        n=$($COMPOSE $COMPOSE_OPTS ps $s | grep -P 'running|Up' | wc -l)
        if [ $n -eq 1 ]; then
            status="running"
        else
            status="not running"
        fi
        printf "%-20s%s\n" "$s" "$status"
        [ "$status" != "running" ] && is_ok=false
    done
    [ "$is_ok" = "false" ] && return 1
    return 0
}

do_images(){
    $COMPOSE $COMPOSE_OPTS images ${SERVICES[@]}
}

do_config(){
    $COMPOSE $COMPOSE_OPTS config ${SERVICES[@]}
}

main(){
    init $@
    do_$COMMAND
}

# 支持多架构
if [ "$ARCH" == "x86_64" ] || [ "$ARCH" == "aarch64" ] || [ "$ARCH" == "loongarch64" ];then
    main $@
else
    echo "architecture $ARCH is not supported!"
    exit 1
fi
