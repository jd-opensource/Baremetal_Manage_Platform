set timeout=1
set default=1
menuentry 'bmp agent {{arch}} uefi' {
    linuxefi {{kernel_path}} {{console}}  net.ifnames=0 biosdevname=0 ksdevice=bootif kssendmac text BMP_MQ_HOST={{bmp_mq_host}} BMP_MQ_PORT={{bmp_mq_port}} BMP_MQ_USER={{bmp_mq_user}} BMP_MQ_PASSWORD='{{bmp_mq_password}}' BMP_MQ_VHOST={{bmp_mq_vhost}} BMP_MQ_EXCHANGE_ROUTING_KEY={{bmp_mq_exchange_routing_key}} BMP_IMAGE_HOST={{bmp_image_host}} BMP_IMAGE_PORT={{bmp_image_port}}  BMP_RSYSLOG_HSOT={{bmp_rsyslog_host}} BMP_RSYSLOG_PORT={{bmp_rsyslog_port}} 
    initrdefi {{initramfs_path}}
}