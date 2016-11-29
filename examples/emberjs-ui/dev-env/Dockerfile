# Copyright 2016 Cisco Systems, Inc.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
# http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM library/node:6.3

ENV EMBER_CLI_VERSION=2.7

RUN npm install -g ember-cli@$EMBER_CLI_VERSION \
    bower \
    phantomjs

## watchman is used for live-reloading server
ENV WATCHMAN_VERSION=v3.9.0
RUN \
  cd /tmp &&\
	git clone https://github.com/facebook/watchman.git &&\
	cd watchman &&\
	git checkout $WATCHMAN_VERSION &&\
	./autogen.sh &&\
	./configure --without-python &&\
	make -j2 &&\
	make install &&\
  cd / &&\
  rm -rf /tmp/watchman
