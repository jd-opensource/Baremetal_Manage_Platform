import hashlib
import os


def ensure_dir(path, mode=None):
    if path == "":
        path = "."
    if not os.path.exists(path):
        os.makedirs(path)

    if mode is not None:
        os.chmod(path, mode)


def md5sum(path):
    hashlib_md5 = hashlib.md5()
    with open(path, 'rb') as f:
        for chunk in iter(lambda: f.read(4096), b""):
            hashlib_md5.update(chunk)
    return hashlib_md5.hexdigest()
