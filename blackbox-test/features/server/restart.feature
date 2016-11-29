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

Feature: server restart task

  @teardown
  Scenario: with no server running
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server restart`
    Then it should succeed with 'starting service \"devserver\"'

  @teardown
  Scenario: with a dev server running
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server start`
    When I run `lc server restart`
    Then it should succeed
    And the output should contain all of these:
      | Stopping server                |
      | starting service \"devserver\" |

  @teardown
  Scenario: with a prod server running
    Given a file named "docker-compose.yml" with:
    """yaml
    prodserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server start -prod`
    When I run `lc server restart`
    Then it should succeed
    And the output should contain all of these:
      | Stopping server                |
      | starting service \"prodserver\" |
