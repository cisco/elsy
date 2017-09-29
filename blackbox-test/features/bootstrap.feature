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

Feature: bootstrap task

  Scenario: with valid image based services
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: busybox
    """
    And a file named "lc.yml" with:
    """yaml
    name: testbootstrap
    """
    When I run `lc bootstrap`
    Then it should succeed with "Pulling test"
    When I run `lc --disable-parallel-pull bootstrap`
    Then it should succeed with "Pulling from library/busybox"

  Scenario: with invalid images
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: fdsafdsa
    """
    And a file named "lc.yml" with:
    """yaml
    name: testbootstrap
    """
    When I run `lc bootstrap`
    Then it should fail pulling "fdsafdsa"

  Scenario: with a build service
    Given a file named "dev-env/Dockerfile" with:
    """
    FROM busybox
    """
    And a file named "docker-compose.yml" with:
    """yaml
    test:
      build: dev-env
    """
    And a file named "lc.yml" with:
    """yaml
    name: testbootstrap
    """
    When I run `lc bootstrap`
    Then it should succeed with "Building test"

  Scenario: with an invalid build service
    Given a file named "dev-env/Dockerfile" with:
    """
    FRO busybox
    """
    And a file named "docker-compose.yml" with:
    """yaml
    test:
      build: dev-env
    """
    And a file named "lc.yml" with:
    """yaml
    name: testbootstrap
    """
    When I run `lc bootstrap`
    Then it should fail with "Service 'test' failed to build"

  Scenario: with an image matching the docker artifact
    It is common to utilize the project's docker image artifact in a docker
    compose service. When docker-compose attempts to pull that service, it will
    produce an error. In order to minimize developer confusion. That error
    should be squelched while still showing errors where services pull images
    which are not expected
    Given a file named "docker-compose.yml" with:
    """yaml
    prodserver:
      image: baz
    other_service:
      image: fdsafdsa
    """
    And a file named "lc.yml" with:
    """yaml
    name: testbootstrap
    docker_image_name: baz
    """
    When I run `lc bootstrap`
    Then it should fail pulling "fdsafdsa"
    And it should not fail pulling "baz"

  Scenario: running in offline mode
    If we run with --offline, we should not try to pull any images.
    Given a file named "docker-compose.yml" with:
    """yaml
    foo:
      image: fdsafdsa
    """
    And a file named "lc.yml" with:
    """yaml
    name: testbootstrap
    docker_image_name: baz
    """
    When I run `lc bootstrap`
    Then it should fail pulling "fdsafdsa"
    When I run `lc --offline bootstrap`
    Then it should succeed
    And the output should not contain "Pulling repository docker.io/library/baz"

  Scenario: running with docker-compose.yml v2 file
    At a minimum, lc should not fail to parse a docker-compose.yml v2 file format
    if there is no lc-template being used.
    Given a file named "dev-env/Dockerfile" with:
    """
    FROM busybox
    """
    And a file named "docker-compose.yml" with:
    """yaml
    version: '2'

    services:
      prodserver:
        build: dev-env
    """
    And a file named "lc.yml" with:
    """yaml
    name: testbootstrap
    """
    When I run `lc bootstrap`
    Then it should succeed with "Building prodserver"
