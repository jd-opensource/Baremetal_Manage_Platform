from bmpa.command.common import Command
from bmpa.service import partitions
import bmpa.utils.validates as validates


class MakePartitions(Command):
    """Make partition.

    :param auto_create_efi:       Automatically create the UEFI partition in UEFI mode.
    :param auto_create_bios_grub: Automatically create a 1M BIOS_GRUB partition in BIOS mode with GPT.
    """

    def __init__(self,
                 volumes,
                 boot_mode=None,
                 auto_create_efi=True,
                 auto_create_bios_grub=True,
                 **kwargs):
        super(MakePartitions, self).__init__(**kwargs)
        self.volumes = volumes
        self.boot_mode = boot_mode
        self.auto_create_efi = auto_create_efi
        self.auto_create_bios_grub = auto_create_bios_grub
        validates.is_required(volumes, 'volumes')

        self.res_volumes = None

    def execute(self):
        self.res_volumes = partitions.make_partition(
            self.volumes, self.boot_mode, self.auto_create_efi, self.auto_create_bios_grub)

    def execute_after(self):
        self.data["volumes"] = self.res_volumes
