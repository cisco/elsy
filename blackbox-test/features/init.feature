# Copyright 2016 Cisco Systems, Inc.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
# http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

Feature: init command

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
    When I run `lc init --template=mvn --project-name=mvnservice --docker-image-name=mvnserviceimage --docker-registry=internal-registry.something.com:5000 fullinit`
    Then it should succeed
    And the output should contain all of these:
      | Initializing lc project         |
      | using project-name: mvnservice  |
    And the file "fullinit/lc.yml" should contain the following:
      | project_name: mvnservice                              |
      | template: mvn                                         |
      | docker_image_name: mvnserviceimage                    |
      | docker_registry: internal-registry.something.com:5000 |
    And the file "fullinit/docker-compose.yml" should contain the following:
      | noop            |
      | image: alpine   |
    And the file "fullinit/Dockerfile" should contain the following:
      | FROM scratch    |

  Scenario: initializing a repo with all options and multiple registries
    When I run `lc init --template=mvn --project-name=mvnservice --docker-image-name=mvnserviceimage --docker-registry=internal-registry.something.com:5000 --docker-registry=internal-registry2.somethingelse.com fullinit`
    Then it should succeed
    And the output should contain all of these:
      | Initializing lc project         |
      | using project-name: mvnservice  |
    And the file "fullinit/lc.yml" should contain the following:
      | project_name: mvnservice                            |
      | template: mvn                                       |
      | docker_image_name: mvnserviceimage                  |
      | docker_registries: ["internal-registry.something.com:5000","internal-registry2.somethingelse.com",] |
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
