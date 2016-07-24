FROM library/ruby:latest

RUN apt-get update && apt-get install -y netcat

RUN gem install \
    turnip

## setup git so blackbox tests can make git calls
RUN git config --global user.email "test@lancopetest.com"
RUN git config --global user.name "Project Lifecycle"

ENV DOCKER_VERSION 1.8.3
RUN curl -so /usr/bin/docker https://get.docker.com/builds/Linux/x86_64/docker-${DOCKER_VERSION}
RUN chmod a+x /usr/bin/docker

ENV DOCKER_COMPOSE_VERSION=1.5.2
RUN curl -Lo /usr/local/bin/docker-compose https://github.com/docker/compose/releases/download/$DOCKER_COMPOSE_VERSION/docker-compose-Linux-x86_64
RUN chmod a+x /usr/local/bin/docker-compose
