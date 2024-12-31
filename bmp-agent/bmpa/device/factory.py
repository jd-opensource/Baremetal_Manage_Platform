from bmpa.device.default import DefaultDevice
from bmpa.device.sogon import SogonW580G20


def get_device(manufacturer, product_name):
    return _get_device(manufacturer, product_name)


def _get_device(manufacturer, product_name):
    if manufacturer == SogonW580G20.MANUFACTURER and product_name == SogonW580G20.PRODUCT_NAME:
        return SogonW580G20()

    return DefaultDevice()
