package template

import (
	"bytes"
	"fmt"
	"sort"
	"github.com/elsy/helpers"
	tmpl "text/template"
)

var sharedExternalDataContainers = make(map[string][]helpers.DockerDataContainer)

func AddSharedExternalDataContainer(templateName string, ddc helpers.DockerDataContainer) {
	dataContainers := GetSharedExternalDataContainers(templateName)
	sharedExternalDataContainers[templateName] = append(dataContainers, ddc)
}

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

var templates = make(map[string]template)

// Add will add the given template to a registry for use by external packages
func Add(template template) error {
	if _, ok := templates[template.name]; ok {
		return fmt.Errorf("template %q already exists", template.name)
	}
	templates[template.name] = template
	return nil
}

// Get will return the template if it exists
// If 'enableScratchVolume' is true and the target template supports
// scratch-space optimization then Get will enable it.
func Get(name string, enableScratchVolume bool) (string, error) {
	template, present := templates[name]
	if !present {
		return "", fmt.Errorf("template %q is not registered", name)
	}
	yml, err := template.toYml(enableScratchVolume)
	if err != nil {
		return "", err
	}
	return yml, nil
}

func List() []string {
	keys := make([]string, 0, len(templates))

	for k := range templates {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}
