FROM centos:7.9.2009

MAINTAINER minping@jd.com
USER root
RUN groupadd -g 2001 ironic-driver
RUN useradd -d /home/ironic-driver -u 2001 -g 2001 ironic-driver

WORKDIR /home/ironic-driver

ADD ./ironic-driver /home/ironic-driver/
RUN mkdir -p /home/ironic-driver/conf
ADD ./conf/* /home/ironic-driver/conf/

RUN chmod 755 /home/ironic-driver/ironic-driver
USER ironic-driver
CMD ./ironic-driver