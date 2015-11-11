FROM ruby:latest

RUN gem install \
    turnip

ENV DOCKER_VERSION 1.7.1
RUN curl -so /usr/bin/docker https://get.docker.com/builds/Linux/x86_64/docker-${DOCKER_VERSION}
RUN chmod a+x /usr/bin/docker

ENV DOCKER_COMPOSE_VERSION=1.5.0
RUN curl -Lo /usr/local/bin/docker-compose https://github.com/docker/compose/releases/download/$DOCKER_COMPOSE_VERSION/docker-compose-Linux-x86_64
RUN chmod a+x /usr/local/bin/docker-compose
