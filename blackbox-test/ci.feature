Feature: ci task
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
    docker_image_name: projectlifecycleblackbox_docker_artifact
    docker_registry: terrapin-registry0.eng.lancope.local:5000
    """
    And I run `lc ci --git-branch=origin/master`
    Then it should succeed with "Pushing repository terrapin-registry0.eng.lancope.local:5000/projectlifecycleblackbox_docker_artifact"
