package template

var emberTemplate = template{
	name: "ember",
	composeYmlTmpl: `
  emberdata:
    image: arch-docker.eng.lancope.local:5000/ember
    volumes:
     - /opt/project/bower_components
     - /opt/project/dist
     - /opt/project/node_modules
     - /opt/project/vendor
     - /opt/project/tmp
    labels:
      com.lancope.docker-gc.keep: "True"
    entrypoint: "/bin/true"
  ember: &ember
    image: arch-docker.eng.lancope.local:5000/ember
    volumes:
     - .:/opt/project
    working_dir: /opt/project
    entrypoint: /usr/local/bin/ember
    volumes_from:
     - emberdata
  npm:
    <<: *ember
    entrypoint: /usr/local/bin/npm
  bower:
    <<: *ember
    entrypoint: /usr/local/bin/bower
  installdependencies:
    <<: *ember
    entrypoint: bash
    command: -c "npm install && npm update && bower install && bower update"
  test:
    <<: *ember
    command: [test]
  package:
    <<: *ember
    command: [build, "--environment='production'",  "--output-path=dist-production"]
  `,
}

func init() {
	Add(emberTemplate)
}
