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

Feature: system upgrade task
  Scenario: try to upgrade snapshot build
    When I run `lc system upgrade`
    Then it should fail with "Upgrade not available on snapshot versions"

  Scenario: upgrade a non-snapshot build
    Given I run `mkdir -p /opt/bin`
    And I run `cp /opt/project/target/lc-blackbox /opt/bin/lc-test`
    And I run `chmod 755 /opt/bin/lc-test && ln -s /opt/bin/lc-test /opt/bin/lc`
    And I run `/opt/project/blackbox-test/bin/patch-strings-in-binary /opt/bin/lc-test "$(/opt/bin/lc --version|awk '{print $3}')" "v1.9.9"`
    When I run `/opt/bin/lc system upgrade`
    Then it should succeed
    And the output should contain all of these:
        | Upgrading to |
        | Done!        |
        | 100.00%      |
    When I run `/opt/bin/lc --version`
    Then it should succeed
    And the output should not contain "test"
    Then I run `rm -rf /opt/bin`

  Scenario: try to upgrade same version
    Given I run `mkdir -p /opt/bin`
    And I run `cp /opt/project/target/lc-blackbox /opt/bin/lc-test`
    And I run `chmod 755 /opt/bin/lc-test && ln -s /opt/bin/lc-test /opt/bin/lc`
    And I run `/opt/project/blackbox-test/bin/patch-strings-in-binary /opt/bin/lc-test "$(/opt/bin/lc --version|awk '{print $3}')" $(curl https://api.github.com/repos/cisco/elsy/releases/latest 2>/dev/null|grep tag_name|awk '{print $2}'|sed "s/\"\|,//g")`
    When I run `/opt/bin/lc system upgrade`
    Then it should succeed
    And the output should contain "No new version available"
