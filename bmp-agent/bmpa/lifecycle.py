import logging
from threading import Condition

LOG = logging.getLogger(__name__)

STOPPED = "STOPPED"
FAILED = "FAILED"
STARTING = "STARTING"
STARTED = "STARTED"
STOPPING = "STOPPING"
RUNNING = "RUNNING"


class LifeCycleException(Exception):

    def __init__(self, message):
        super(LifeCycleException, self).__init__(message)


class LifeCycle(object):

    def __init__(self):
        self.lock = Condition()
        self.state = STOPPED

    def start(self):
        with self.lock:
            try:
                if self.state == STARTED or self.state == STARTING:
                    return
                self.state = STARTING
                LOG.info("%s is starting.", self.__class__.__name__)
                self.do_start()
                self.state = STARTED
                LOG.info("%s is started.", self.__class__.__name__)
            except Exception as e:
                self.set_failed(e)
                LOG.info("%s started failed.", self.__class__.__name__)
                raise

    def stop(self):
        with self.lock:
            try:
                if self.state == STOPPING or self.state == STOPPED:
                    return
                self.state = STOPPING
                LOG.info("%s is stopping.", self.__class__.__name__)
                self.lock.notify()
                self.do_stop()
                self.state = STOPPED
                LOG.info("%s is stopped.", self.__class__.__name__)
            except Exception as e:
                self.set_failed(e)
                LOG.info("%s stop failed.", self.__class__.__name__)
                raise

    def is_running(self):
        return self.state == STARTED or self.state == STARTING

    def is_started(self):
        return self.state == STARTED

    def is_starting(self):
        return self.state == STARTING

    def is_stoping(self):
        return self.state == STOPPING

    def is_stopped(self):
        return self.state == STOPPED

    def is_failed(self):
        return self.state == FAILED

    def do_start(self):
        raise LifeCycleException("%s not impl doStart method." %
                                 self.__class__.__name__)

    def do_stop(self):
        raise LifeCycleException("%s not impl doStart method." %
                                 self.__class__.__name__)

    def set_failed(self, exception):
        self.state = FAILED
        LOG.error("err:%s", str(exception))
        LOG.error("%s", __name__, exc_info=1)

    def get_state(self):
        return self.state

    def should_stop(self):
        return self.state in [STOPPING, STOPPED, FAILED]
