import logging
import threading
import time


from bmpa import config
from bmpa import engine
from bmpa import log
from bmpa.utils import executor
from bmpa.utils import interfaceutils
from bmpa.utils import system

LOG = logging.getLogger('bmpa')

ascii_snake = """\
=============================================================================================

    --..,_                     _,.--.
       `'.'.                .'`__ o  `;__.
          '.'.            .'.'`  '---'`  `  BMP Agent
            '.`'--....--'`.'
              `'--....--'`

=============================================================================================
"""


def set_rsyslog_conf():
    try:
        rsyslog_tpl_path = system.get_relative_path(
            'template/rsyslog/bmp-agent.conf.tpl')
        rsyslog_config = system.render_template(rsyslog_tpl_path, {
                                                'log_path': config.log_path,
                                                'rsyslog_host': config.rsyslog_host,
                                                'rsyslog_port': config.rsyslog_port})
        rsyslog_path = '/etc/rsyslog.d/bmp-agent.conf'
        with open(rsyslog_path, 'w') as f:
            f.write(rsyslog_config)
        executor.execute(*['systemctl', 'restart', 'rsyslog'])
    except Exception as e:
        LOG.error("Failed to set rsyslog config: %s", e)


def run():
    print(ascii_snake)
    log.init()
    LOG.debug("Cmdline path %s.", config.cmdline_path)
    set_rsyslog_conf()
    interfaceutils.dhclient()
    interfaceutils.wait_for_dhcp()
    eng = engine.Engine()
    eng.start()
    LOG.info("Engine is running.")
    time.sleep(threading.TIMEOUT_MAX)


if __name__ == '__main__':
    run()
