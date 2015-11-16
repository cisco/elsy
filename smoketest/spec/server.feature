Feature: server task

  Scenario: running a devserver
  Given a file named "docker-compose.yml" with:
  """yaml
  emberdata:
    image: arch-docker.eng.lancope.local:5000/ember
    volumes:
     - /opt/app/bower_components
     - /opt/app/dist
     - /opt/app/node_modules
     - /opt/app/vendor
     - /opt/app/tmp
    entrypoint: "/bin/true"
  unison:
    image: leighmcculloch/unison
    ports:
     - "5000"
    volumes:
     - /opt/app/app
     - /opt/app/config
     - /opt/app/public
     - /opt/app/tests
  defaults: &defaults
    image: arch-docker.eng.lancope.local:5000/ember
    volumes_from:
     - emberdata
    volumes:
     - .:/opt/app
    entrypoint: /usr/local/bin/ember
  ember:
    <<: *defaults
  npm:
    <<: *defaults
    entrypoint: /usr/local/bin/npm
  bower:
    <<: *defaults
    entrypoint: /usr/local/bin/bower
  testserver:
    <<: *defaults
    volumes_from:
     - emberdata
     - unison
    command: test --server --host 0.0.0.0
    ports:
     - "7357:7357"
  devserver:
    <<: *defaults
    volumes_from:
     - emberdata
     - unison
    command: serve -p 80 --live-reload false
    ports:
     - "80"
  prodserver:
    image: vision15-iris-ui
    ports:
     - "80"
  """
  When I run `lc server start`
  Then it should report a correct address

  Scenario: stopping a running devserver
  Given a file named "docker-compose.yml" with:
  """yaml
  emberdata:
    image: arch-docker.eng.lancope.local:5000/ember
    volumes:
     - /opt/app/bower_components
     - /opt/app/dist
     - /opt/app/node_modules
     - /opt/app/vendor
     - /opt/app/tmp
    entrypoint: "/bin/true"
  unison:
    image: leighmcculloch/unison
    ports:
     - "5000"
    volumes:
     - /opt/app/app
     - /opt/app/config
     - /opt/app/public
     - /opt/app/tests
  defaults: &defaults
    image: arch-docker.eng.lancope.local:5000/ember
    volumes_from:
     - emberdata
    volumes:
     - .:/opt/app
    entrypoint: /usr/local/bin/ember
  ember:
    <<: *defaults
  npm:
    <<: *defaults
    entrypoint: /usr/local/bin/npm
  bower:
    <<: *defaults
    entrypoint: /usr/local/bin/bower
  testserver:
    <<: *defaults
    volumes_from:
     - emberdata
     - unison
    command: test --server --host 0.0.0.0
    ports:
     - "7357:7357"
  devserver:
    <<: *defaults
    volumes_from:
     - emberdata
     - unison
    command: serve -p 80 --live-reload false
    ports:
     - "80"
  prodserver:
    image: vision15-iris-ui
    ports:
     - "80"
  """
  When I run `lc server start && lc server stop`
  Then it should report "Stopping devserver"

  Scenario: trying to get status when server not running
  Given a file named "docker-compose.yml" with:
  """yaml
  emberdata:
    image: arch-docker.eng.lancope.local:5000/ember
    volumes:
     - /opt/app/bower_components
     - /opt/app/dist
     - /opt/app/node_modules
     - /opt/app/vendor
     - /opt/app/tmp
    entrypoint: "/bin/true"
  unison:
    image: leighmcculloch/unison
    ports:
     - "5000"
    volumes:
     - /opt/app/app
     - /opt/app/config
     - /opt/app/public
     - /opt/app/tests
  defaults: &defaults
    image: arch-docker.eng.lancope.local:5000/ember
    volumes_from:
     - emberdata
    volumes:
     - .:/opt/app
    entrypoint: /usr/local/bin/ember
  ember:
    <<: *defaults
  npm:
    <<: *defaults
    entrypoint: /usr/local/bin/npm
  bower:
    <<: *defaults
    entrypoint: /usr/local/bin/bower
  testserver:
    <<: *defaults
    volumes_from:
     - emberdata
     - unison
    command: test --server --host 0.0.0.0
    ports:
     - "7357:7357"
  devserver:
    <<: *defaults
    volumes_from:
     - emberdata
     - unison
    command: serve -p 80 --live-reload false
    ports:
     - "80"
  prodserver:
    image: vision15-iris-ui
    ports:
     - "80"
  """
  When I run `lc server stat`
  Then it should report "devserver: down"
