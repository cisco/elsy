FROM library/golang:latest
RUN GOOS=darwin GOARCH=amd64 go install -v std

RUN go get -u github.com/kardianos/govendor

COPY test package /opt/lifecycle/
