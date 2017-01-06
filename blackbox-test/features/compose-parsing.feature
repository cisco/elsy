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


Feature: elsy supports parsing docker-compose yaml files

Scenario: using the v2 extended build syntax
  And a file named "docker-compose.yml" with:
  """yaml
  version: '2'
  services:
    test:
      build:
        context: .
  """
  When I run `lc dc config`
  Then it should succeed
