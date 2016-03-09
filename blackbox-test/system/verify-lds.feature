Feature: system verify-lds task

## We can't really test the disk mounting repeatedly, so just verify that the file check works
## and assume the disk mounting is always working (it will be in ci since it is all linux)

  Scenario: calling with missing lc.yml file
    When I run `lc system verify-lds`
    Then it should succeed
    And the output should contain "could not find 'lc.yml' in the current directory, skipping lds verification"

  Scenario: calling with lc.yml file
    Given a file named "lc.yml" with:
    """yaml
    name: test-verify-lds
    """
    When I run `lc system verify-lds`
    Then it should succeed
    And the output should not contain "could not find 'lc.yml' in the current directory, skipping lds verification"
