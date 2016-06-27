FROM scenario

RUN mkdir -p /usr/store

COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf
COPY /bin/store /usr/bin/

EXPOSE 80

ENTRYPOINT ["/usr/bin/supervisord"]
