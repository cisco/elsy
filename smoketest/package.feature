Feature: package task

  Scenario: with a package service
    Given a file named "docker-compose.yml" with:
    """yaml
    package:
      image: busybox
      command: echo foo
    """
    When I run `lc package`
    Then it should succeed with "foo"

  Scenario: with a failing package service
    Given a file named "docker-compose.yml" with:
    """yaml
    package:
      image: busybox
      command: /bin/false
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
    FROM alpine
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
    FROM alpine
    """
    And a file named "lc.yml" with:
    """yaml
    docker_image_name: projectlifecyclesmoketests_docker_artifact
    """
    When I run `lc package`
    Then it should succeed with "Image is up to date for alpine:latest"
