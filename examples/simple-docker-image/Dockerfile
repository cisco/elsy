FROM library/alpine:latest

RUN apk update && apk add lighttpd
RUN echo "Lighttpd is running..." > /var/www/localhost/htdocs/index.html

EXPOSE 80
ENTRYPOINT ["/usr/sbin/lighttpd", "-D", "-f", "/etc/lighttpd/lighttpd.conf"]
