from bmpa.utils import utils


class Disk:

    def __init__(self,
                 serial: str = None,
                 enclosure: int = None,
                 slot: int = None,
                 pd_type: str = None,
                 media_type: str = None,
                 raw_size: float = None,
                 raw_size_unit: str = None,
                 size: float = None,
                 size_unit: str = None,
                 inquiry_data: str = None,
                 device_firware_level: str = None,
                 device_path: str = None):
        self.serial = serial
        self.enclosure = enclosure
        self.slot = slot
        self.pd_type = pd_type
        self.media_type = media_type
        self.raw_size, self.raw_size_unit = utils.convert_disk_size(
            raw_size, raw_size_unit)
        self.size, self.size_unit = utils.convert_disk_size(size, size_unit)
        self.inquiry_data = inquiry_data
        self.device_firware_level = device_firware_level
        self.device_path = device_path

    def __str__(self):
        return (f"Disk(serial='{self.serial}',"
                f"enclosure='{self.enclosure}',"
                f"slot='{self.slot}',"
                f"pd_type='{self.pd_type}',"
                f"media_type='{self.media_type}',"
                f"raw_size='{self.raw_size}',"
                f"raw_size_unit='{self.raw_size_unit}',"
                f"size='{self.size}',"
                f"size_unit='{self.size_unit}',"
                f"inquiry_data='{self.inquiry_data}',"
                f"device_firware_level='{self.device_firware_level}',"
                f"device_path='{self.device_path}')")
