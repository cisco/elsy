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

Feature: make command

    Scenario: compile a valid C program using a Makefile
        Given a file named "docker-compose.yml" with:
        """yaml
        make:
          image: gcc:6.1
          volumes:
            - ./:/project
          working_dir: /project
          command: "make"
        """
        And a file named "foo.c" with:
        """
        #include <stdio.h>
        int main() {
            printf("foo\n");
            return 0;
        }
        """
        And a file named "Makefile" with:
        """
        .RECIPEPREFIX = >

        foo: foo.c
        > gcc -o foo foo.c
        """
        When I run `lc make`
        Then it should succeed
        And the file "foo" should be executable

  Scenario: compile a valid C program using a Makefile and a compose v2 file
      Given a file named "docker-compose.yml" with:
      """yaml
      version: '2'
      services:  
        make:
          image: gcc:6.1
          volumes:
            - ./:/project
          working_dir: /project
          command: "make"
      """
      And a file named "foo.c" with:
      """
      #include <stdio.h>
      int main() {
          printf("foo\n");
          return 0;
      }
      """
      And a file named "Makefile" with:
      """
      .RECIPEPREFIX = >

      foo: foo.c
      > gcc -o foo foo.c
      """
      When I run `lc make`
      Then it should succeed
      And the file "foo" should be executable

    Scenario: compile an invalid C program using a Makefile
        Given a file named "docker-compose.yml" with:
        """yaml
        make:
          image: gcc:6.1
          volumes:
            - ./:/project
          working_dir: /project
          command: "make"
        """
        And a file named "foo.c" with:
        """
        #include <stdio.h>
        int main() {
            printf("foo\n");XXX
            return 0;
        }
        """
        And a file named "Makefile" with:
        """
        .RECIPEPREFIX = >

        foo: foo.c
        > gcc -fno-diagnostics-color -o foo foo.c
        """
        When I run `lc make`
        Then it should fail with "error: 'XXX' undeclared"

    Scenario: try to compile without a make service
        Given a file named "docker-compose.yml" with:
        """yaml
        test:
          image: busybox
          command: /bin/true
        """
        And a file named "foo.c" with:
        """
        #include <stdio.h>
        int main() {
            printf("foo\n");XXX
            return 0;
        }
        """
        And a file named "Makefile" with:
        """
        .RECIPEPREFIX = >

        foo: foo.c
        > gcc -o foo foo.c
        """
        When I run `lc make`
        Then it should fail with "No such service: make"

    Scenario: try to compile without a make service, but with template: make
        Given a file named "lc.yml" with:
        """
        template: make
        """
        And a file named "foo.c" with:
        """
        #include <stdio.h>
        int main() {
            printf("foo\n");
            return 0;
        }
        """
        And a file named "Makefile" with:
        """
        .RECIPEPREFIX = >

        foo: foo.c
        > gcc -o foo foo.c
        """
        When I run `lc make`
        Then it should succeed

    Scenario: correct make template
      When I run `lc system view-template make`
      Then it should succeed
      And the output should contain all of these:
        | make:     |
        | test:    |
