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
    docker_image_name: projectlifecycleblackbox_docker_artifact_ci
    docker_registry: localhost:5000
    """
    And I run `lc ci --git-branch=origin/master`
    Then it should succeed
    And the output should contain all of these:
      | The push refers to a repository [localhost:5000/projectlifecycleblackbox_docker_artifact_ci]|
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
    docker_image_name: projectlifecycleblackbox_docker_artifact_ci_labels
    docker_registry: localhost:5000
    """
    And I run `lc ci --git-commit=d8dfd9f --git-branch=origin/master`
    Then it should succeed
    And the output should contain all of these:
      | The push refers to a repository [localhost:5000/projectlifecycleblackbox_docker_artifact_ci_labels]|
      | latest: digest: sha256                                                                             |
      | Attaching image label: com.elsy.metadata.git-commit=d8dfd9f                                        |
    And the image 'projectlifecycleblackbox_docker_artifact_ci_labels' should exist
    And it should have the following labels:
      | com.elsy.metadata.git-commit:d8dfd9f              |
