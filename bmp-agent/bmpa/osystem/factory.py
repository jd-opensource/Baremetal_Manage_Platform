import re

from bmpa.osystem.centos import CentOS7X
from bmpa.osystem.centos import CentOS8X
from bmpa.osystem.debian import Debian10x
from bmpa.osystem.default import DefaultOsystem
from bmpa.osystem.kylin import Kylin10X
from bmpa.osystem.loongnix import Loongnix8X
from bmpa.osystem.openEuler import OpenEuler20X
from bmpa.osystem.ubuntu import Ubuntu18X
from bmpa.osystem.ubuntu import Ubuntu20X
from bmpa.osystem.uniontech import UnionTech20X
from bmpa.osystem.windows import Windows10
from bmpa.osystem.windows import Windows2012
from bmpa.osystem.windows import Windows2016
from bmpa.osystem.windows import Windows2019


def get_osystem(os_type, os_version):
    return _get_osystem(os_type, os_version)


def _get_osystem(os_type, os_version):
    if os_type.lower() == "centos":
        if re.compile(r"7\.*").match(os_version):
            return CentOS7X()
        return CentOS8X()
    elif os_type.lower() == "ubuntu":
        if re.compile(r"18\.*").match(os_version):
            return Ubuntu18X()
        return Ubuntu20X()
    elif os_type.lower() == "kylin":
        return Kylin10X()
    elif os_type.lower() == "openeuler":
        return OpenEuler20X()
    elif os_type.lower() == "uniontech":
        return UnionTech20X()
    elif os_type.lower() == "debian":
        return Debian10x()
    elif os_type.lower() == "loongnix":
        return Loongnix8X()
    elif os_type.lower() == "windowspc":
        return Windows10()
    elif os_type.lower() == "windows":
        if re.compile(r"2012").match(os_version):
            return Windows2012()
        elif re.compile(r"2016").match(os_version):
            return Windows2016()
        return Windows2019()

    return DefaultOsystem()
