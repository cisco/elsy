Feature: smoketest task

  As a developer, I want to be confused by using the smoketest task to smoketest the smoketest task

  Scenario: with passing smoketests
    Given a file named "docker-compose.yml" with:
    """yaml
    smoketest:
      image: busybox
      command: /bin/true
    """
    And a file named "lc.yml" with:
    """yaml
    name: testsmokes
    """
    When I run `lc smoketest`
    Then it should succeed

  Scenario: with failing smoketests
    Given a file named "docker-compose.yml" with:
    """yaml
    smoketest:
      image: busybox
      command: /bin/false
    """
    And a file named "lc.yml" with:
    """yaml
    name: testsmokes
    """
    When I run `lc smoketest`
    Then it should fail

  Scenario: forwarding arguments
    Given a file named "docker-compose.yml" with:
    """yaml
    smoketest:
      image: busybox
      entrypoint: echo
    """
    And a file named "lc.yml" with:
    """yaml
    name: testsmokes
    """
    When I run `lc smoketest fdsa`
    Then it should succeed with "fdsa"

  ## see US7549
  Scenario: with defaulting to first running lc package
    Given I run `docker rmi -f projectlifecyclesmoketests_docker_artifact_smoketest`
    And a file named "docker-compose.yml" with:
    """yaml
    smoketest:
      image: projectlifecyclesmoketests_docker_artifact_smoketest
      command: /bin/true
    """
    And a file named "Dockerfile" with:
    """
    FROM library/alpine
    """
    And a file named "lc.yml" with:
    """yaml
    docker_image_name: projectlifecyclesmoketests_docker_artifact_smoketest
    """
    When I run `lc smoketest --skip-package`
    Then it should fail with 'image library/projectlifecyclesmoketests_docker_artifact_smoketest:latest not found'
    When I run `lc smoketest`
    Then it should succeed
    And the output should contain all of these:
      | Running package before executing smoketests     |
      | Successfully built                              |
