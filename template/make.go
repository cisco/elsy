package template

var makeTemplate = template{
	name: "make",
	composeYmlTmpl: `
make: &make
  image: gcc:6.1
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
