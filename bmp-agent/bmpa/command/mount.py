import logging


from bmpa.command.common import Command
from bmpa.service import mount
import bmpa.utils.validates as validates

LOG = logging.getLogger(__name__)


class MountPartitions(Command):
    def __init__(self, mounts, auto_mount_efi=True, **kwargs):
        self.mounts = mounts
        self.auto_mount_efi = auto_mount_efi
        super(MountPartitions, self).__init__(**kwargs)

        validates.is_required(mounts, 'mounts')

    def execute(self):
        mount.mount_partitions(
            auto_mount_efi=self.auto_mount_efi, mounts=self.mounts)
