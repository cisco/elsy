package template

var makeTemplateV1 = template{
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

var makeTemplateV2 = template{
	name: "make",
	composeYmlTmpl: `
version: '2'
services:
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
	addV1(makeTemplateV1)
	addV2(makeTemplateV2)
}
