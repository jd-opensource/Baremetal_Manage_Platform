import functools
import logging
import os

import bmpa.config as config
import bmpa.utils.file as file

app_name = 'bmpa'


def _setupLogging(name='bmpa'):
    log_path = config.log_path
    log_dir = os.path.dirname(log_path)
    file.ensure_dir(log_dir)

    logger = logging.getLogger(name)
    logger.setLevel(logging.getLevelName(config.logging_level))
    formatter = logging.Formatter(
        '%(asctime)s %(name)s/%(filename)s[line:%(lineno)d] %(levelname)s %(message)s')

    file_handler = logging.FileHandler(log_path, encoding='utf-8')
    file_handler.setFormatter(formatter)

    stream_handler = logging.StreamHandler()
    stream_handler.setFormatter(formatter)

    logger.addHandler(file_handler)
    logger.addHandler(stream_handler)


def init(name='bmpa'):
    global app_name
    app_name = name
    _setupLogging(name)


def log_complete(func):

    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        result = func(*args, **kwargs)
        logger = logging.getLogger(app_name)
        logger.info("%s complete.", func.__name__)
        return result

    return wrapper
