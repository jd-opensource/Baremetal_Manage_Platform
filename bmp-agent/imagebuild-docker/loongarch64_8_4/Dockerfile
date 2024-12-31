FROM scratch
ENV bmp_oss_url https://bmp.s3-internal.cn-north-1.jdcloud-oss.com
ENV bmp_oss_liveos_prefix LiveOS
ENV image_name v1.7.0-loongnix-8.4-2023053018.tar.xz 

ADD ${bmp_oss_url}/${bmp_oss_liveos_prefix}/${image_name} /
RUN rm -rf /v1.7.0-loongnix-8.4-2023053018.tar.xz
ENV ARCH=loongarch64
ENV BASE_DIR=imagebuild-docker/$ARCH
RUN yum install -y rsyslog \
&& yum install -y python3 \
&& yum install -y python3-devel \
&& yum install -y libyaml-devel \
&& yum install -y lshw \
&& yum install -y hdparm \
&& yum install -y ipmitool \
&& yum install -y dmidecode \
&& yum install -y gdisk \
&& yum install -y smartmontools \
&& yum install -y vim \
&& yum install -y sysfsutils \
&& yum install -y net-tools \
&& yum install -y curl \
&& yum install -y sg3_utils \
&& yum install -y kexec-tools \
&& yum install -y python3-pip \
&& yum install -y sysstat \
&& yum install -y tcpdump \
&& yum install -y telnet \
&& yum install -y parted \
&& yum install -y psmisc \
&& yum install -y efibootmgr \
&& yum install -y dosfstools \
&& yum install -y xfsprogs \
&& yum install -y make \
&& yum install -y gcc \
&& yum install -y tar
# base config
# fstab
COPY $BASE_DIR/chroot/fstab /etc/fstab

RUN echo "root:beijing" | chpasswd \
&& echo "SELINUX=disabled" > /etc/selinux/config \
&& mkdir -p /etc/systemd/getty/root@.service.d \
&& echo "[Service]"  >> /etc/systemd/getty/root@.service.d/noclear.conf \
&& echo "TTYVTDisallocate=no"  >> /etc/systemd/getty/root@.service.d/noclear.conf 
# network
RUN for i in $(seq 0 8);  \
do \
echo DEVICE=eth${i} >> /etc/sysconfig/network-scripts/ifcfg-eth${i}; \
echo "BOOTPROTO=dhcp" >> /etc/sysconfig/network-scripts/ifcfg-eth${i}; \
echo "ONBOOT=yes" >> /etc/sysconfig/network-scripts/ifcfg-eth${i}; \
done \
# If has not this file, interface will up failed
&& touch network /etc/sysconfig/network

RUN yum install -y nvme-cli

# storcli
COPY $BASE_DIR/packages/raid-static-conf-loongarch-storcli-v0.6 /tmp/raid-static-conf-loongarch-storcli-v0.6
RUN cd /tmp/raid-static-conf-loongarch-storcli-v0.6 && ./installer \
&& ln -s /usr/bin/storcli64_ls /usr/sbin/storcli64


COPY $BASE_DIR/chroot/sshd_config /etc/ssh/
# bmp-agent service
COPY $BASE_DIR/chroot/bmpa.service /lib/systemd/system/
COPY bmpa /tmp/bmp-agent/bmpa
COPY MANIFEST.in /tmp/bmp-agent/
COPY requirements.txt /tmp/bmp-agent/
COPY setup.py /tmp/bmp-agent/
RUN  pip3 install -r /tmp/bmp-agent/requirements.txt \
&&  cd /tmp/bmp-agent && python3 setup.py install \
&& systemctl enable bmpa.service \

# openssh-server
&& yum install -y openssh-server \
&& mkdir -p /root/.ssh \
&& chmod 700 /root/.ssh \
&& touch /root/.ssh/authorized_keys \
&& chmod 600 /root/.ssh/authorized_keys
COPY $BASE_DIR/chroot/id_rsa.pub  /tmp/
RUN cat /tmp/id_rsa.pub /root/.ssh/authorized_keys \
# dns
&& echo UseDNS no >> /etc/ssh/sshd_config \

# clean base system
&& systemctl disable smartd \
&& systemctl disable systemd-tmpfiles-setup \
&& yum remove -y python3-devel \
&& yum remove -y libyaml-devel \
&& yum clean all \
&& rm -rf /var/cache/* \
&& rm -rf /var/lib/apt/lists/* \
&& rm -rf /var/log/* \
&& rm -rf /usr/share/doc/* \
&& rm -rf /usr/share/man/* \
&& rm -rf /usr/share/locale/* \
&& rm -rf /tmp/*
