Feature: package task

  Scenario: with a package service
    Given a file named "docker-compose.yml" with:
    """yaml
    package:
      image: busybox
      command: echo foo
    """
    And a file named "lc.yml" with:
    """yaml
    name: testpackage
    """
    When I run `lc package`
    Then it should succeed with "foo"

  Scenario: with a package service and a test service
    Given a file named "docker-compose.yml" with:
    """yaml
    package:
      image: busybox
      command: echo foo
    test:
      image: busybox
      command: echo running tests
    """
    And a file named "lc.yml" with:
    """yaml
    name: testpackage
    """
    When I run `lc package`
    Then it should succeed
    And the output should contain all of these:
      | Running tests before packaging |
      | running tests  |
      | foo |

  Scenario: with a package service and a test service, skipping tests
    Given a file named "docker-compose.yml" with:
    """yaml
    package:
      image: busybox
      command: echo foo
    test:
      image: busybox
      command: echo running tests
    """
    And a file named "lc.yml" with:
    """yaml
    name: testpackage
    """
    When I run `lc package -skip-tests`
    Then it should succeed with "foo"
    And the output should not contain any of these:
      | Running tests before packaging |
      | running tests |

  Scenario: with a package service and no test service
    Given a file named "docker-compose.yml" with:
    """yaml
    package:
      image: busybox
      command: echo foo
    """
    And a file named "lc.yml" with:
    """yaml
    name: testpackage
    """
    When I run `lc package`
    Then it should succeed with "foo"
    And the output should not contain any of these:
      | Running tests before packaging |
      | running tests  |

  Scenario: with a package service and no test service, skipping tests
    Given a file named "docker-compose.yml" with:
    """yaml
    package:
      image: busybox
      command: echo foo
    """
    And a file named "lc.yml" with:
    """yaml
    name: testpackage
    """
    When I run `lc package -skip-tests`
    Then it should succeed with "foo"
    And the output should not contain any of these:
      | Running tests before packaging |
      | running tests |

  Scenario: without a package service
    Given a file named "lc.yml" with:
    """yaml
    name: testpackage
    """
    When I run `lc package`
    Then it should succeed

  Scenario: with a docker artifact
    Given a file named "Dockerfile" with:
    """
    FROM library/alpine
    """
    And a file named "lc.yml" with:
    """yaml
    name: testpackage
    """
    When I run `lc package --docker-image-name=projectlifecycleblackbox_docker_artifact`
    Then it should succeed with "Image is up to date for alpine:latest"
    When I run `lc package --docker-image-name=projectlifecycleblackbox_docker_artifact --skip-docker`
    Then it should succeed
    And the output should not contain "Successfully built "

  Scenario: with a docker artifact based on a local image
    Given a file named "Dockerfile" with:
    """
    FROM projectlifecycleblackbox_docker_artifact
    """
    And a file named "lc.yml" with:
    """yaml
    name: testpackage
    """
    When I run `lc package --docker-image-name=projectlifecycleblackbox_docker_artifact`
    Then it should succeed with "Successfully built "

  Scenario: with a failing package service
    Given a file named "docker-compose.yml" with:
    """yaml
    package:
      image: busybox
      command: /bin/false
    """
    And a file named "lc.yml" with:
    """yaml
    name: testpackage
    """
    When I run `lc package`
    Then it should fail

  Scenario: with a docker artifact and insufficient args
    Given a file named "docker-compose.yml" with:
    """yaml
    package:
      image: busybox
      command: /bin/true
    """
    And a file named "Dockerfile" with:
    """
    FROM busybox
    """
    And a file named "lc.yml" with:
    """yaml
    name: testpackage
    """
    When I run `lc package`
    Then it should fail with "you must use `--docker-image-name` to package a docker image"

  Scenario: with a docker artifact and configured image name
    Given a file named "docker-compose.yml" with:
    """yaml
    package:
      image: busybox
      command: /bin/true
    """
    And a file named "Dockerfile" with:
    """
    FROM library/alpine
    """
    And a file named "lc.yml" with:
    """yaml
    name: testpackage
    """
    When I run `lc package --docker-image-name=projectlifecycleblackbox_docker_artifact`
    Then it should succeed with "Image is up to date for alpine:latest"

  Scenario: with a docker artifact and configured image name via lc.yml
    Given a file named "docker-compose.yml" with:
    """yaml
    package:
      image: busybox
      command: /bin/true
    """
    And a file named "Dockerfile" with:
    """
    FROM library/alpine
    """
    And a file named "lc.yml" with:
    """yaml
    docker_image_name: projectlifecycleblackbox_docker_artifact
    """
    When I run `lc package`
    Then it should succeed with "Image is up to date for alpine:latest"

  Scenario: custom package script generating the Dockerfile
    Given a file named "docker-compose.yml" with:
    """yaml
    package:
      image: busybox
      volumes:
        - .:/opt/project
      command: ["sh", "-c", "echo 'FROM library/alpine' > /opt/project/Dockerfile"]
    """
    And a file named "lc.yml" with:
    """yaml
    docker_image_name: projectlifecycleblackbox_docker_artifact
    """
    When I run `lc package`
    Then it should succeed with "Image is up to date for alpine:latest"
