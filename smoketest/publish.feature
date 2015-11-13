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
