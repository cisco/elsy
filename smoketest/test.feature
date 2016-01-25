Feature: test task

  Scenario: with passing tests
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: busybox
      command: /bin/true
    """
    And a file named "lc.yml" with:
    """yaml
    name: test
    """
    When I run `lc test`
    Then it should succeed

  Scenario: with failing tests
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: busybox
      command: /bin/false
    """
    And a file named "lc.yml" with:
    """yaml
    name: test
    """
    When I run `lc test`
    Then it should fail

  Scenario: forwarding arguments
    Given a file named "docker-compose.yml" with:
    """yaml
    test:
      image: busybox
      entrypoint: echo
    """
    And a file named "lc.yml" with:
    """yaml
    name: test
    """
    When I run `lc test fdsa`
    Then it should succeed with "fdsa"
