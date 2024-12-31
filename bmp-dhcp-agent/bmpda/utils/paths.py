import os


def target_path(target, path=None):
    if target in (None, ""):
        target = "/"

    target = os.path.abspath(target)

    if not path:
        return target

    path = path.lstrip("/")

    return os.path.join(target, path)
