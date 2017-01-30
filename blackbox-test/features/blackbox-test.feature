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

Feature: blackbox-test task

  As a developer, I want to be confused by using the blackbox-test task to blackbox test the blackbox-test task

  Scenario: with passing blackbox-test
    Given a file named "docker-compose.yml" with:
    """yaml
    blackbox-test:
      image: busybox
      command: /bin/true
    """
    And a file named "lc.yml" with:
    """yaml
    name: testsmokes
    """
    When I run `lc blackbox-test`
    Then it should succeed

  Scenario: with passing bbtest alias
    Given a file named "docker-compose.yml" with:
    """yaml
    blackbox-test:
      image: busybox
      command: /bin/true
    """
    And a file named "lc.yml" with:
    """yaml
    name: testsmokes
    """
    When I run `lc bbtest`
    Then it should succeed

  Scenario: with failing blackbox-test
    Given a file named "docker-compose.yml" with:
    """yaml
    blackbox-test:
      image: busybox
      command: /bin/false
    """
    And a file named "lc.yml" with:
    """yaml
    name: testsmokes
    """
    When I run `lc blackbox-test`
    Then it should fail

  Scenario: forwarding arguments
    Given a file named "docker-compose.yml" with:
    """yaml
    blackbox-test:
      image: busybox
      entrypoint: echo
    """
    And a file named "lc.yml" with:
    """yaml
    name: testsmokes
    """
    When I run `lc blackbox-test fdsa`
    Then it should succeed with "fdsa"

  ## see US7549
  Scenario: with defaulting to first running lc package
    Given I run `docker rmi -f elsyblackbox_docker_artifact_blackbox`
    And a file named "docker-compose.yml" with:
    """yaml
    blackbox-test:
      image: elsyblackbox_docker_artifact_blackbox
      command: /bin/true
    """
    And a file named "Dockerfile" with:
    """
    FROM library/alpine
    """
    And a file named "lc.yml" with:
    """yaml
    docker_image_name: elsyblackbox_docker_artifact_blackbox
    """
    When I run `lc blackbox-test --skip-package`
    Then it should fail with 'image library/elsyblackbox_docker_artifact_blackbox:latest not found'
    When I run `lc blackbox-test`
    Then it should succeed
    And the output should contain all of these:
      | Running package before executing blackbox tests |
      | Successfully built                              |

  Scenario: with a package service and a test service, should not run tests
    Given a file named "docker-compose.yml" with:
    """yaml
    blackbox-test:
      image: elsyblackbox_docker_artifact_blackbox
      command: /bin/true
    package:
      image: busybox
      command: echo foo
    test:
      image: busybox
      command: /bin/false
    """
    And a file named "lc.yml" with:
    """yaml
    name: testpackage
    """
    When I run `lc blackbox-test`
    Then it should succeed with "Running package before executing blackbox tests"
    And the output should not contain "Running tests before packaging"

    ## Only works in docker 1.11 and higher
  Scenario: with a Docker project and image labels
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: busybox
      command: /bin/true
    blackbox-test:
      image: elsyblackbox_docker_artifact_blackbox_labels
      command: /bin/true
    package:
      image: busybox
      command: /bin/true
    """
    And a file named "Dockerfile" with:
    """
    FROM alpine
    """
    And a file named "lc.yml" with:
    """yaml
    docker_image_name: elsyblackbox_docker_artifact_blackbox_labels
    docker_registry: localhost:5000
    """
    And I run `lc blackbox-test --git-commit=d8dfd9f`
    Then it should succeed
    And the output should contain all of these:
      | Attaching image label: com.elsy.metadata.git-commit=d8dfd9f                                   |
    And the image 'elsyblackbox_docker_artifact_blackbox_labels' should exist
    And it should have the following labels:
      | com.elsy.metadata.git-commit:d8dfd9f |
