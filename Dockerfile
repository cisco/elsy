FROM golang:latest
RUN go get -v github.com/codegangsta/cli
RUN GOOS=darwin GOARCH=amd64 go install -v github.com/codegangsta/cli std
