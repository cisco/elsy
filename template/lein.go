package template

import "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"

var leinTemplate = template{
	name: "lein",
	composeYmlTmpl: `
{{if .ScratchVolumes}}
mvnscratch:
  image: busybox
  volumes:
    {{.ScratchVolumes}}
  entrypoint: /bin/true
{{end}}
lein: &lein
  image: clojure:lein-2.6.1
  volumes:
    - ./:/opt/project
  working_dir: /opt/project
  entrypoint: lein
  volumes_from:
    - lc_shared_mvndata
{{if .ScratchVolumes}}
    - mvnscratch
{{end}}
test:
  <<: *lein
  entrypoint: [lein, test]
package:
  <<: *lein
  command: [jar, "-DskipTests=true"]
publish:
  <<: *lein
  entrypoint: /bin/true
`,
	scratchVolumes: `
  - /opt/project/target/classes
  - /opt/project/target/journal
  - /opt/project/target/maven-archiver
  - /opt/project/target/maven-status
  - /opt/project/target/snapshots
  - /opt/project/target/test-classes
  - /opt/project/target/war/work
  - /opt/project/target/webappDirectory
`}

func init() {
	AddSharedExternalDataContainer("lein", helpers.DockerDataContainer{
		Image:     "busybox:latest",
		Name:      "lc_shared_mvndata",
		Volumes:   []string{"/root/.m2/repository"},
		Resilient: true,
	})

	Add(leinTemplate)
}
