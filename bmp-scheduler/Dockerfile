 FROM centos:7.9.2009

MAINTAINER minping@jd.com
USER root
RUN groupadd -g 2001 ironic-scheduler
RUN useradd -d /home/ironic-scheduler -u 2001 -g 2001 ironic-scheduler

WORKDIR /home/ironic-scheduler

ADD ./ironic-scheduler /home/ironic-scheduler/
RUN mkdir -p /home/ironic-scheduler/conf
ADD ./conf/* /home/ironic-scheduler/conf/

RUN chmod 755 /home/ironic-scheduler/ironic-scheduler
USER ironic-scheduler
CMD ./ironic-scheduler