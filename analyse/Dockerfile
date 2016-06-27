FROM scenario

RUN mkdir -p /usr/analyse

COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf
COPY /bin/analyse /usr/bin/

EXPOSE 80

ENTRYPOINT ["/usr/bin/supervisord"]
