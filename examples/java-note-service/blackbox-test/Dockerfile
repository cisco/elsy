FROM library/ruby:latest

RUN apt-get update && apt-get install -y netcat mysql-client

RUN gem install \
    turnip \
    excon \
    multi_json
