set timeout=1
set default=1
menuentry 'bmp agent x86_64 uefi' {
    linuxefi ${BMP_KERNEL_PATH_x86}  net.ifnames=0 biosdevname=0 ksdevice=bootif kssendmac text BMP_MQ_HOST=${BMP_HOST_IP} BMP_MQ_PORT=${BMP_MQ_PORT} BMP_MQ_USER=${BMP_MQ_USER} BMP_MQ_PASSWORD='${BMP_MQ_PASSWORD}' BMP_MQ_VHOST=${BMP_MQ_VHOST} BMP_MQ_EXCHANGE_ROUTING_KEY=${BMP_MQ_EXCHANGE_ROUTING_KEY} BMP_IMAGE_HOST=${BMP_HOST_IP} BMP_IMAGE_PORT=${BMP_IMAGE_PORT}
    initrdefi ${BMP_INITRAMFS_PATH_x86}
}
