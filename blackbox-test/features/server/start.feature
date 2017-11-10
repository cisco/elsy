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

Feature: server start task

  @teardown
  Scenario: starting a devserver
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server start`
    Then it should succeed
    And it should report a correct address

  @teardown
  Scenario: starting a prodserver
    Given a file named "docker-compose.yml" with:
    """yaml
    prodserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server start -prod`
    Then it should succeed
    And it should report a correct address

  @teardown
  Scenario: starting server with no devserver service defined
    Given a file named "docker-compose.yml" with:
    """yaml
    """
    When I run `lc server start`
    Then it should fail with 'no \"devserver\" service defined'

  ## this is mainly testing that this error condition provides a sensible error to the user
  Scenario: starting server with no image should attempt to pull the image
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: somefakeimagethatdoesntexist
      ports:
       - "80"
    """
    When I run `lc server start`
    Then it should fail
    And the output should contain one of the following:
      |image library/somefakeimagethatdoesntexist:latest not found                                   |
      |repository somefakeimagethatdoesntexist not found: does not exist or no pull access           |
      |pull access denied for somefakeimagethatdoesntexist, repository does not exist or may require |

  @teardown
  Scenario: starting prod server with no prodserver service defined
    Given a file named "docker-compose.yml" with:
    """yaml
    """
    When I run `lc server start -prod`
    Then it should fail with 'no \"prodserver\" service defined'

  @teardown
  Scenario: starting prod when devserver is already running
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: nginx
      ports:
       - "80"
    prodserver:
     image: nginx
     ports:
      - "80"
    """
    When I run `lc server start`
    And I run `lc server start -prod`
    Then it should succeed with '\"devserver\" already running'

  @teardown
  Scenario: starting dev when prodserver is already running
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: nginx
      ports:
       - "80"
    prodserver:
     image: nginx
     ports:
      - "80"
    """
    When I run `lc server start -prod`
    And I run `lc server start`
    Then it should succeed with '\"prodserver\" already running'
