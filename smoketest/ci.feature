Feature: ci task

  Scenario: with a failing test
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: busybox
      command: /bin/false
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
    docker_image_name: projectlifecyclesmoketests_docker_artifact
    docker_registry: terrapin-registry0.eng.lancope.local:5000
    """
    And I run `lc ci --git-branch=origin/master`
    Then it should succeed with "Pushing repository terrapin-registry0.eng.lancope.local:5000/projectlifecyclesmoketests_docker_artifact"
