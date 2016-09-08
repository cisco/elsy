Feature: system list-templates task

  Scenario: calling with no args
    When I run `lc system list-templates`
    Then it should succeed
    And the output should contain all of these:
      | Run `lc system view-template <template-name>` to see the template contents. |
      | mvn |
      | sbt |
