FROM library/node:6.3

ENV EMBER_CLI_VERSION=2.7

RUN npm install -g ember-cli@$EMBER_CLI_VERSION \
    bower \
    phantomjs

## watchman is used for live-reloading server
ENV WATCHMAN_VERSION=v3.9.0
RUN \
  cd /tmp &&\
	git clone https://github.com/facebook/watchman.git &&\
	cd watchman &&\
	git checkout $WATCHMAN_VERSION &&\
	./autogen.sh &&\
	./configure --without-python &&\
	make -j2 &&\
	make install &&\
  cd / &&\
  rm -rf /tmp/watchman
