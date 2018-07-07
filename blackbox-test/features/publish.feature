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

Feature: publish task

  ## To run this test in a way that does not require an engineer or build-server to
  ## edit their 'docker daemon' '--insecure-registry' flag, we are relying on the
  ## fact that 'localhost' is automatically trusted. To that end we are making the
  ## following assumptions in this test:
  ##
  ## - The docker-compose.yml maps 'registry1:5000' maps to 'docker-daemon-host:5000'
  ## - The docker-compose.yml maps 'registry2:5000' maps to 'docker-daemon-host:5001'
  ## - From the perspecitive of the docker-daemon host:
  ##    - registry1 is available at http://localhost:5000
  ##    - registry1 is available at http://localhost:5001
  Background:
    Given registry1 is listening on port 5000
    And registry2 is listening on port 5000

  Scenario: with an empty project, calling publish without the git branch
    When I run `lc publish`
    Then it should fail with "expecting a git branch and/or a git tag be set, found neither"

  Scenario: with both docker_registry and docker_registries defined
    Given a file named "lc.yml" with:
    """yaml
    docker_registry: localhost:5000
    docker_registries:
      - localhost:5000
    """
    When I run `lc bootstrap`
    Then it should fail with "multiple docker registry configs found, pick either"
    When I run `lc test`
    Then it should fail with "multiple docker registry configs found, pick either"
    When I run `lc package`
    Then it should fail with "multiple docker registry configs found, pick either"
    When I run `lc publish`
    Then it should fail with "multiple docker registry configs found, pick either"


  Scenario: with a publish service, calling publish on the master branch
    Given a file named "docker-compose.yml" with:
    """yaml
    publish:
      image: busybox
      environment:
        - LC_PUBLISH_DOCKER_TAG
      command: echo custom publish of tag $LC_PUBLISH_DOCKER_TAG
    """
    When I run `lc publish --git-branch origin/master`
    Then it should succeed with "custom publish of tag latest"

  Scenario: with a publish service, calling publish on a release branch
    Given a file named "docker-compose.yml" with:
    """yaml
    publish:
      image: busybox
      environment:
        - LC_PUBLISH_DOCKER_TAG
      command: echo custom publish of tag $LC_PUBLISH_DOCKER_TAG
    """
    When I run `lc publish --git-branch origin/release/1.5`
    Then it should succeed with "custom publish of tag 1.5"

  Scenario: with a publish service, calling publish on a feature branch
    Given a file named "docker-compose.yml" with:
    """yaml
    publish:
      image: busybox
      environment:
        - LC_PUBLISH_DOCKER_TAG
      command: echo custom publish of tag $LC_PUBLISH_DOCKER_TAG
    """
    When I run `lc publish --git-branch origin/fix-thing`
    Then it should succeed with "skipping custom publish task"

  Scenario: with a Docker project, calling publish on the master branch
    Given a file named "docker-compose.yml" with:
    """yaml
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
    docker_image_name: elsyblackbox_docker_artifact1
    docker_registry: localhost:5000
    """
    When I run `lc package`
    And I run `lc publish --git-branch=origin/master`
    And the output should contain all of these:
      | The push refers to repository [localhost:5000/elsyblackbox_docker_artifact1]   |
      | latest: digest: sha256                                                                       |

  Scenario: with a Docker project, calling publish on the master branch with multiple registries
    Given a file named "docker-compose.yml" with:
    """yaml
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
    docker_image_name: elsyblackbox_docker_artifact2
    docker_registries:
      - localhost:5000
      - localhost:5001
    """
    When I run `lc package`
    And I run `lc publish --git-branch=origin/master`
    And the output should contain all of these:
      | The push refers to repository [localhost:5000/elsyblackbox_docker_artifact2]   |
      | The push refers to repository [localhost:5001/elsyblackbox_docker_artifact2]   |
      | latest: digest: sha256                                                                       |

  Scenario: with a Docker project, calling publish on a feature branch
    Given a file named "docker-compose.yml" with:
    """yaml
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
    docker_image_name: elsyblackbox_docker_artifact3
    docker_registry: localhost:5000
    """
    When I run `lc package`
    And I run `lc publish --git-branch=origin/feature/thing`
    Then it should succeed
    And the output should contain all of these:
      | The push refers to repository [localhost:5000/elsyblackbox_docker_artifact3]|
      | snapshot.feature.thing: digest: sha256                                                    |

  Scenario: with a publish service, calling publish on a release tag
    Given a file named "docker-compose.yml" with:
    """yaml
    publish:
      image: busybox
      environment:
        - LC_PUBLISH_DOCKER_TAG
      command: echo custom publish of tag $LC_PUBLISH_DOCKER_TAG
    """
    When I run `lc publish --git-tag=v0.2.3`
    Then it should succeed with "custom publish of tag v0.2.3"

  Scenario: with a publish service, calling publish on a non-release tag
    Given a file named "docker-compose.yml" with:
    """yaml
    publish:
      image: busybox
      environment:
        - LC_PUBLISH_DOCKER_TAG
      command: echo custom publish of tag $LC_PUBLISH_DOCKER_TAG
    """
    When I run `lc publish --git-tag=foo-test`
    Then it should succeed with "custom publish of tag snapshot.foo-test"

  Scenario: with a publish service, calling publish on a bogus tag
    Given a file named "docker-compose.yml" with:
    """yaml
    publish:
      image: busybox
      command: echo foo
    """
    When I run `lc publish --git-tag=x/[z`
    Then it should fail
    And the output should contain all of these:
      | snapshot.x.[z |
      | not valid |

  Scenario: with a Docker project, calling publish on a new release tag
    Given a file named "docker-compose.yml" with:
    """yaml
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
    docker_image_name: elsyblackbox_docker_artifact4
    docker_registry: localhost:5000
    """
    When I run `lc package`
    And I run `lc publish --git-tag=v0.2.2 --git-branch=origin/master`
    Then it should succeed
    And the output should contain all of these:
      | The push refers to repository [localhost:5000/elsyblackbox_docker_artifact4]|
      | v0.2.2: digest: sha256                                                                    |

  Scenario: with a Docker project, calling publish on a non-release tag
    Given a file named "docker-compose.yml" with:
    """yaml
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
    docker_image_name: elsyblackbox_docker_artifact5
    docker_registry: localhost:5000
    """
    When I run `lc package`
    And I run `lc publish --git-tag=foo-test --git-branch=origin/master`
    Then it should succeed
    And the output should contain all of these:
      | The push refers to repository [localhost:5000/elsyblackbox_docker_artifact5|
      | snapshot.foo-test: digest: sha256                                                        |
