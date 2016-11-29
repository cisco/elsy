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

Feature: command execution


Scenario: with failing command
  Given a file named "docker-compose.yml" with:
  """yaml
  test:
    image: busybox
    command: /bin/false
  """
  And a file named "lc.yml" with:
  """yaml
  name: testcommand
  """
  When I run `lc test`
  Then it should fail
  When I run `lc --debug test`
  Then it should fail

Scenario: with successful command
  Given a file named "docker-compose.yml" with:
  """yaml
  test:
    image: busybox
    command: /bin/true
  """
  And a file named "lc.yml" with:
  """yaml
  name: testcommand
  """
  When I run `lc test`
  Then it should succeed
  When I run `lc --debug test`
  Then it should succeed

Scenario: with no command
  When I run `lc`
  Then it should succeed
