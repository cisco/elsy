Feature: passing through to docker-compose

  Scenario: with a standard project
    Given a file named "docker-compose.yml" with:
    """yaml
    foo:
      image: busybox
      command: echo bar
    """
    When I run `lc dc -- run --rm foo`
    Then it should succeed with "bar"
