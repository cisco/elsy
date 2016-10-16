package template

import "github.com/elsy/helpers"

var mvnTemplateV1 = template{
	name: "mvn",
	composeYmlTmpl: `
{{if .ScratchVolumes}}
mvnscratch:
  image: busybox
  volumes:
    {{.ScratchVolumes}}
  entrypoint: /bin/true
{{end}}
mvn: &mvn
  image: maven:3.2-jdk-8
  volumes:
    - ./:/opt/project
  working_dir: /opt/project
  entrypoint: mvn
  volumes_from:
    - lc_shared_mvndata
{{if .ScratchVolumes}}
    - mvnscratch
{{end}}
test:
  <<: *mvn
  entrypoint: [mvn, test]
package:
  <<: *mvn
  command: [package, "-DskipTests=true"]
publish:
  <<: *mvn
  entrypoint: /bin/true
{{if .ScratchVolumes}}
clean:
  <<: *mvn
  entrypoint: [mvn, clean, "-Dmaven.clean.failOnError=false"]
{{else}}
clean:
  <<: *mvn
  entrypoint: [mvn, clean]
{{end}}
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

var mvnTemplateV2 = template{
	name: "mvn",
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
  mvn: &mvn
    image: maven:3.2-jdk-8
    volumes:
      - ./:/opt/project
    working_dir: /opt/project
    entrypoint: mvn
    volumes_from:
      - container:lc_shared_mvndata
  {{if .ScratchVolumes}}
      - mvnscratch
  {{end}}
  test:
    <<: *mvn
    entrypoint: [mvn, test]
  package:
    <<: *mvn
    command: [package, "-DskipTests=true"]
  publish:
    <<: *mvn
    entrypoint: /bin/true
  {{if .ScratchVolumes}}
  clean:
    <<: *mvn
    entrypoint: [mvn, clean, "-Dmaven.clean.failOnError=false"]
  {{else}}
  clean:
    <<: *mvn
    entrypoint: [mvn, clean]
  {{end}}
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
	addSharedExternalDataContainer("mvn", helpers.DockerDataContainer{
		Image:     "busybox:latest",
		Name:      "lc_shared_mvndata",
		Volumes:   []string{"/root/.m2/repository"},
		Resilient: true,
	})

	addV1(mvnTemplateV1)
	addV2(mvnTemplateV2)
}
