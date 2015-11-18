Feature: system view-template task

  Scenario: calling with no args
    When I run `lc system view-template`
    Then it should fail with "view-template requires an argument"

  Scenario: calling with an invalid template
    When I run `lc system view-template foo`
    Then it should fail with 'template \"foo\" is not registered'

  Scenario: calling with a valid template
    When I run `lc system view-template mvn`
    Then it should succeed
    And the output should contain all of these:
      | mvn:     |
      | test:    |
      | package: |
