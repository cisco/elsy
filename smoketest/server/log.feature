Feature: server log task

  Scenario: with no server services defined
    Given a file named "docker-compose.yml" with:
    """yaml
    """
    When I run `lc server log`
    Then it should fail with "No devserver or prodserver service defined"

  Scenario: with a devserver not running
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server log`
    Then it should fail with "Server is not running"

  Scenario: with a prodserver not running
    Given a file named "docker-compose.yml" with:
    """yaml
    prodserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server log`
    Then it should fail with "Server is not running"
