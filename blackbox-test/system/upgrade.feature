Feature: upgrade the system

  Scenario: trying to upgrade lc should result in a message
    When I run `lc system upgrade`
    Then it should succeed
    And the output should contain "lc no longer updates itself. Run `lds upgrade`, to upgrade it."
