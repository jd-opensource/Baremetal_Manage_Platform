import os

from bmpa import constants
from bmpa.utils import utils

cmdline_path: str = os.getenv('BMD_CMDLINE_PATH', '/proc/cmdline')
AGENT_PARAMS: dict = utils.get_agent_params(cmdline_path)
heart_period: int = int(AGENT_PARAMS.get('BMP_HEART_PERIOD', 10))
mq_host: str = AGENT_PARAMS.get('BMP_MQ_HOST')  # required
mq_port: int = int(AGENT_PARAMS.get('BMP_MQ_PORT', 5672))
mq_username: str = AGENT_PARAMS.get('BMP_MQ_USER')  # required
mq_password: str = AGENT_PARAMS.get('BMP_MQ_PASSWORD')  # required
mq_password = mq_password.strip("'")
mq_is_use_ssl: int = int(AGENT_PARAMS.get('BMP_MQ_IS_USE_SSL',
                                          0))  # 0:not use 1:use
mq_vhost: str = AGENT_PARAMS.get('BMP_MQ_VHOST', '/bmp')
mq_exchange: str = AGENT_PARAMS.get('BMP_MQ_EXCHANGE', 'CPS_IRONIC_SCHEDULER')
mq_exchange_type: str = AGENT_PARAMS.get('BMP_MQ_EXCHANGE_TYPE', 'direct')
image_host: str = AGENT_PARAMS.get('BMP_IMAGE_HOST')
image_port: int = int(AGENT_PARAMS.get('BMP_IMAGE_PORT', 10000))
protocol: str = AGENT_PARAMS.get('BMP_PROTOCAL', constants.PROTOCAL_RABBITMQ)
log_path: str = AGENT_PARAMS.get('BMP_LOG_PATH',
                                 '/var/log/bmp/bmp-agent.log')
# DEBUG,INFO,WARNING
logging_level = AGENT_PARAMS.get('LOGGING_LEVEL', 'INFO')
rsyslog_host = AGENT_PARAMS.get('BMP_RSYSLOG_HSOT', mq_host)
rsyslog_port = AGENT_PARAMS.get('BMP_RSYSLOG_PORT', 1514)
