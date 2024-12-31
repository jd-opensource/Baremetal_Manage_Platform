import logging
import threading
import time

from bmpda import agent
from bmpda import config
from bmpda import log

LOG = logging.getLogger('bmpda')

ascii_snake = """\
=============================================================================================

    --..,_                     _,.--.
       `'.'.                .'`__ o  `;__.
          '.'.            .'.'`  '---'`  `  BMP Agent
            '.`'--....--'`.'
              `'--....--'`

=============================================================================================
"""


def run():
    print(ascii_snake)
    log.init()
    LOG.debug("Env %s.", config.env)
    eng = agent.Agent()
    eng.start()
    LOG.info("Agent is running.")
    time.sleep(threading.TIMEOUT_MAX)


if __name__ == '__main__':
    run()
