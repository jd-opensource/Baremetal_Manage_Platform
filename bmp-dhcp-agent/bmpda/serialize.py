import json
import uuid


class Serializable(object):
    """Base class for things that can be serialized."""
    serializable_fields = ()

    def serialize(self):
        """Turn this object into a dict."""
        return dict((f, getattr(self, f)) for f in self.serializable_fields
                    if (hasattr(self, f) and getattr(self, f) is not None))


class SerializableComparable(Serializable):
    """A Serializable class which supports some comparison operators

    This class supports the '__eq__' and '__ne__' comparison operators, but
    intentionally disables the '__hash__' operator as some child classes may be
    mutable.  The addition of these comparison operators is mainly used to
    assist with unit testing.
    """

    __hash__ = None

    def __eq__(self, other):
        return self.serialize() == other.serialize()

    def __ne__(self, other):
        return self.serialize() != other.serialize()


class DefaultJsonEncoder(json.JSONEncoder):
    """A slightly customized JSON encoder."""

    def __init__(self, **kwargs):
        super(DefaultJsonEncoder, self).__init__(**kwargs)

    def encode(self, o):
        """Turn an object into JSON.

        Appends a newline to responses when configured to pretty-print,
        in order to make use of curl less painful from most shells.
        """
        delimiter = ''

        # if indent is None, newlines are still inserted, so we should too.
        if self.indent is not None:
            delimiter = '\n'

        return super(DefaultJsonEncoder, self).encode(o) + delimiter

    def default(self, o):
        """Turn an object into a serializable object.

        In particular, by calling :meth:`.Serializable.serialize` on `o`.
        """
        if isinstance(o, Serializable):
            return o.serialize()
        elif isinstance(o, uuid.UUID):
            return str(o)
        else:
            return json.JSONEncoder.default(self, o)


class DefaultJsonDecoder(json.JSONDecoder):
    """Customized JSON decoder."""

    def __init__(self, **kwargs):
        super(DefaultJsonDecoder, self).__init__(**kwargs)
