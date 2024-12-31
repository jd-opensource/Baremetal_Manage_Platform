import io
import os
import sys

from setuptools import setup, find_packages

NAME = "bmpda"
DESCRIPTION = "BMP DHCP Agent"
URL = ""
EMAIL = ""
AUTHOR = "JD Cloud Pysical System"
REQUIRES_PYTHON = ">=3.6"
VERSION = "2.0.1"

here = os.path.abspath(os.path.dirname(__file__))


# What packages are required for this module to be executed
def read_requires():
    with open(os.path.join(here, "requirements.txt")) as f:
        return f.read().splitlines()


setup(
    name=NAME,
    version=VERSION,
    description=DESCRIPTION,
    author=AUTHOR,
    author_email=EMAIL,
    platforms="any",
    python_requires=REQUIRES_PYTHON,
    url=URL,
    packages=find_packages(
        exclude=["*.tests", "*.tests.*", "tests.*", "tests"]),
    entry_points={
        "console_scripts": [
            "bmp-agent = bmpda.cmd.agent:run",
        ],
    },
    install_requires=read_requires(),
    # These files are included via the MANIFEST.in will be
    # automatically installed with your package
    include_package_data=True,
    ## license="",
    zip_safe=False)
