from bmpa import serialize


class RESTError(Exception, serialize.Serializable):
    """Base class for errors."""
    message = 'Internal Server Error.'
    details = 'An unexpected error occurred. Please try back later.'
    status_code = 500
    serializable_fields = ('type', 'code', 'message', 'details')

    def __init__(self, details: str = None, *args, **kwargs):
        super(RESTError, self).__init__(*args, **kwargs)
        self.type = self.__class__.__name__
        self.code = self.status_code
        if details:
            self.details = details

    def __str__(self):
        return "{}: {}".format(self.message, self.details)

    def __repr__(self):
        return "{}('{}')".format(self.__class__.__name__, self.__str__())


class UnkownCommandError(RESTError):
    """Error which occurs if not has the command."""
    message = 'Not found'
    status_code = 404
    details = 'The  Command was not found.'

    def __init__(self, details: str = None):
        if details:
            self.details = details
        super(UnkownCommandError, self).__init__(details)


class IncompatibleHardwareMethodError(RESTError):
    """Error raised when HardwareManager method incompatible with hardware."""

    message = 'HardwareManager method is not compatible with hardware.'

    def __init__(self, details: str = None):
        super(IncompatibleHardwareMethodError, self).__init__(details)


class BlockDeviceError(RESTError):
    """Error raised when a block devices causes an unknown error."""

    message = 'Block device caused unknown error.'

    def __init__(self, details):
        super(BlockDeviceError, self).__init__(details)


class BlockDeviceEraseError(RESTError):
    """Error raised when an error occurs erasing a block device."""

    message = 'Error erasing block device.'

    def __init__(self, details):
        super(BlockDeviceEraseError, self).__init__(details)


class DeviceNotFound(RESTError):
    status_code = 404
    """Error raised when the device to deploy the image onto is not found."""

    message = ('Error finding the disk or partition device to deploy '
               'the image onto')

    def __init__(self, details):
        super(DeviceNotFound, self).__init__(details)


class ImageDownloadError(RESTError):
    """Error raised when an image cannot be downloaded."""

    message = 'Error downloading image.'

    def __init__(self, image_id, msg):
        details = 'Download of image id {} failed: {}.'.format(image_id, msg)
        self.secondary_message = msg
        super(ImageDownloadError, self).__init__(details)


class ImageChecksumError(RESTError):
    """Error raised when an image fails to verify against its checksum."""

    message = 'Error verifying image checksum.'
    details_str = ('Image failed to verify against checksum. location: {}; '
                   'image ID: {}; image checksum: {}; verification '
                   'checksum: {}')

    def __init__(self, image_id, image_location, checksum,
                 calculated_checksum):
        details = self.details_str.format(image_location, image_id, checksum,
                                          calculated_checksum)
        super(ImageChecksumError, self).__init__(details)


class BmpError(RESTError):
    """Error bmp error."""

    message = 'Bmp error.'

    def __init__(self, details):
        super(BmpError, self).__init__(details)
