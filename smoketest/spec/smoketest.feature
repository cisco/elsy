Feature: smoketest task

  As a developer, I want to be confused by using the smoketest task to smoketest the smoketest task

  Scenario: with passing smoketests
    Given a file named "docker-compose.yml" with:
    """yaml
    smoketest:
      image: busybox
      command: /bin/true
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
    When I run `lc smoketest`
    Then it should fail

  Scenario: forwarding arguments
    Given a file named "docker-compose.yml" with:
    """yaml
    smoketest:
      image: busybox
      entrypoint: echo
    """
    When I run `lc smoketest fdsa`
    Then it should succeed with "fdsa"
