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

Feature: server stop task

  @teardown
  Scenario: stopping a running devserver
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server start && lc server stop`
    Then it should succeed with "Stopping server"

  @teardown
  Scenario: stopping a running prodserver
    Given a file named "docker-compose.yml" with:
    """yaml
    prodserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server start -prod && lc server stop`
    Then it should succeed with "Stopping server"

  @teardown
  Scenario: stopping with no server services
    Given a file named "docker-compose.yml" with:
    """yaml
    """
    When I run `lc server stop`
    Then it should succeed
