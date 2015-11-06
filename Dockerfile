FROM golang:latest
RUN GOOS=darwin GOARCH=amd64 go install -v std

RUN go get -v github.com/codegangsta/cli
RUN GOOS=darwin GOARCH=amd64 go install -v github.com/codegangsta/cli

RUN go get -v gopkg.in/yaml.v2
RUN GOOS=darwin GOARCH=amd64 go install -v gopkg.in/yaml.v2

RUN go get -v github.com/fsouza/go-dockerclient
RUN GOOS=darwin GOARCH=amd64 go install -v github.com/fsouza/go-dockerclient
