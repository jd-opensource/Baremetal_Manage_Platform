FROM ubuntu:22.04
ENV ARCH=arm64
ENV BASE_DIR=imagebuild-docker/$ARCH
COPY $BASE_DIR/chroot/sources.22.04.list /etc/apt/sources.list
RUN apt-get update \
&& apt-get install -y rsyslog \
&& apt-get install -y python3 \
&& apt-get install -y python3-dev  \
&& apt-get install -y libyaml-dev \
&& apt-get install -y --assume-yes apt-utils \
&& apt-get install -y  initramfs-tools \
&& apt-get install -y lshw \
&& apt-get install -y hdparm \
&& apt-get install -y ipmitool \
&& apt-get install -y qemu-utils \
&& apt-get install -y dmidecode \
&& apt-get install -y gdisk \
&& DEBIAN_FRONTEND=noninteractive apt-get install -y -o Dpkg::Options::="--force-confdef" -o Dpkg::Options::="--force-confold" smartmontools \
&& apt-get install -y vim \
&& apt-get install -y sysfsutils \
&& apt-get install -y dkms \
&& apt-get install -y net-tools \
&& apt-get install -y ifupdown \
&& apt-get install -y curl \
&& apt-get install -y ntpdate \
&& apt-get install -y sg3-utils \
&& DEBIAN_FRONTEND=noninteractive apt-get install -y kexec-tools \
&& apt-get install -y python3-pip \
&& apt-get install -y ntfs-3g \
&& apt-get install -y sysstat \
&& apt-get install -y tcpdump \
&& apt-get install -y telnet \
&& apt-get install -y parted \
&& apt-get install -y psmisc \
&& apt-get install -y efibootmgr \
&& apt-get install -y dosfstools \
&& apt-get install -y xfsprogs  \
&& apt-get install -y nvme-cli
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
# network
RUN rm -rvf /etc/netplan/*
COPY $BASE_DIR/chroot/interfaces /etc/network/

# storcli64
COPY $BASE_DIR/packages/storcli64 $CHROOT_DIR/usr/sbin
RUN chmod +x /usr/sbin/storcli64

# sas3ircu
COPY $BASE_DIR/packages/sas3ircu $CHROOT_DIR/usr/sbin
RUN chmod +x /usr/sbin/sas3ircu

COPY $BASE_DIR/chroot/sshd_config $CHROOT_DIR/etc/ssh/
# openssh-server
RUN apt-get install -y openssh-server \
&& mkdir -p /root/.ssh \
&& chmod 700 /root/.ssh \
&& touch /root/.ssh/authorized_keys \
&& chmod 600 /root/.ssh/authorized_keys
COPY $BASE_DIR/chroot/id_rsa.pub  /tmp/
RUN cat /tmp/id_rsa.pub /root/.ssh/authorized_keys \
# dns
&& echo UseDNS no >> /etc/ssh/sshd_config \
# kernel
&& apt-get install -y linux-image-generic \
&& apt-get remove -y linux-image-generic 

# bmp-agent service
COPY $BASE_DIR/chroot/bmpa.service /lib/systemd/system/
COPY requirements.txt /tmp/bmp-agent/
RUN pip3 install  -r /tmp/bmp-agent/requirements.txt
COPY bmpa /tmp/bmp-agent/bmpa
COPY MANIFEST.in /tmp/bmp-agent/
COPY requirements.txt /tmp/bmp-agent/
COPY setup.py /tmp/bmp-agent/
RUN pip3 install  -r /tmp/bmp-agent/requirements.txt \
&& cd /tmp/bmp-agent && python3 setup.py install \
&& systemctl enable bmpa.service \
# clean base system
&& systemctl disable smartd \
&& systemctl disable systemd-tmpfiles-setup \
&& apt-get remove -y python3-dev \
&& apt-get remove -y libyaml-dev \
&& apt-get clean -y \
&& rm -rf /var/cache/* \
&& rm -rf /var/lib/apt/lists/* \
&& rm -rf /var/log/* \
&& rm -rf /usr/share/doc/* \
&& rm -rf /usr/share/man/* \
&& rm -rf /usr/share/locale/* \
&& rm -rf /tmp/* \
&& rm -rf /usr/lib/firmware/iwlwifi-* \
&& rm -rf /usr/lib/modules/*/kernel/drivers/media \
&& rm -rf /usr/lib/modules/*/kernel/drivers/usb \
&& rm -rf /usr/lib/modules/*/kernel/drivers/video \
&& rm -rf /usr/lib/modules/*/kernel/drivers/net/wireless \
&& rm -rf /usr/lib/modules/*/kernel/drivers/net/usb \
&& rm -rf /usr/lib/firmware/netronome/nic-sriov \
&& rm -rf /usr/lib/modules/5.15.0-112-generic/kernel/sound \
&& rm -rf /usr/lib/modules/5.15.0-112-generic/kernel/zfs \
&& rm -rf /usr/lib/modules/5.15.0-112-generic/kernel/net/wireless \
&& rm -rf /usr/lib/modules/5.15.0-112-generic/kernel/net/mac80211 \
&& rm -rf /usr/lib/modules/5.15.0-112-generic/kernel/net/bluetooth \
&& rm -rf /usr/lib/modules/5.15.0-112-generic/kernel/net/sunrpc \
&& rm -rf /usr/lib/modules/5.15.0-112-generic/kernel/net/openvswitch \
&& rm -rf /usr/lib/modules/5.15.0-112-generic/kernel/net/ceph \
&& rm -rf /usr/lib/modules/5.15.0-112-generic/kernel/drivers/infiniband 