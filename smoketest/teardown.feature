Feature: teardown task

  As a developer, I want to teardown containers associated with my project

  Scenario: with no gc labels
    Given a file named "docker-compose.yml" with:
    """yaml
    teardowntestcontainer:
      image: busybox
      command: /bin/true
    """
    When I run `lc dc up teardowntestcontainer`
    Then it should succeed
    When I run `lc dc ps`
    Then the output should contain 'teardowntestcontainer'
    When I run `lc teardown`
    Then it should succeed
    And I run `lc dc ps`
    Then the output should not contain 'teardowntestcontainer'

  Scenario: with gc labels
    Given a file named "docker-compose.yml" with:
    """yaml
    teardowntestcontainerwithgc:
      image: busybox
      labels:
        com.lancope.docker-gc.keep: "True"
      command: /bin/true
    """
    When I run `lc dc up teardowntestcontainerwithgc`
    Then it should succeed
    When I run `lc dc ps`
    Then the output should contain 'teardowntestcontainerwithgc'
    When I run `lc teardown`
    Then it should succeed
    And I run `lc dc ps`
    Then the output should contain 'teardowntestcontainerwithgc'