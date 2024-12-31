#!/bin/bash
set -ue
WORKSPACE=$(cd $(dirname $0); pwd)
set -a
. $WORKSPACE/.env
. $WORKSPACE/config/image.cfg
. $WORKSPACE/config/service.cfg
set +a

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

usage(){
    help='
Usage:  ./bmp-build-image.sh build SERVICE...

COMMAND:
    build       Build image

SERVICE:
    all                 All SERVICES
'
    for s in ${bmp_service_app[@]}; do
        help="${help}    $s\n"
    done
    for s in ${bmp_service_other[@]}; do
        help="${help}    $s\n"
    done


    printf "$help"
}


parse $@
if [ ${#POSITIONAL_ARGS[@]} -lt 2 ]; then
        usage
fi

#check command
COMMAND=${POSITIONAL_ARGS[0]}
if [[ "build" != "$COMMAND" ]]; then
   echo "COMMAND $COMMAND is not support!"
   usage
   exit 1
fi

ALL_SERVICES=("${bmp_service_app[@]}" "${bmp_service_other[@]}")
if [ "${POSITIONAL_ARGS[1]}" = "all" ]; then
    SERVICES=(${ALL_SERVICES[@]})
else
    SERVICES=(${POSITIONAL_ARGS[@]:1})
    for s in ${SERVICES[@]}; do
        if [[ ! " ${ALL_SERVICES[@]} " =~ " $s " ]]; then
            echo "service $s is not support!"
            usage
            exit 1
        fi
    done
fi

# build image
bmp_code_dir=$(dirname ${WORKSPACE})
for service_name in ${SERVICES[@]}; do
    dockerfile_name=${bmp_code_dir}/${service_name}/${service_name}.Dockerfile
    docker_tag=${BMP_REGISTRY}/${service_name}:${BMP_VERSION}
    build_code_dir=${bmp_code_dir}/${service_name}

    if [ ${service_name} == "bmp-tftp" ] || [ ${service_name} == "bmp-rsyslog" ];then
        dockerfile_name=${WORKSPACE}/dockerfile/${service_name}.Dockerfile
        docker_tag=${BMP_REGISTRY}/${service_name}:${BMP_VERSION}
        build_code_dir=${bmp_code_dir}

    elif [ ${service_name} == "bmp-oob-agent" ];then
        dockerfile_name=${bmp_code_dir}/oob-log-agent/oob-log-agent.Dockerfile
        docker_tag=${BMP_REGISTRY}/oob-log-agent:${BMP_VERSION}
        build_code_dir=${bmp_code_dir}/oob-log-agent

    elif [ ${service_name} == "bmp-oob-alert" ];then
        dockerfile_name=${bmp_code_dir}/oob-log-alert/oob-log-alert.Dockerfile
        docker_tag=${BMP_REGISTRY}/oob-log-alert:${BMP_VERSION}
        build_code_dir=${bmp_code_dir}/

    elif [ ${service_name} == "bmp-console-api" ] || [ ${service_name} == "bmp-operation-api" ];then
        dockerfile_name=${bmp_code_dir}/${service_name}/${service_name}.Dockerfile
        docker_tag=${BMP_REGISTRY}/${service_name}:${BMP_VERSION}
        build_code_dir=${bmp_code_dir}/

    else
        dockerfile_name=${bmp_code_dir}/${service_name}/${service_name}.Dockerfile
        docker_tag=${BMP_REGISTRY}/${service_name}:${BMP_VERSION}
        build_code_dir=${bmp_code_dir}/${service_name}
    fi

    if [ ! -f "${dockerfile_name}" ];then
        echo ${service_name} dockerfile : "${dockerfile_name}" do not exists!!
        exit 1
    fi
    echo "${service_name} begin build image ... "
    echo "image name: ${docker_tag}"
    echo "build code path: ${build_code_dir}"
    docker build -f ${dockerfile_name} -t ${docker_tag} ${build_code_dir}
    echo "build success !!"
    echo "---------------------------------------------------------"
done
