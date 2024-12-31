def load_class(class_path):
    parts = class_path.rsplit('.', 1)
    module = __import__(parts[0], fromlist=parts[1])
    return getattr(module, parts[1])
