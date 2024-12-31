from typing import Dict


def is_required(arg: object, param_name):
    not_null(arg, param_name)
    not_empty(arg, param_name)


def check_field(dict_obj: Dict, field: str):
    not_field(dict_obj, field)
    not_null(dict_obj[field], field)
    not_empty(dict_obj[field], field)


def not_field(dict_obj: Dict, field: str):
    if field not in dict_obj:
        raise ValueError('Missing field: [{0}]'.format(field))


def not_null(arg: object, param_name: str):
    if arg is None:
        raise ValueError("Parameter `{param_name}` cannot not be None".format(
            param_name=param_name))


def not_empty(arg: object, param_name: str):
    if arg == '':
        raise ValueError("Parameter `{param_name}` cannot be empty".format(
            param_name=param_name))


def sn_match(sn, sn_now):
    if sn != sn_now:
        raise ValueError("Sn `{sn}` not match `{sn_now}`".format(
            sn=sn, sn_now=sn_now))
