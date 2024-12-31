#!/bin/bash

tmp_file=/tmp/bmp-agentd.tgz
target_dir=/opt/bmp

if [[ $(id -u) -ne 0 ]]; then
    echo "ERROR: Insufficient privilege - SUPER privileges required";
    exit 1
fi

sed -n -e '1,/^exit 0$/!p' $0 > ${tmp_file} 2>/dev/null

tar xzf ${tmp_file} -C /tmp

if [ "$1" = "-v=true" ]; then # 加上-v=true参数时，只打印版本号就退出
    res=$(/tmp/bmp-agent -v=true)
    echo $res
    exit 0
fi

# remove old version
if [ -f /etc/init.d/bmpd ]; then
    if which chkconfig >/dev/null 2>&1; then
        service bmpd stop
        chkconfig --del bmpd
    elif which update-rc.d >/dev/null 2>&1; then
        service bmpd stop
        update-rc.d -f bmpd remove
    else
        echo "ERROR: Stop BMP Agent Failed";
    fi
    /bin/rm -f /etc/init.d/bmpd
fi

if [ -d ${target_dir} ]; then
    /bin/rm -rf ${target_dir}
fi

# unarchive tar and install 
# tar xzf ${tmp_file} -C /tmp
mkdir -p ${target_dir}/bin
cp -f /tmp/bmp-agent ${target_dir}/bin/
cp -f /tmp/bmpd ${target_dir}/bin/
ln -fs ${target_dir}/bin/bmpd /etc/init.d/bmpd

#auto start
if which chkconfig >/dev/null 2>&1; then
    chkconfig --add bmpd
    chkconfig bmpd on
elif which update-rc.d >/dev/null 2>&1; then
    update-rc.d bmpd defaults 90
else
    echo "ERROR: Auto Start JD CPS Agent Failed";
fi

# start
if which service >/dev/null 2>&1; then
    service bmpd start
	sleep 1
	service bmpd status
else
    echo "ERROR: Start JD CPS Agent Failed";
fi

if [ $? -ne 0 ]; then
    echo "ERROR: Start JD CPS Agent Failed";
fi

# clean
/bin/rm -rf /tmp/bmp-agent /tmp/bmpd ${tmp_file}

exit 0
