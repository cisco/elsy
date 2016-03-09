Feature: server start task

  @teardown
  Scenario: starting a devserver
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server start`
    Then it should succeed
    And it should report a correct address

  @teardown
  Scenario: starting a prodserver
    Given a file named "docker-compose.yml" with:
    """yaml
    prodserver:
      image: nginx
      ports:
       - "80"
    """
    When I run `lc server start -prod`
    Then it should succeed
    And it should report a correct address

  @teardown
  Scenario: starting server with no devserver service defined
    Given a file named "docker-compose.yml" with:
    """yaml
    """
    When I run `lc server start`
    Then it should fail with 'no \"devserver\" service defined'

  ## this is mainly testing that this error condition provides a sensible error to the user
  Scenario: starting server with no image should attempt to pull the image
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: somefakeimagethatdoesntexist
      ports:
       - "80"
    """
    When I run `lc server start`
    Then it should fail with 'image library/somefakeimagethatdoesntexist:latest not found'

  @teardown
  Scenario: starting prod server with no prodserver service defined
    Given a file named "docker-compose.yml" with:
    """yaml
    """
    When I run `lc server start -prod`
    Then it should fail with 'no \"prodserver\" service defined'

  @teardown
  Scenario: starting prod when devserver is already running
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: nginx
      ports:
       - "80"
    prodserver:
     image: nginx
     ports:
      - "80"
    """
    When I run `lc server start`
    And I run `lc server start -prod`
    Then it should succeed with '\"devserver\" already running'

  @teardown
  Scenario: starting dev when prodserver is already running
    Given a file named "docker-compose.yml" with:
    """yaml
    devserver:
      image: nginx
      ports:
       - "80"
    prodserver:
     image: nginx
     ports:
      - "80"
    """
    When I run `lc server start -prod`
    And I run `lc server start`
    Then it should succeed with '\"prodserver\" already running'
