import logging
import sys
import time

LOG = logging.getLogger(__name__)


def download(url: str, path: str):
    """Download url to path"""
    if sys.version_info >= (3, ):
        import urllib.request as urllib2
    else:
        import urllib2

    buf_len = 8192
    info = None
    with open(path, 'wb') as f:
        try:
            buf = None
            fsize = 0
            start = time.time()
            last_time = start
            u = urllib2.urlopen(url)
            info = u.info()
            total_size = int(info.get('Content-Length', 1))
            while True:
                buf = u.read(buf_len)
                if not buf:
                    break
                f.write(buf)
                fsize += len(buf)
                time_delta = time.time() - last_time
                if time_delta >= 10:
                    progress = fsize / total_size * 100
                    LOG.info(
                        "Downloaded %d bytes from %s to %s in %.2fs (%.2fMbps), %.2f%%.", fsize,
                        url, path, time.time() - start, fsize / time_delta / 1024 / 1024, progress)
                    last_time = time.time()

            time_delta = time.time() - start
            LOG.info(
                "Downloaded %d bytes from %s to %s in %.2fs (%.2fMbps) success.", fsize,
                url, path, time_delta, fsize / time_delta / 1024 / 1024)
        except Exception as e:
            raise e
    return path, info
