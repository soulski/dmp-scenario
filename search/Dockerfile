FROM scenario

RUN mkdir -p /usr/search

COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf
COPY /bin/search /usr/bin/

EXPOSE 80

ENTRYPOINT ["/usr/bin/supervisord"]
