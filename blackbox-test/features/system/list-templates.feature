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

Feature: system list-templates task

  Scenario: calling with no args
    When I run `lc system list-templates`
    Then it should succeed
    And the output should contain all of these:
      | Run `lc system view-template <template-name>` to see the template contents. |
      | Compose v1 Templates: |
      | mvn |
      | sbt |
      | Compose v2 Templates: |
