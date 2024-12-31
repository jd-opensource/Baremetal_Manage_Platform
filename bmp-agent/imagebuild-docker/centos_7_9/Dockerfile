FROM centos:7.9.2009
ENV ARCH=x86_64
ENV BASE_DIR=imagebuild-docker/$ARCH
RUN yum install -y dhcp \
&& yum install -y rsyslog \
&& yum install -y python3 \
&& yum install -y python3-devel \
&& yum install -y libyaml-devel \
&& yum install -y biosdevname \
&& yum install -y lshw \
&& yum install -y hdparm \
&& yum install -y ipmitool \
&& yum install -y qemu-kvm libvirt libvirt-python libguestfs-tools virt-install \
&& yum install -y dmidecode \
&& yum install -y gdisk \
&& yum install -y smartmontools \
&& yum install -y vim \
&& yum install -y sysfsutils \
&& yum install -y epel-release dkms \
&& yum install -y net-tools \
&& yum install -y curl \
&& yum install -y ntpdate \
&& yum install -y sg3_utils \
&& yum install -y kexec-tools \
&& yum install -y python3-pip \
&& yum install -y ntfs-3g \
&& yum install -y sysstat \
&& yum install -y tcpdump \
&& yum install -y telnet \
&& yum install -y parted \
&& yum install -y psmisc \
&& yum install -y efibootmgr \
&& yum install -y dosfstools \
&& yum install -y xfsprogs \
&& yum install -y make \
&& yum install -y gcc
# base config
RUN mkdir -p /root/.pip
COPY $BASE_DIR/chroot/pip.conf /root/.pip/
# fstab
COPY $BASE_DIR/chroot/fstab $CHROOT_DIR/etc/fstab

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
# if has not this file, interface will up failed
&& touch network /etc/sysconfig/network

# nvme-cli
COPY $BASE_DIR/packages/nvme-cli-1.8.1-3.el7.x86_64.rpm /tmp/
RUN rpm -i /tmp/nvme-cli-1.8.1-3.el7.x86_64.rpm

# storcli
COPY $BASE_DIR/packages/storcli-007.1613.0000.0000-1.noarch.rpm /tmp/
RUN rpm -i /tmp/storcli-007.1613.0000.0000-1.noarch.rpm \
&& ln -s /opt/MegaRAID/storcli/storcli64 /usr/sbin/storcli64
# megacli
COPY $BASE_DIR/packages/Lib_Utils2-1.00-10.noarch.rpm /tmp/
RUN rpm -i /tmp/Lib_Utils2-1.00-10.noarch.rpm
COPY $BASE_DIR/packages/MegaCli-8.07.14-1.noarch.rpm /tmp/
RUN rpm -i /tmp/MegaCli-8.07.14-1.noarch.rpm \
&& ln -s /opt/MegaRAID/MegaCli/MegaCli64 /usr/sbin/MegaCli64
# perccli
COPY $BASE_DIR/packages/perccli-007.1623.0000.0000-1.noarch.rpm /tmp/
RUN rpm -i /tmp/perccli-007.1623.0000.0000-1.noarch.rpm \
&& ln -s /opt/MegaRAID/perccli/perccli64 /usr/sbin/perccli64
# sas2ircu
COPY $BASE_DIR/packages/sas2ircu /usr/sbin
RUN chmod +x /usr/sbin/sas2ircu
# sas3ircu
COPY $BASE_DIR/packages/sas3ircu /usr/sbin
RUN chmod +x /usr/sbin/sas3ircu

# arcconf
COPY $BASE_DIR/packages/arcconf /usr/sbin
RUN chmod +x /usr/sbin/arcconf

COPY $BASE_DIR/chroot/sshd_config $CHROOT_DIR/etc/ssh/
# bmp-agent service
COPY $BASE_DIR/chroot/bmpa.service /lib/systemd/system/
COPY requirements.txt /tmp/bmp-agent/
RUN pip3 install  -r /tmp/bmp-agent/requirements.txt
COPY bmpa /tmp/bmp-agent/bmpa
COPY MANIFEST.in /tmp/bmp-agent/
COPY setup.py /tmp/bmp-agent/
RUN  cd /tmp/bmp-agent && python3 setup.py install \
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