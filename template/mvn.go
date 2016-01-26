package template

import "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"

var mvnTemplate = template{
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
{{if not .ScratchVolumes}}
teardown:
  <<: *mvn
  command: [clean]
{{end}}
`,
  scratchVolumes: `
  - /opt/project/target/classes
  - /opt/project/target/journal
  - /opt/project/target/maven-archiver
  - /opt/project/target/maven-status
  - /opt/project/target/snapshots
  - /opt/project/target/test-classes
`,}

func init() {
  AddSharedExternalDataContainer("mvn", helpers.DockerDataContainer{
    Image: "busybox:latest",
    Name: "lc_shared_mvndata",
    Volumes: []string{"/root/.m2/repository"},
    Resilient: true,
  })

  Add(mvnTemplate)
}
