#!/bin/bash
###############################################################################
#编译脚本的原理是将编译结果放到output目录中，这个样例模版提供一个产生
#一个最基本golang运行程序包的编译脚本，对于特殊的需求请酌情考虑
#
#1、该脚本支持参数化，参数将传入build_package函数（内容为最终执行的编译命令）
#   ，用$1,$2....表示，第1,2...个参数
#2、部署需要启动程序，所以需要提供control文件放在当前目录中，用于启动和
#   监控程序状态

###############用户修改部分################
readonly PROJECT_NAME="ironic-jdstack"    #项目名称,相对于$GOPATH/src文件夹所在的目录,可选项
readonly PACKAGE_DIR_NAME=main.go   #main文件相对于src文件夹所在的目录,可选项
readonly PACKAGE_BIN_NAME=ironic-jdstack   #定义产出的运行程序名,必填项
readonly CONF_DIR_NAME="conf"       #定义配置文件目录,此路径为相对路径,可选项
readonly OTHERS_DIR_NAME=""       #定义其它文件目录,多个目录以空格区分，此路径为相对路径,可选项
#最终的抽包路径为$OUTPUT
###########################################

if [[ "${PACKAGE_BIN_NAME}" == "" ]];then
    echo "Please set "PACKAGE_BIN_NAME" value"
    exit 1
fi

function set_work_dir
{
    readonly OUTPUT=$(pwd)/output
    readonly WORKSPACE_DIR=$(pwd)
}

#清理编译构建目录操作
function clean_before_build
{
    cd ${WORKSPACE_DIR}
    rm -rf bin pkg
    rm -rf ${OUTPUT}
}

#初始化编译环境
function prepare_package()
{
    if [ ! -z "$PROJECT_NAME" ];then
        mkdir -p src/$PROJECT_NAME
	    copy_files=`ls |grep -vw 'src'`
	    cp -rf $copy_files src/$PROJECT_NAME/
	fi
}

#实际的编译命令
#这个函数中可使用$1,$2...获取第1,2...个参数
function build_package()
{
    cd ${WORKSPACE_DIR}
    # export GOPATH=$(pwd)
	go clean
    echo $(pwd)
    echo $(ls)
	go build -v -gcflags="-N -l" -o ${PACKAGE_BIN_NAME} ${PACKAGE_DIR_NAME} || return 1
}

#建立最终发布的目录
function build_dir
{
    mkdir -p ${OUTPUT}/bin || return 1
}

function dir_not_empty()
{
    if [[ ! -d $1 ]];then
        return 1
    fi
    if [[ $(ls $1|wc -l) -eq 0 ]];then
        return 1
    fi
    return 0
}

#拷贝编译结果到发布的目录
function copy_result
{
    cd ${WORKSPACE_DIR}
    cp -r ${PACKAGE_BIN_NAME} ${OUTPUT}/bin/ || return 1
    cp -r ./control ${OUTPUT}/bin || return 1
    dir_not_empty ${WORKSPACE_DIR}/${CONF_DIR_NAME} && mkdir -p ${OUTPUT}/${CONF_DIR_NAME}
    cp -rf ./${CONF_DIR_NAME}/* ${OUTPUT}/${CONF_DIR_NAME}/
    test -z "$OTHERS_DIR_NAME" || cp -rf $OTHERS_DIR_NAME ${OUTPUT}/
    return 0
}

#执行
function main()
{
    cd $(dirname $0)
    set_work_dir

    echo "At: "$(date "+%Y-%m-%d %H:%M:%S") 'Cleaning...'
    clean_before_build || exit 1
    echo "At: "$(date "+%Y-%m-%d %H:%M:%S") 'Clean completed'
    echo
	
    echo "At: "$(date "+%Y-%m-%d %H:%M:%S") 'preparng...'
    prepare_package || exit 1
    echo "At: "$(date "+%Y-%m-%d %H:%M:%S") 'prepare completed'
    echo
	
    echo "At: "$(date "+%Y-%m-%d %H:%M:%S") 'Building...'
    build_package $@ || exit 1
    echo "At: "$(date "+%Y-%m-%d %H:%M:%S") 'Build completed'
    echo

    echo "At: "$(date "+%Y-%m-%d %H:%M:%S") 'Making dir...'
    build_dir || exit 1
    echo "At: "$(date "+%Y-%m-%d %H:%M:%S") 'Make completed'
    echo

    echo "At: "$(date "+%Y-%m-%d %H:%M:%S") 'Copy result to publish dir...'
    copy_result || exit 1
    echo "At: "$(date "+%Y-%m-%d %H:%M:%S") 'Copy completed'
    echo

    exit 0
}

main $@
