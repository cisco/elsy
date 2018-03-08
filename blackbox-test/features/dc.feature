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

Feature: passing through to docker-compose

Scenario: with a standard project
  Given a file named "docker-compose.yml" with:
  """yaml
  foo:
    image: busybox
    command: echo bar
  """
  When I run `lc dc -- run --rm foo`
  Then it should succeed with "bar"

Scenario: retaining exit code
  Given a file named "docker-compose.yml" with:
  """yaml
  foo:
    image: busybox
    command: ["sh", "-c", "exit 42"]
  """
  When I run `lc dc -- run --rm foo`
  Then it should fail with exit code 42
