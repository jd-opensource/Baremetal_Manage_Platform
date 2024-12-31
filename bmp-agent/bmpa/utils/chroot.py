from bmpa.utils import executor
from bmpa.utils import mountutils
from bmpa.utils import paths


class ChrootTarget:
    def __init__(self, target,
                 mounts=None):
        if target is None:
            target = "/"
        self.target = paths.target_path(target)
        if mounts is not None:
            self.mounts = mounts
        else:
            self.mounts = ["/dev", "/proc", "/run", "/sys"]
        self.umounts = []

    def __enter__(self):
        for p in self.mounts:
            tpath = paths.target_path(self.target, p)
            mount_success = True
            if p == '/sys':
                mount_success = mountutils.do_mount(
                    'none', tpath, opts=['-t', 'sysfs'])
            else:
                mount_success = mountutils.do_mount(p, tpath, opts='--bind')
            if mount_success:
                self.umounts.append(tpath)
        return self

    def __exit__(self, etype, value, trace):

        for p in reversed(self.umounts):
            mountutils.do_umount(p)

    def execute(self, *args, **kwargs):
        kwargs['target'] = self.target
        return executor.execute(*args, **kwargs)
