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

import "testing"

func TestMvnTemplate(t *testing.T) {
	if dataContainers := GetSharedExternalDataContainers("mvn"); len(dataContainers) != 1 {
		t.Error("expected mvn template to register one shared external data container")
	}

	if _, err := GetV1("mvn", false); err != nil {
		t.Error("expected mvn template to be registered")
	}

	if _, err := GetV2("mvn", false); err != nil {
		t.Error("expected mvn template to be registered")
	}
}
