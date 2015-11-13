package template

import "stash0.eng.lancope.local/dev-infrastructure/project-lifecycle/helpers"

func init() {
  AddSharedExternalDataContainer("sbt", helpers.DockerDataContainer{
    Image: "busybox:latest",
    Name: "lc_shared_sbtdata",
    Volumes: []string{"/root/.ivy2"},
    Resilient: true,
  })

  Add("sbt", `
sbtscratch:
  image: busybox
  command: /bin/true
  volumes:
    - /opt/project/target/resolution-cache
    - /opt/project/target/scala-2.11/classes
    - /opt/project/target/scala-2.11/test-classes
    - /opt/project/target/streams
    - /opt/project/project/project
    - /opt/project/project/target

sbt: &sbt
  image: arch-docker.eng.lancope.local:5000/sbt
  volumes:
    - ./:/opt/project
  working_dir: /opt/project
  entrypoint: sbt
  volumes_from:
    - lc_shared_sbtdata
    - sbtscratch
test:
  <<: *sbt
  entrypoint: [sbt, test]
package:
  <<: *sbt
  command: [assembly]
`)
}
