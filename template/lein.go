package template

import "github.com/elsy/helpers"

var leinTemplateV1 = template{
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
clean:
  <<: *lein
  entrypoint: [lein, clean]
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

var leinTemplateV2 = template{
	name: "lein",
	composeYmlTmpl: `
version: '2'
services:
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
      - container:lc_shared_mvndata
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
  clean:
    <<: *lein
    entrypoint: [lein, clean]
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
	addSharedExternalDataContainer("lein", helpers.DockerDataContainer{
		Image:     "busybox:latest",
		Name:      "lc_shared_mvndata",
		Volumes:   []string{"/root/.m2/repository"},
		Resilient: true,
	})

	addV1(leinTemplateV1)
	addV2(leinTemplateV2)
}
