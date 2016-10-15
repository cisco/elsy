package template

import (
	"bytes"
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

// addV1 will add the given compose v1 template to a registry for use by external packages
func addV1(template template) error {
	if _, ok := templatesV1[template.name]; ok {
		return fmt.Errorf("template %q already exists", template.name)
	}
	templatesV1[template.name] = template
	return nil
}

// GetV1 will return the compose V1 template if it exists
// If 'enableScratchVolume' is true and the target template supports
// scratch-space optimization then Get will enable it.
func GetV1(name string, enableScratchVolume bool) (string, error) {
	tmpl, present := templatesV1[name]
	if !present {
		return "", fmt.Errorf("template %q is not registered", name)
	}
	yml, err := tmpl.toYml(enableScratchVolume)
	if err != nil {
		return "", err
	}
	return yml, nil
}

// List will return a slice of all known templates
func List() []string {
	keys := make([]string, 0, len(templatesV1))

	for k := range templatesV1 {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}
