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

Feature: test task

  Scenario: with passing tests
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: busybox
      command: /bin/true
    """
    And a file named "lc.yml" with:
    """yaml
    name: test
    """
    When I run `lc test`
    Then it should succeed

  Scenario: with failing tests
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: busybox
      command: /bin/false
    """
    And a file named "lc.yml" with:
    """yaml
    name: test
    """
    When I run `lc test`
    Then it should fail

  Scenario: forwarding arguments
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: busybox
      entrypoint: echo
    """
    And a file named "lc.yml" with:
    """yaml
    name: test
    """
    When I run `lc test fdsa`
    Then it should succeed with "fdsa"
