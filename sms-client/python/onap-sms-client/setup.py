# Copyright 2018 Intel Corporation, Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


from setuptools import setup, find_packages

setup(

    name='onap-sms-client',
    keywords=("secret", "consul", "onap"),
    description="ONAP python SMS client library",
    long_description="python-package onap-sms-client client library for using"
                     " Secret Management Service (SMS) inside ONAP. Refer "
                     "https://wiki.onap.org/display/DW/Secret+Management+Service for more details.",
    version="0.0.1",
    url="https://gerrit.onap.org/r/gitweb?p=aaf%2Fsms.git;a=summary",
    license="Apache 2",
    author="Kiran Kamineni",
    packages=find_packages(),
    platforms=["all"],
    classifiers=[
        "Intended Audience :: Developers",
        "Programming Language :: Python :: 2.7"
    ]
)
