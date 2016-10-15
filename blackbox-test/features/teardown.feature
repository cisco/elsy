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
        com.lancope.docker-gc.keep: "value-does-not-matter"
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

  Scenario: with gc labels and -f flag
    Given a file named "docker-compose.yml" with:
    """yaml
    teardowntestcontainer:
      image: busybox
      labels:
        com.lancope.docker-gc.keep: "True"
      command: /bin/true
    """
    When I run `lc dc up teardowntestcontainer`
    Then it should succeed
    When I run `lc dc ps`
    Then the output should contain 'teardowntestcontainer'
    When I run `lc teardown -f`
    Then it should succeed
    And I run `lc dc ps`
    Then the output should not contain 'teardowntestcontainer'

  Scenario: with v2 networks and -f flag
    Given a file named "docker-compose.yml" with:
    """yaml
    version: '2'

    services:
      test:
        image: busybox
        command: "/bin/true"
        networks:
          - elsy_teardown_lc_bbt_network_test

    networks:
      elsy_teardown_lc_bbt_network_test:
    """
    When I run `lc test`
    Then it should succeed
    When I run `docker network ls`
    Then the output should contain 'elsy_teardown_lc_bbt_network_test'
    When I run `lc teardown -f`
    Then it should succeed
    When I run `docker network ls`
    Then the output should not contain 'elsy_teardown_lc_bbt_network_test'

  Scenario: with v2 volumes
    Given a file named "docker-compose.yml" with:
    """yaml
    version: '2'

    volumes:
      teardown_test_build_cache:

    services:
      test:
        image: busybox
        volumes:
          - teardown_test_build_cache:/opt/cache
        command: "touch /opt/cache/foo"
    """
    When I run `lc test`
    Then it should succeed
    When I run `docker volume ls -q`
    Then the output should contain 'teardown_test_build_cache'
    When I run `lc teardown -f`
    Then it should succeed
    And I run `docker volume ls -q`
    Then the output should not contain 'teardown_test_build_cache'
