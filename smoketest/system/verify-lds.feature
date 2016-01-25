Feature: system verify-lds task

## We can't really test the disk mounting repeatedly, so just verify that the file check works
## and assume the disk mounting is always working (it will be in ci since it is all linux)

  Scenario: calling with missing lc.yml file
    When I run `lc system verify-lds`
    Then it should fail with "It appears that your local disk is not mounted into the LDS VM"

  Scenario: calling with missing lc.yml file
    Given a file named "lc.yml" with:
    """yaml
    name: test-verify-lds
    """
    When I run `lc system verify-lds`
    Then it should succeed
