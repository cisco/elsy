FROM library/ruby:latest
RUN apt-get update && \
    apt-get install -y \
      netcat-openbsd && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN gem install \
    turnip \
    turnip_formatter \
    rspec_junit_formatter \
    byebug \
    rspec-instafail

ENV DOCKER_VERSION 1.11.2
RUN curl -L https://get.docker.com/builds/Linux/x86_64/docker-$DOCKER_VERSION.tgz | tar -xzvf - --strip-components=1 -C /usr/bin docker/docker
RUN chmod a+x /usr/bin/docker

ENV DOCKER_COMPOSE_VERSION 1.8.1
RUN curl -Lo /usr/local/bin/docker-compose https://github.com/docker/compose/releases/download/$DOCKER_COMPOSE_VERSION/docker-compose-Linux-x86_64
RUN chmod a+x /usr/local/bin/docker-compose

## setup git so blackbox tests can make git calls
RUN git config --global user.email "test@ciscotest.com"
RUN git config --global user.name "Project Lifecycle"

WORKDIR /opt/project
ENTRYPOINT ["rspec", "--require", "./blackbox-test/spec_helper.rb"]
CMD ["--pattern", "blackbox-test/features/**/*.feature"]
