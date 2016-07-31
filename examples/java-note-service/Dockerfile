FROM library/java:8-jre

RUN apt-get update && apt-get install -y netcat

COPY ./target/java-note-service-1.0-SNAPSHOT.jar /opt/service/java-note-service.jar
COPY ./config.yml /opt/service/config.yml

COPY ./init.d /opt/init.d
WORKDIR /opt/service

EXPOSE 8080
ENTRYPOINT ["/opt/init.d/server-entrypoint.sh"]
CMD ["server", "config.yml"]
