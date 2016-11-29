/*
 *  Copyright 2016 Cisco Systems, Inc.
 *  
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *  
 *  http://www.apache.org/licenses/LICENSE-2.0
 *  
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package template

var makeTemplateV1 = template{
	name: "make",
	composeYmlTmpl: `
make: &make
  image: gcc:6.1
  volumes:
    - ./:/opt/project
  working_dir: /opt/project
  entrypoint: make
test:
  <<: *make
  entrypoint: [make, test]
clean:
  <<: *make
  entrypoint: [make, clean]
`}

var makeTemplateV2 = template{
	name: "make",
	composeYmlTmpl: `
version: '2'
services:
  make: &make
    image: gcc:6.1
    volumes:
      - ./:/opt/project
    working_dir: /opt/project
    entrypoint: make
  test:
    <<: *make
    entrypoint: [make, test]
  clean:
    <<: *make
    entrypoint: [make, clean]
`}

func init() {
	addV1(makeTemplateV1)
	addV2(makeTemplateV2)
}
