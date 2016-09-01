Feature: server status task

  @teardown
  Scenario: with no server services defined
    Given a file named "docker-compose.yml" with:
    """yaml
    """
    When I run `lc server status`
    Then it should fail with "No devserver or prodserver service defined"

  @teardown
  Scenario: with a devserver not running
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server status`
    Then it should fail with "Server is not running"

  @teardown
  Scenario: with a prodserver not running
    Given a file named "docker-compose.yml" with:
    """yaml
    prodserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server status`
    Then it should fail with "Server is not running"

  @teardown
  Scenario: with a devserver running
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server start`
    And I run `lc server status`
    Then it should succeed
    And the output should contain all of these:
     | Server is running using \"devserver\" |
     | 80/tcp is available at 127.0.0.1:     |

  @teardown
  Scenario: with a prodserver running
    Given a file named "docker-compose.yml" with:
    """yaml
    prodserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server start -prod`
    And I run `lc server status`
    Then it should succeed
    And the output should contain all of these:
     | Server is running using \"prodserver\" |
     | 80/tcp is available at 127.0.0.1:      |

   @teardown
   Scenario: with multiple ports
     Given a file named "docker-compose.yml" with:
     """yaml
     devserver:
       image: nginx
       ports:
        - "8080"
        - "8081"
     """
     When I run `lc server start`
     And I run `lc server status`
     Then it should succeed
     And the output should contain all of these:
      | Server is running using \"devserver\" |
      | 8080/tcp is available at 127.0.0.1:     |
      | 8081/tcp is available at 127.0.0.1:     |
