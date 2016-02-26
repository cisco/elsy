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
    When I run `lc package --docker-image-name=projectlifecyclesmoketests_docker_artifact`
    Then it should succeed with "Image is up to date for alpine:latest"
    When I run `lc package --docker-image-name=projectlifecyclesmoketests_docker_artifact --skip-docker`
    Then it should succeed
    And the output should not contain "Successfully built "

  Scenario: with a docker artifact based on a local image
    Given a file named "Dockerfile" with:
    """
    FROM projectlifecyclesmoketests_docker_artifact
    """
    And a file named "lc.yml" with:
    """yaml
    name: testpackage
    """
    When I run `lc package --docker-image-name=projectlifecyclesmoketests_docker_artifact2`
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
    When I run `lc package --docker-image-name=projectlifecyclesmoketests_docker_artifact`
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
    docker_image_name: projectlifecyclesmoketests_docker_artifact
    """
    When I run `lc package`
    Then it should succeed with "Image is up to date for alpine:latest"
