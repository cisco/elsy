package template

var makeTemplate = template{
	name: "make",
	composeYmlTmpl: `
make: &make
  image: arch-docker.eng.lancope.local:5000/c-dev-env:v1.0.0
  volumes:
    - ./:/opt/project
  working_dir: /opt/project
  entrypoint: make
test:
  <<: *make
  entrypoint: [make, test]
clean:
  <<: *make
  entrypoint: [make, clean]
`}

func init() {
	Add(makeTemplate)
}
