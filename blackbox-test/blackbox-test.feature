Feature: blackbox-test task

  As a developer, I want to be confused by using the blackbox-test task to blackbox test the blackbox-test task

  Scenario: with passing blackbox-test
    Given a file named "docker-compose.yml" with:
    """yaml
    blackbox-test:
      image: busybox
      command: /bin/true
    """
    And a file named "lc.yml" with:
    """yaml
    name: testsmokes
    """
    When I run `lc blackbox-test`
    Then it should succeed

  Scenario: with failing blackbox-test
    Given a file named "docker-compose.yml" with:
    """yaml
    blackbox-test:
      image: busybox
      command: /bin/false
    """
    And a file named "lc.yml" with:
    """yaml
    name: testsmokes
    """
    When I run `lc blackbox-test`
    Then it should fail

  Scenario: forwarding arguments
    Given a file named "docker-compose.yml" with:
    """yaml
    blackbox-test:
      image: busybox
      entrypoint: echo
    """
    And a file named "lc.yml" with:
    """yaml
    name: testsmokes
    """
    When I run `lc blackbox-test fdsa`
    Then it should succeed with "fdsa"

  ## see US7549
  Scenario: with defaulting to first running lc package
    Given I run `docker rmi -f projectlifecycleblackbox_docker_artifact_blackbox`
    And a file named "docker-compose.yml" with:
    """yaml
    blackbox-test:
      image: projectlifecycleblackbox_docker_artifact_blackbox
      command: /bin/true
    """
    And a file named "Dockerfile" with:
    """
    FROM library/alpine
    """
    And a file named "lc.yml" with:
    """yaml
    docker_image_name: projectlifecycleblackbox_docker_artifact_blackbox
    """
    When I run `lc blackbox-test --skip-package`
    Then it should fail with 'image library/projectlifecycleblackbox_docker_artifact_blackbox:latest not found'
    When I run `lc blackbox-test`
    Then it should succeed
    And the output should contain all of these:
      | Running package before executing blackbox tests |
      | Successfully built                              |
