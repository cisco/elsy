FROM library/ruby:latest

EXPOSE 9090

RUN apt-get update && apt-get install -y netcat

ENV PHANTOMJS_VERSION 2.1.1
RUN \
  apt-get update && \
  apt-get install -y libfreetype6 libfontconfig && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/* && \
  mkdir -p /srv/var && \
  curl -sL https://bitbucket.org/ariya/phantomjs/downloads/phantomjs-$PHANTOMJS_VERSION-linux-x86_64.tar.bz2 | tar -C /tmp -xj && \
  mv /tmp/phantomjs-$PHANTOMJS_VERSION-linux-x86_64/ /opt/phantomjs && \
  ln -s /opt/phantomjs/bin/phantomjs /usr/local/bin/phantomjs

RUN gem install \
    turnip \
    poltergeist \
    gnawrnip \
    rspec
