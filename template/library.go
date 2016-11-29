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

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	tmpl "text/template"

	"github.com/elsy/helpers"
)

var sharedExternalDataContainers = make(map[string][]helpers.DockerDataContainer)

func addSharedExternalDataContainer(templateName string, ddc helpers.DockerDataContainer) {
	dataContainers := GetSharedExternalDataContainers(templateName)
	sharedExternalDataContainers[templateName] = append(dataContainers, ddc)
}

// GetSharedExternalDataContainers will return a slice of data containers used by the given template
func GetSharedExternalDataContainers(templateName string) []helpers.DockerDataContainer {
	return sharedExternalDataContainers[templateName]
}

// template contains the data necessary to construct a lc template
type template struct {
	name           string
	composeYmlTmpl string
	scratchVolumes string
}

// toYml will take a template and prepare it for use by other packages
// the string returned will be a valid docker-compose.yml string.
func (t *template) toYml(enableScratchVolume bool) (string, error) {
	goTemplate, err := tmpl.New("composeTemplate").Parse(t.composeYmlTmpl)
	if err != nil {
		return "", fmt.Errorf("could not parse docker-compose yml for template %q, error: %q", t.name, err)
	}

	var scratchVolumes string
	if enableScratchVolume {
		scratchVolumes = t.scratchVolumes
	}
	data := struct {
		ScratchVolumes string
	}{
		scratchVolumes,
	}
	var finalYml bytes.Buffer
	err = goTemplate.Execute(&finalYml, data)
	if err != nil {
		return "", fmt.Errorf("could not interpolate docker-compose yml for template %q, error: %q", t.name, err)
	}
	return finalYml.String(), nil
}

var templatesV1 = make(map[string]template)
var templatesV2 = make(map[string]template)

// addV1 will add the given compose v1 template to a registry for use by external packages
func addV1(template template) error {
	if _, ok := templatesV1[template.name]; ok {
		return fmt.Errorf("template %q already exists", template.name)
	}
	templatesV1[template.name] = template
	return nil
}

// addV2 will add the given compose v2 template to a registry for use by external packages
func addV2(template template) error {
	if _, ok := templatesV2[template.name]; ok {
		return fmt.Errorf("template %q already exists", template.name)
	}
	templatesV2[template.name] = template
	return nil
}

// GetTemplate returns the template if found
// will inspect the compose file at "docker-comopse.yml" to determine the correct
// file version to use when retrieving the template
//
// currently defaults to V1 if no version found
func GetTemplate(templateName string, enableScratchVolume bool) (string, error) {
	version := helpers.GetComposeFileVersion("docker-compose.yml", helpers.V1)
	switch version {
	case helpers.V1:
		return GetV1(templateName, enableScratchVolume)
	case helpers.V2:
		return GetV2(templateName, enableScratchVolume)
	}
	return "", errors.New("could not determine docker-compose.yml file version")
}

// GetV1 will return the compose V1 template if it exists
// If 'enableScratchVolume' is true and the target template supports
// scratch-space optimization then Get will enable it.
func GetV1(name string, enableScratchVolume bool) (string, error) {
	tmpl, present := templatesV1[name]
	if !present {
		return "", fmt.Errorf("template %q is not a registered v1 template", name)
	}
	yml, err := tmpl.toYml(enableScratchVolume)
	if err != nil {
		return "", err
	}
	return yml, nil
}

// GetV2 will return the compose V2 template if it exists
// If 'enableScratchVolume' is true and the target template supports
// scratch-space optimization then Get will enable it.
func GetV2(name string, enableScratchVolume bool) (string, error) {
	tmpl, present := templatesV2[name]
	if !present {
		return "", fmt.Errorf("template %q is not a registered v2 template", name)
	}
	yml, err := tmpl.toYml(enableScratchVolume)
	if err != nil {
		return "", err
	}
	return yml, nil
}

// ListV1 will return a slice of all known v1 templates
func ListV1() []string {
	keys := make([]string, 0, len(templatesV1))

	for k := range templatesV1 {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}

// ListV2 will return a slice of all known v2 templates
func ListV2() []string {
	keys := make([]string, 0, len(templatesV2))

	for k := range templatesV1 {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}
