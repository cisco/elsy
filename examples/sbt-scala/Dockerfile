# FROM hseeberger/scala-sbt
FROM library/java:8-jre

COPY target/scala-2.11/hello-akka-assembly-1.0.jar /opt/app.jar

COPY ./init.d /opt/init.d
WORKDIR /opt/service

EXPOSE 8080

ENTRYPOINT ["/opt/init.d/server-entrypoint.sh"]
# CMD ["server", "config.yml"]
