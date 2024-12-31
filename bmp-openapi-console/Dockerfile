FROM centos:7.9.2009

MAINTAINER minping@jd.com
USER root
RUN groupadd -g 2001 ironic-api
RUN useradd -d /home/ironic-api -u 2001 -g 2001 ironic-api

WORKDIR /home/ironic-api

ADD ./ironic-api /home/ironic-api/
RUN mkdir -p /home/ironic-api/conf
ADD ./conf/* /home/ironic-api/conf/

RUN chmod 755 /home/ironic-api/ironic-api
USER ironic-api
EXPOSE 9999
CMD ./ironic-api