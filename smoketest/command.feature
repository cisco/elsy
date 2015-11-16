Feature: command execution


Scenario: with failing command
  Given a file named "docker-compose.yml" with:
  """yaml
  test:
    image: busybox
    command: /bin/false
  """
  When I run `lc test`
  Then it should fail
  When I run `lc --debug test`
  Then it should fail

Scenario: with successful command
  Given a file named "docker-compose.yml" with:
  """yaml
  test:
    image: busybox
    command: /bin/true
  """
  When I run `lc test`
  Then it should succeed
  When I run `lc --debug test`
  Then it should succeed