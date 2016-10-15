Feature: templates

  Scenario: with an invalid template
    Given a file named "lc.yml" with:
    """yaml
    name: testsmokes
    """
    When I run `lc --template foo bootstrap`
    Then it should fail
    And the output should contain 'template \\\"foo\\\" is not a registered v1 template'
