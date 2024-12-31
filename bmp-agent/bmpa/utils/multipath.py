from bmpa.utils import udev


def mpath_partition_to_mpath_id_and_partnumber(devpath):
    """Return the mpath id and partition number of a multipath partition. """
    info = udev.udevadm_info(devpath)
    if 'DM_MPATH' in info and 'DM_PART' in info:
        return info['DM_MPATH'], info['DM_PART']

    return None
