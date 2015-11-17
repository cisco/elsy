FROM golang:latest
RUN GOOS=darwin GOARCH=amd64 go install -v std

COPY dependencies go-gets go-installs /opt/build/
RUN /opt/build/go-gets
RUN /opt/build/go-installs

COPY publish package /opt/lifecycle/
