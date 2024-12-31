from bmpda import serialize


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


class InvalidArgumentError(RESTError):

    message = 'Invalid Argument error.'

    def __init__(self, details):
        super(InvalidArgumentError, self).__init__(details)


class BmpError(RESTError):
    """Error bmp error."""

    message = 'Bmp error.'

    def __init__(self, details):
        super(BmpError, self).__init__(details)
