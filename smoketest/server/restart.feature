Feature: server restart task

  @teardown
  Scenario: with no server running
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server restart`
    Then it should succeed with 'starting service \"devserver\"'

  @teardown
  Scenario: with a dev server running
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server start`
    When I run `lc server restart`
    Then it should succeed
    And the output should contain all of these:
      | Stopping server                |
      | starting service \"devserver\" |

  @teardown
  Scenario: with a prod server running
    Given a file named "docker-compose.yml" with:
    """yaml
    prodserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server start -prod`
    When I run `lc server restart`
    Then it should succeed
    And the output should contain all of these:
      | Stopping server                |
      | starting service \"prodserver\" |
