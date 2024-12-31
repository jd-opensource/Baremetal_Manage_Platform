#! /bin/bash

set -e

script_dir=$(cd $(dirname "$0") && pwd)

#load env
#set -a
#. $WORKSPACE/.env
#set +a

# 帮助文档
show_help() {
        echo "Usage: $0 [-h|--help] [-e|--env <public|private>]"
        echo
        echo "Options:"
        echo "  -h, --help       Show this help message and exit"
        echo "  -e, --env        Set the environment to 'public' or 'private'"
}

# 初始化变量
deploy_env=""

# 解析命令行参数
while [[ "$#" -gt 0 ]]; do
    case $1 in
        -h|--help)
            show_help
            exit 0
            ;;
        -e|--env)
            if [[ -n "$2" ]]; then
                if [[ "$2" == "public" || "$2" == "private" ]]; then
                    deploy_env="$2"
                    shift
                else
                    echo "Error: The -e|--env option can only be 'public' or 'private'."
                    exit 1
                fi
            else
                echo "Error: The -e|--env option requires an argument."
                exit 1
            fi
            ;;
        *)
            echo "Error: Invalid option $1"
            show_help
            exit 1
            ;;
    esac
    shift
done

if [[ -n "$deploy_env" ]]; then
    echo "Environment set to: $deploy_env"
else
    echo "No environment set."
    echo "  -e, --env        Set the environment to 'public' or 'private'"
    exit 1
fi

function check_docker(){
	# 检查 Docker 是否安装
	if ! docker version &> /dev/null; then
    		echo "Error: Docker is not installed."
		# offline install
		if [ -d $script_dir/docker_install_files/docker ];then
			echo "Now starting offline installation of Docker."
			cp $script_dir/docker_install_files/docker/bin/* /usr/bin/
    			cp $script_dir/docker_install_files/docker/docker.service /etc/systemd/system/
    			chmod +x /usr/bin/docker*
    			chmod 644 /etc/systemd/system/docker.service
			systemctl daemon-reload &> /dev/null
			systemctl enable docker &> /dev/null
			systemctl start docker &> /dev/null 
			if [ $? -nq 0]; then
				echo "Error: Docker installation failed, please install manually or contact technical support."
				exit 1
			fi
			echo "Docker installation success."
 		fi
	fi

	if ! (docker-compose version &>/dev/null || docker compose version &>/dev/null); then
		echo "Docker Compose is not installed."
                # offline install
                if [ -d $script_dir/docker_install_files/docker-compose ];then
                        echo "Now starting offline installation of Docker-compose."			
                        cp $script_dir/docker_install_files/docker-compose/* /usr/bin/
			chmod u+x $script_dir/docker_install_files/docker-compose/* 
			docker-compose version &>/dev/null
                        if [ $? -nq 0 ]; then
                                echo "Error: Docker-compose installation failed, please install manually or contact technical support."
                                exit 1
                        fi
			echo "Docker-compose installation success."
                fi
	fi
	# 获取 Docker 版本
	# docker_version=$(docker --version 2>&1)

	# 使用正则表达式从版本字符串中提取主要版本号
	#major_version=$(echo "$docker_version" | grep -oP 'Docker version \K(\d+)')

	# 比较主要版本号
	#if (( $major_version < 2 )); then
    	#	echo "Error: Docker version is too low. Please update to latest version."
    	#	exit 1
	# fi

	echo "Docker and docker compose are installed."
}

function check_port(){
	port_list=$(grep -E -o "\- ([0-9]+):([0-9]+)" docker-compose.yml | awk -F "[ :]" '{print $2}')
	check_port_res=0
	for every_port in ${port_list[@]};do
		if (lsof -i :$every_port | grep -E "LISTEN|UDP" &> /dev/null);then
			echo "port $every_port is already in use. "
			check_port_res=1
		fi
	done
	if [ $check_port_res -ne 0 ];then
		echo "ERROR: The BMP system needs to use these ports, but these ports is already in use. Please confirm."
		exit 1
	fi
	echo "All ports are available."
}


function update_env_ip(){
	# 获取本机的所有 IP 地址
	ips=$(ip addr show | awk '/^[0-9]+: / {}; /inet.*global/ {print $2}' | cut -d '/' -f 1)

	# 如果有多个 IP 地址，提示用户选择一个作为主 IP
	if [[ $(echo "$ips" | wc -l) -gt 1 ]]; then
    		echo "This machine has multiple IP addresses:"
    		echo "$ips"
		while true; do
    			read -p "Please enter the IP address to use as the main IP for external communication: " main_ip
			if echo "$ips" | grep -q "$main_ip"; then
				echo "$main_ip is a valid IP address"
				break;
			else
				echo "$main_ip is not a valid IP address, please check the spelling and re-enter"
			fi
		done
	else
    	# 如果只有一个 IP 地址，直接使用它作为主 IP
    		main_ip=$ips
	fi
	sed -i "s/^BMP_HOST_IP=.*/BMP_HOST_IP=$main_ip/g" $script_dir/.env
	echo "The main IP address for external communication is: $main_ip"
} 


function private_load_docker_images(){
	docker_image=$script_dir/private_deploy_docker_images/bmp-image.tar.gz
	if [ ! -f $docker_image ];then
		echo "$docker_image does not exist. Please confirm whether the file download is correct, or contact the BMP colleague."
		exit 1
	fi
	docker load -i $docker_image 
	if [ $? -ne 0 ]; then
  		echo "Failed to load Docker image: $docker_image" >&2
  		exit 1
	fi
}

function public_load_docker_images(){
	. $script_dir/.env
	. $script_dir/config/image.cfg
	images_list=(
	bmp_rsyslog_image
	bmp_driver_image
	bmp_dhcp_agent_image
	bmp_openapi_image
	bmp_openapi_console_image
	bmp_scheduler_image
	bmp_console_api_image
	bmp_operation_api_image
	bmp_console_web_image
	bmp_operation_web_image
	bmp_monitor_proxy_image
	bmp_pronoea_image
	bmp_tftp_image
	oob_log_alert_image
	oob_log_agent_image
	bmp_redis_image
	bmp_mq_image
	bmp_image_image
	bmp_pushgateway_image
	bmp_alertmanager_image
	bmp_prometheus_image
	bmp_db_image
	)	
	for image_var in "${images_list[@]}"; do
		image_name=${!image_var}
    		echo "begin docker pull ${image_name}"
		if docker images --format "{{.Repository}}:{{.Tag}}" | grep -q "^${image_name}$"; then
  			echo "Image ${image_name} already exists locally. Skipping docker pull."
		else
			docker pull ${image_name}
			if [ $? -ne 0 ];then
				echo "docker pull failed"
				exit 1
			fi
			echo "docker pull ${!image_var} success!"
		fi
	done
}



function main(){
	echo -ne "Step 1: Check if docker is installed.\n"
	echo ""
	check_docker
	echo "-------------------------------------------------"
	echo -ne "Step 2: Check if port is available.\n"
        echo ""
        check_port
        echo "-------------------------------------------------"
	echo -ne "Step 3: update ip in .env \n"
        echo ""
	# update_env_ip
	echo "-------------------------------------------------"
	echo -ne "Step 4: docker load images \n"
        echo ""
        if [[ "$deploy_env" == "private" ]]; then
          private_load_docker_images 
        elif [[ "$deploy_env" == "public" ]]; then
          public_load_docker_images
	else
    	  echo "No valid environment set."
          exit 1
	fi
	echo "
		now bmp-deploy-precheck.sh success 
                please run: bash bmp-deploy.sh start all

	"
}

main
