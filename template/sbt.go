package template

import "github.com/elsy/helpers"

var sbtTemplateV1 = template{
	name: "sbt",
	composeYmlTmpl: `
{{if .ScratchVolumes}}
sbtscratch:
  image: busybox
  command: /bin/true
  volumes:
    {{.ScratchVolumes}}
{{end}}
sbt: &sbt
  image: arch-docker.eng.lancope.local:5000/sbt
  volumes:
    - ./:/opt/project
  working_dir: /opt/project
  entrypoint: sbt
  volumes_from:
    - lc_shared_sbtdata
{{if .ScratchVolumes}}
    - sbtscratch
{{end}}
test:
  <<: *sbt
  entrypoint: [sbt, test]
package:
  <<: *sbt
  command: [assembly]
publish:
  <<: *sbt
  entrypoint: /bin/true
clean:
  <<: *sbt
  entrypoint: [sbt, clean]
`,
	scratchVolumes: `
    - /opt/project/target/resolution-cache
    - /opt/project/target/scala-2.11/classes
    - /opt/project/target/scala-2.11/test-classes
    - /opt/project/target/streams
    - /opt/project/project/project
    - /opt/project/project/target
`}

var sbtTemplateV2 = template{
	name: "sbt",
	composeYmlTmpl: `
version: '2'
services:
  {{if .ScratchVolumes}}
  sbtscratch:
    image: busybox
    command: /bin/true
    volumes:
      {{.ScratchVolumes}}
  {{end}}
  sbt: &sbt
    image: arch-docker.eng.lancope.local:5000/sbt
    volumes:
      - ./:/opt/project
    working_dir: /opt/project
    entrypoint: sbt
    volumes_from:
      - container:lc_shared_sbtdata
  {{if .ScratchVolumes}}
      - sbtscratch
  {{end}}
  test:
    <<: *sbt
    entrypoint: [sbt, test]
  package:
    <<: *sbt
    command: [assembly]
  publish:
    <<: *sbt
    entrypoint: /bin/true
  clean:
    <<: *sbt
    entrypoint: [sbt, clean]
`,
	scratchVolumes: `
      - /opt/project/target/resolution-cache
      - /opt/project/target/scala-2.11/classes
      - /opt/project/target/scala-2.11/test-classes
      - /opt/project/target/streams
      - /opt/project/project/project
      - /opt/project/project/target
`}

func init() {
	addSharedExternalDataContainer("sbt", helpers.DockerDataContainer{
		Image:     "busybox:latest",
		Name:      "lc_shared_sbtdata",
		Volumes:   []string{"/root/.ivy2"},
		Resilient: true,
	})

	addV1(sbtTemplateV1)
	addV2(sbtTemplateV2)
}
