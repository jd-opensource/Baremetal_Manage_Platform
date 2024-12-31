FROM centos:7.9.2009

MAINTAINER minping@jd.com
USER root
RUN groupadd -g 2001 ironic-console
RUN useradd -d /home/ironic-console -u 2001 -g 2001 ironic-console

WORKDIR /home/ironic-console

ADD ./ironic-console /home/ironic-console/
RUN mkdir -p /home/ironic-console/conf
ADD ./conf/* /home/ironic-console/conf/

RUN chmod 755 /home/ironic-console/ironic-console
USER ironic-console
EXPOSE 9998
CMD ./ironic-console