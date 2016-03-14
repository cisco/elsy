Feature: init command

- test new directory
- test existing directory
- test three logical branches: https://lancope.slack.com/archives/dev-infrastructure/p1457726449000640


  Scenario: initializing the current directory
    When I run `mkdir inittest && cd inittest && lc init`
    Then it should succeed
    And the output should contain all of these:
      | Initializing lc project       |
      | using project-name: inittest  |
    And the file "inittest/lc.yml" should contain the following:
      | project_name: inittest  |
    And the file "inittest/lc.yml" should not contain the following:
      | template  |
      | docker    |

  Scenario: initializing a new directory
    When I run `lc init inittest-new`
    Then it should succeed
    And the output should contain all of these:
      | Initializing lc project       |
      | using project-name: inittestnew  |
    And the file "inittest-new/lc.yml" should contain the following:
      | project_name: inittestnew  |
    And the file "inittest-new/lc.yml" should not contain the following:
      | template  |
      | docker    |

  Scenario: initializing a pre-existing lc repo
    Given a file named "lc.yml" with:
      """yaml
      name: testpackage
      """
    When I run `lc init`
    Then it should fail with "repo already initialized"

  Scenario: initializing a repo with all options
    When I run `lc init --template=mvn --project-name=mvnservice --docker-image-name=mvnserviceimage --docker-registry=arch-docker.eng.lancope.local:5000 fullinit`
    Then it should succeed
    And the output should contain all of these:
      | Initializing lc project         |
      | using project-name: mvnservice  |
    And the file "fullinit/lc.yml" should contain the following:
      | project_name: mvnservice                            |
      | template: mvn                                       |
      | docker_image_name: mvnserviceimage                  |
    And the file "fullinit/docker-compose.yml" should contain the following:
      | noop            |
      | image: alpine   |
    And the file "fullinit/Dockerfile" should contain the following:
      | FROM scratch    |

  Scenario: initializing a repo should work
    When I run `lc init fullinittestrun`
    Then it should succeed
    And the output should contain all of these:
      | Initializing lc project         |
      | using project-name: fullinittestrun  |
    And the file "fullinittestrun/lc.yml" should contain the following:
      | project_name: fullinittestrun                       |
    When I run `cd fullinittestrun && lc bootstrap`
    Then it should succeed
    When I run `cd fullinittestrun && lc package`
    Then it should succeed
