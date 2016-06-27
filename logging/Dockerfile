FROM scenario

RUN mkdir -p /usr/logging

COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf
COPY /bin/logging /usr/bin/

EXPOSE 80

ENTRYPOINT ["/usr/bin/supervisord"]
