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

Feature: ci task

  ## See ./blackbox-test/publish.feature for a description of how the registry
  ## is setup for this test.
  Background:
    Given registry1 is listening on port 5000

  Scenario: with a package service and a test service, should not run tests
    Given a file named "docker-compose.yml" with:
    """yaml
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
    When I run `lc ci`
    Then it should fail
    And the output should not contain "Running tests before packaging"

  Scenario: with a failing test
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: busybox
      command: /bin/false
    """
    And a file named "lc.yml" with:
    """yaml
    name: testci
    """
    When I run `lc ci`
    Then it should fail

  Scenario: with a Docker project
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: busybox
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
    docker_image_name: elsyblackbox_docker_artifact_ci
    docker_registry: localhost:5000
    """
    And I run `lc ci --git-branch=origin/master`
    Then it should succeed
    And the output should contain all of these:
      | The push refers to a repository [localhost:5000/elsyblackbox_docker_artifact_ci]|
      | latest: digest: sha256                                                                      |

  ## Only works in docker 1.11 and higher
  Scenario: with a Docker project and image labels
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: busybox
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
    docker_image_name: elsyblackbox_docker_artifact_ci_labels
    docker_registry: localhost:5000
    """
    And I run `lc ci --git-commit=d8dfd9f --git-branch=origin/master`
    Then it should succeed
    And the output should contain all of these:
      | The push refers to a repository [localhost:5000/elsyblackbox_docker_artifact_ci_labels]|
      | latest: digest: sha256                                                                             |
      | Attaching image label: com.elsy.metadata.git-commit=d8dfd9f                                        |
    And the image 'elsyblackbox_docker_artifact_ci_labels' should exist
    And it should have the following labels:
      | com.elsy.metadata.git-commit:d8dfd9f              |

  Scenario: with no publish service and no Dockerfile
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: busybox
      command: /bin/true
    package:
      image: busybox
      command: /bin/true
    """
    And a file named "lc.yml" with:
    """yaml
    docker_image_name: elsyblackbox_docker_artifact_ci
    docker_registry: localhost:5000
    """
    And I run `lc ci --git-branch=origin/master`
    Then it should succeed
    And the output should contain "No publish service defined, and no Dockerfile present."

  Scenario: Printing build logs
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: busybox
      command: /bin/true
      links:
        - stderrtest
        - stdouttest
    stderrtest:
      image: busybox
      command: /bin/sh -c "echo test stderr capture >&2 && tail -f /dev/null"
    stdouttest:
      image: busybox
      command: /bin/sh -c "echo test stdout capture >&2 && tail -f /dev/null"
    """
    And a file named "lc.yml" with:
    """yaml
    docker_image_name: elsyblackbox_docker_artifact_ci
    docker_registry: localhost:5000
    build_logs_dir: test_build_logs
    """
    And I run `lc ci --git-branch=origin/master`
    Then it should succeed
    And the output should contain "writing build logs to build-logs-dir"
    And the file "test_build_logs/test" should contain the following:
      | Attaching to         |
    And the file "test_build_logs/stderrtest" should contain the following:
      | test stderr capture  |
    And the file "test_build_logs/stdouttest" should contain the following:
      | test stdout capture  |

  Scenario: Run clean during ci
    Given a file named "docker-compose.yml" with:
    """yaml
    package:
      image: busybox
      command: echo foo
    test:
      image: busybox
      command: /bin/true
    clean:
      image: busybox
      volumes:
        - .:/test
      command: ["rm", "/test/foo.txt"]
    """
    And a file named "lc.yml" with:
    """yaml
    name: testpackage
    """
    And a file named "foo.txt" with:
    """
    garbage file
    """
    When I run `lc ci`
    Then it should succeed
    And the file "/test/foo.txt" should not exist

  Scenario: Run clean during ci with scratch volumes enabled
    Given a file named "docker-compose.yml" with:
    """yaml
    package:
      image: busybox
      command: echo foo
    test:
      image: busybox
      command: /bin/true
    clean:
      image: busybox
      volumes:
        - .:/test
      command: ["rm", "/test/foo.txt"]
    """
    And a file named "lc.yml" with:
    """yaml
    name: testpackage
    """
    And a file named "foo.txt" with:
    """
    garbage file
    """
    When I run `lc --enable-scratch-volumes ci`
    Then it should succeed
    And the file "/test/foo.txt" should not exist
