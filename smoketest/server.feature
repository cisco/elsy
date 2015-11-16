Feature: server task

  Scenario: running a devserver
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

  Scenario: stopping a running devserver
  Given a file named "docker-compose.yml" with:
  """yaml
  devserver:
    image: nginx
    ports:
     - "80"
  """
  When I run `lc server start && lc server stop`
  Then it should succeed with "Stopping devserver"

  Scenario: trying to get status when server not running
  Given a file named "docker-compose.yml" with:
  """yaml
  devserver:
    image: nginx
    ports:
     - "80"
  """
  When I run `lc server stat`
  Then it should succeed with "devserver: down"
