import os

from bmpda import constants
from bmpda.utils import utils

env: str = os.getenv('ENV')
if env:
    cfg_path = os.path.join('cfg', env)
    AGENT_PARAMS: dict = utils.get_agent_params(
        utils.get_relative_path(cfg_path))
else:
    AGENT_PARAMS = os.environ
heart_period: int = int(AGENT_PARAMS.get('BMP_HEART_PERIOD', 10))
mq_host: str = AGENT_PARAMS.get('BMP_MQ_HOST')  # required
mq_port: int = int(AGENT_PARAMS.get('BMP_MQ_PORT', 5672))
mq_username: str = AGENT_PARAMS.get('BMP_MQ_USER')  # required
mq_password: str = AGENT_PARAMS.get('BMP_MQ_PASSWORD')  # required
mq_password = mq_password.strip("'")
mq_vhost: str = AGENT_PARAMS.get('BMP_MQ_VHOST', '/bmp')
mq_exchange: str = AGENT_PARAMS.get('BMP_MQ_EXCHANGE', 'BMP_SCHEDULER')
mq_exchange_type: str = AGENT_PARAMS.get('BMP_MQ_EXCHANGE_TYPE', 'topic')
mq_exchange_type_publish: str = AGENT_PARAMS.get(
    'BMP_MQ_EXCHANGE_TYPE_PUBLISH', 'direct')
mq_exchange_routing_key: str = AGENT_PARAMS.get(
    'BMP_MQ_EXCHANGE_ROUTING_KEY', 'az')
mq_exchange_consumer: str = AGENT_PARAMS.get(
    'BMP_MQ_EXCHANGE_CONSUMER', 'BMP_SCHEDULER_TOPIC')
mq_queue_name_pefix_consumer: str = AGENT_PARAMS.get(
    'BMP_MQ_QUEUE_NAME_PEFIX_CONSUMER', 'dhcp-az')
omapi_host = AGENT_PARAMS.get('BMP_OMAPI_HOST', '127.0.0.1')
omapi_port = AGENT_PARAMS.get('BMP_OMAPI_PORT', 7911)
omapi_key = AGENT_PARAMS.get('BMP_OMAPI_KEY', '')
dhcp_config_dir = AGENT_PARAMS.get('BMP_DHCP_CONFIG_DIR', '/data')
dhcp_control_bin_path = utils.get_relative_path('shell/dhcpd_control.sh')
dhcp_control_bin = AGENT_PARAMS.get(
    'BMP_DHCP_CONTROL_BIN', dhcp_control_bin_path)
protocol: str = AGENT_PARAMS.get('BMP_PROTOCAL', constants.PROTOCAL_RABBITMQ)
log_path: str = AGENT_PARAMS.get('BMP_LOG_PATH',
                                 '/var/log/bmp/bmp-dhcp-agent/bmp-dhcp-agent.log')
# DEBUG,INFO,WARNING
logging_level = AGENT_PARAMS.get('LOGGING_LEVEL', 'INFO')

# tftp
tftp_config_dir = AGENT_PARAMS.get('BMP_TFTP_CONFIG_DIR', '/data')
tftp_mq_host: str = AGENT_PARAMS.get('BMP_TFTP_MQ_HOST')
tftp_mq_port: int = int(AGENT_PARAMS.get('BMP_TFTP_MQ_PORT', 5672))
tftp_mq_user: str = AGENT_PARAMS.get('BMP_TFTP_MQ_USER')
tftp_mq_password: str = AGENT_PARAMS.get('BMP_TFTP_MQ_PASSWORD')
tftp_mq_password = tftp_mq_password.strip("'")
tftp_mq_vhost: str = AGENT_PARAMS.get('BMP_TFTP_MQ_VHOST', '/bmp')
tftp_mq_exchange_routing_key: str = AGENT_PARAMS.get(
    'BMP_TFTP_MQ_EXCHANGE_ROUTING_KEY', '')
tftp_image_host: str = AGENT_PARAMS.get('BMP_TFTP_IMAGE_HOST', '')
tftp_image_port: str = AGENT_PARAMS.get('BMP_TFTP_IMAGE_PORT', '')
tftp_rsyslog_host: str = AGENT_PARAMS.get('BMP_TFTP_RSYSLOG_HSOT', '')
tftp_rsyslog_port: str = AGENT_PARAMS.get('BMP_TFTP_RSYSLOG_PORT', '')
