from bmpa.device.default import DefaultDevice


class Sogon(DefaultDevice):
    def __init__(self):
        super(Sogon, self).__init__()

    def grub_ttys(self):
        return 'ttyS1'


class SogonW580G20(DefaultDevice):
    MANUFACTURER = "Sugon"
    PRODUCT_NAME = "W580-G20 (Default string)"

    def __init__(self):
        super(Sogon, self).__init__()

    def grub_ttys(self):
        return 'ttyS1'
