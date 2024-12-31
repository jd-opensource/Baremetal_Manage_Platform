from bmpa.utils import executor

_SEP = '='


def split_lvm_name(full):
    """split full lvm name into tuple of (volgroup, lv_name)"""
    # 'dmsetup splitname' is the authoratative source for lvm name parsing
    stdout, _ = executor.execute(*[
        'dmsetup', 'splitname', full, '-c', '--noheadings', '--separator',
        _SEP, '-o', 'vg_name,lv_name'
    ])
    return stdout.strip().split(_SEP)


def _filter_lvm_info(lvtool, match_field, query_field, match_key, args=None):
    """filter output of pv/vg/lvdisplay tools"""
    if args is None:
        args = []
    stdout, _ = executor.execute(*[
        lvtool, '-C', '--separator', _SEP, '--noheadings', '-o', ','.join(
            [match_field, query_field])
    ] + args)
    return [
        qf for (mf, qf) in
        [line.strip().split(_SEP) for line in stdout.strip().splitlines()]
        if mf == match_key
    ]


def get_lvols_in_volgroup(vg_name):
    """get logical volumes in volgroup"""
    return _filter_lvm_info('lvdisplay', 'vg_name', 'lv_name', vg_name)


def get_pvols_in_volgroup(vg_name):
    """get physical volumes used by volgroup"""
    return _filter_lvm_info('pvdisplay', 'vg_name', 'pv_name', vg_name)
