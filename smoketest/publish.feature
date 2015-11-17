Feature: publish task

  Scenario: with no publish service
    Given a file named "docker-compose.yml" with:
    """
    """
    When I run `lc publish`
    Then it should succeed

  Scenario: with a publish service
    Given a file named "docker-compose.yml" with:
    """yaml
    publish:
      image: busybox
      command: echo foo
    """
    When I run `lc publish`
    Then it should succeed with "foo"

  Scenario: with a Docker project
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
    docker_registry: terrapin-registry0.eng.lancope.local:5000
    """
    When I run `lc package`
    And I run `lc publish --git-branch=origin/master`
    Then it should succeed with "Pushing repository terrapin-registry0.eng.lancope.local:5000/projectlifecyclesmoketests_docker_artifact"
