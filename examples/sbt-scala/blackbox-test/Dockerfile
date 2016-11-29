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

FROM library/ruby:latest

EXPOSE 9090

RUN apt-get update && apt-get install -y netcat

ENV PHANTOMJS_VERSION 2.1.1
RUN \
  apt-get update && \
  apt-get install -y libfreetype6 libfontconfig && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/* && \
  mkdir -p /srv/var && \
  curl -sL https://bitbucket.org/ariya/phantomjs/downloads/phantomjs-$PHANTOMJS_VERSION-linux-x86_64.tar.bz2 | tar -C /tmp -xj && \
  mv /tmp/phantomjs-$PHANTOMJS_VERSION-linux-x86_64/ /opt/phantomjs && \
  ln -s /opt/phantomjs/bin/phantomjs /usr/local/bin/phantomjs

RUN gem install \
    turnip \
    poltergeist \
    gnawrnip \
    rspec
