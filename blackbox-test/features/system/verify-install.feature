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

Feature: system verify-install task

## We can't really test the disk mounting repeatedly, so just verify that the file check works
## and assume the disk mounting is always working (it will be in ci since it is all linux)

  Scenario: calling with missing lc.yml file
    When I run `lc system verify-install`
    Then it should succeed
    And the output should contain "could not find 'lc.yml' in the current directory, skipping"

  Scenario: calling with lc.yml file
    Given a file named "lc.yml" with:
    """yaml
    name: test-verify-install
    """
    When I run `lc system verify-install`
    Then it should succeed
    And the output should not contain "could not find 'lc.yml' in the current directory, skipping"
