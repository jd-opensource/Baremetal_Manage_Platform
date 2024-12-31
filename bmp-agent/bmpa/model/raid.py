class RaidConfig:

    def __init__(self,
                 adapter_id: str,
                 raid_level: str,
                 physical_disks: list
                 ):
        self.adapter_id = adapter_id
        self.raid_level = raid_level
        self.physical_disks = physical_disks

    def __str__(self):
        return (f"RaidConfig(adapter_id='{self.adapter_id}',"
                f"raid_level='{self.raid_level}',"
                f"physical_disks='{self.physical_disks}')")
