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

Feature: templates

  Scenario: with an invalid template
    Given a file named "lc.yml" with:
    """yaml
    name: testsmokes
    """
    When I run `lc --template foo bootstrap`
    Then it should fail
    And the output should contain 'template \\\"foo\\\" is not a registered v1 template'

  Scenario: using default maven template image
    Given a file named "lc.yml" with:
    """yaml
    name: testing
    template: mvn
    """
    When I run `lc dc config`
    Then it should succeed
    And the output should contain 'image: maven:3.2-jdk-8'

  Scenario: using overridden maven template image
    Given a file named "lc.yml" with:
    """yaml
    name: testing
    template: mvn
    template_image: example/maven
    """
    When I run `lc dc config`
    Then it should succeed
    And the output should contain 'image: example/maven'
    And the output should not contain 'image: maven:3.2-jdk-8'

  Scenario: using default sbt template image
    Given a file named "lc.yml" with:
    """yaml
    name: testing
    template: sbt
    """
    When I run `lc dc config`
    Then it should succeed
    And the output should contain 'image: paulcichonski/sbt'

  Scenario: using overridden sbt template image
    Given a file named "lc.yml" with:
    """yaml
    name: testing
    template: sbt
    template_image: example/sbt
    """
    When I run `lc dc config`
    Then it should succeed
    And the output should contain 'image: example/sbt'
    And the output should not contain 'image: paulcichonski/sbt'

  Scenario: using default lein template image
    Given a file named "lc.yml" with:
    """yaml
    name: testing
    template: lein
    """
    When I run `lc dc config`
    Then it should succeed
    And the output should contain 'image: clojure:lein-2.6.1'

  Scenario: using overridden lein template image
    Given a file named "lc.yml" with:
    """yaml
    name: testing
    template: lein
    template_image: example/lein
    """
    When I run `lc dc config`
    Then it should succeed
    And the output should contain 'image: example/lein'
    And the output should not contain 'image: clojure:lein-2.6.1'

  Scenario: using default make template image
    Given a file named "lc.yml" with:
    """yaml
    name: testing
    template: make
    """
    When I run `lc dc config`
    Then it should succeed
    And the output should contain 'image: gcc:6.1'

  Scenario: using overridden make template image
    Given a file named "lc.yml" with:
    """yaml
    name: testing
    template: make
    template_image: example/make
    """
    When I run `lc dc config`
    Then it should succeed
    And the output should contain 'image: example/make'
    And the output should not contain 'image: gcc:6.1'

  # V2

  Scenario: using default maven template image with a v2 compose file
    Given a file named "lc.yml" with:
    """yaml
    name: testing
    template: mvn
    """
    And a file named "docker-compose.yml" with:
    """yaml
      version: '2'
      services:
        no-op:
          image: busybox
    """
    When I run `lc dc config`
    Then it should succeed
    And the output should contain 'image: maven:3.2-jdk-8'

  Scenario: using overridden maven template image with a v2 compose file
    Given a file named "lc.yml" with:
    """yaml
    name: testing
    template: mvn
    template_image: example/maven
    """
    And a file named "docker-compose.yml" with:
    """yaml
      version: '2'
      services:
        no-op:
          image: busybox
    """
    When I run `lc dc config`
    Then it should succeed
    And the output should contain 'image: example/maven'
    And the output should not contain 'image: maven:3.2-jdk-8'

  Scenario: using default sbt template image with a v2 compose file
    Given a file named "lc.yml" with:
    """yaml
    name: testing
    template: sbt
    """
    And a file named "docker-compose.yml" with:
    """yaml
      version: '2'
      services:
        no-op:
          image: busybox
    """
    When I run `lc dc config`
    Then it should succeed
    And the output should contain 'image: paulcichonski/sbt'

  Scenario: using overridden sbt template image with a v2 compose file
    Given a file named "lc.yml" with:
    """yaml
    name: testing
    template: sbt
    template_image: example/sbt
    """
    And a file named "docker-compose.yml" with:
    """yaml
      version: '2'
      services:
        no-op:
          image: busybox
    """
    When I run `lc dc config`
    Then it should succeed
    And the output should contain 'image: example/sbt'
    And the output should not contain 'image: paulcichonski/sbt'

  Scenario: using default lein template image with a v2 compose file
    Given a file named "lc.yml" with:
    """yaml
    name: testing
    template: lein
    """
    And a file named "docker-compose.yml" with:
    """yaml
      version: '2'
      services:
        no-op:
          image: busybox
    """
    When I run `lc dc config`
    Then it should succeed
    And the output should contain 'image: clojure:lein-2.6.1'

  Scenario: using overridden lein template image with a v2 compose file
    Given a file named "lc.yml" with:
    """yaml
    name: testing
    template: lein
    template_image: example/lein
    """
    And a file named "docker-compose.yml" with:
    """yaml
      version: '2'
      services:
        no-op:
          image: busybox
    """
    When I run `lc dc config`
    Then it should succeed
    And the output should contain 'image: example/lein'
    And the output should not contain 'image: clojure:lein-2.6.1'

  Scenario: using default make template image with a v2 compose file
    Given a file named "lc.yml" with:
    """yaml
    name: testing
    template: make
    """
    And a file named "docker-compose.yml" with:
    """yaml
      version: '2'
      services:
        no-op:
          image: busybox
    """
    When I run `lc dc config`
    Then it should succeed
    And the output should contain 'image: gcc:6.1'

  Scenario: using overridden make template image with a v2 compose file
    Given a file named "lc.yml" with:
    """yaml
    name: testing
    template: make
    template_image: example/make
    """
    And a file named "docker-compose.yml" with:
    """yaml
      version: '2'
      services:
        no-op:
          image: busybox
    """
    When I run `lc dc config`
    Then it should succeed
    And the output should contain 'image: example/make'
    And the output should not contain 'image: gcc:6.1'



