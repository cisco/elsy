Feature: server stop task

  @teardown
  Scenario: stopping a running devserver
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server start && lc server stop`
    Then it should succeed with "Stopping server"

  @teardown
  Scenario: stopping a running prodserver
    Given a file named "docker-compose.yml" with:
    """yaml
    prodserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server start -prod && lc server stop`
    Then it should succeed with "Stopping server"

  @teardown
  Scenario: stopping with no server services
    Given a file named "docker-compose.yml" with:
    """yaml
    """
    When I run `lc server stop`
    Then it should succeed
