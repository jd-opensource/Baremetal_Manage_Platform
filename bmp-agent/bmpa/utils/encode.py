def to_utf8(text):
    """Encode Unicode to UTF-8, return bytes unchanged.

    Raise TypeError if text is not a bytes string or a Unicode string.
    """
    if isinstance(text, bytes):
        return text
    elif isinstance(text, str):
        return text.encode('utf-8')
    else:
        raise TypeError('bytes or Unicode expected, got %s' %
                        type(text).__name__)
