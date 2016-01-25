Feature: command execution


Scenario: with failing command
  Given a file named "docker-compose.yml" with:
  """yaml
  test:
    image: busybox
    command: /bin/false
  """
  And a file named "lc.yml" with:
  """yaml
  name: testcommand
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
  And a file named "lc.yml" with:
  """yaml
  name: testcommand
  """
  When I run `lc test`
  Then it should succeed
  When I run `lc --debug test`
  Then it should succeed

Scenario: with no command
  When I run `lc`
  Then it should succeed
