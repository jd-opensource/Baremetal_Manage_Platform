from bmpa.command.common import Command
from bmpa.service import image
from bmpa.utils import validates


class WriteImage(Command):

    def __init__(self,
                 url,
                 format,
                 hash,
                 filename,
                 **kwargs):
        super(WriteImage, self).__init__(**kwargs)

        self.url = url
        self.format = format
        self.hash = hash
        self.filename = filename
        self.res = None

        validates.is_required(url, 'url')
        validates.is_required(format, 'format')
        validates.is_required(hash, 'hash')
        validates.is_required(filename, 'filename')

    def execute(self):
        self.res = image.write_image(url=self.url, format=self.format,
                                     hash=self.hash, filename=self.filename)

    def execute_after(self):
        self.data["root_device_hints"] = self.res['root_device_hints']
        if "data_device_hints" in self.res:
            self.data["data_device_hints"] = self.res['data_device_hints']


class WriteImageTar(Command):

    def __init__(self,
                 url,
                 format,
                 hash,
                 filename,
                 **kwargs):
        super(WriteImageTar, self).__init__(**kwargs)

        self.url = url
        self.format = format
        self.hash = hash
        self.filename = filename
        self.res = None

        validates.is_required(url, 'url')
        validates.is_required(format, 'format')
        validates.is_required(hash, 'hash')
        validates.is_required(filename, 'filename')

    def execute(self):
        self.res = image.write_image_tar(url=self.url, format=self.format,
                                         hash=self.hash, filename=self.filename)

    def execute_after(self):
        self.data["root_device_hints"] = self.res['root_device_hints']
        if "data_device_hints" in self.res:
            self.data["data_device_hints"] = self.res['data_device_hints']
