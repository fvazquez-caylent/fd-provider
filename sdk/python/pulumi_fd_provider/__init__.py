# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_fd_provider.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_fd_provider.config')

_utilities.register(
    resource_modules="""
[]
""",
    resource_packages="""
[
 {
  "pkg": "fd-provider",
  "token": "pulumi:providers:fd-provider",
  "fqn": "pulumi_fd_provider",
  "class": "Provider"
 }
]
"""
)