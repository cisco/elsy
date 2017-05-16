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

Feature: resolve-docker-tag task
  Scenario: with an empty project, calling resolve-docker-tag without the git branch
    When I run `lc resolve-docker-tag`
    Then it should fail with "expecting a git branch and/or a git tag be set, found neither"

  Scenario: calling resolve-docker-tag on the master branch
    When I run `lc resolve-docker-tag --git-branch origin/master`
    Then it should succeed with "latest"

  Scenario: calling resolve-docker-tag on a release branch
    When I run `lc resolve-docker-tag --git-branch origin/release/1.5`
    Then it should succeed with "1.5"

  Scenario: calling resolve-docker-tag on a feature branch
    When I run `lc resolve-docker-tag --git-branch origin/fix-thing`
    Then it should succeed with "snapshot.fix-thing"

  Scenario: calling resolve-docker-tag on a tagged commit
    When I run `lc resolve-docker-tag --git-tag v1.2.3`
    Then it should succeed with "v1.2.3"

  Scenario: calling resolve-docker-tag on a tagged master branch
    When I run `lc resolve-docker-tag --git-tag v1.2.3 --git-branch=origin/master`
    Then it should succeed with "v1.2.3"
