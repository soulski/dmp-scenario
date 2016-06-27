FROM scenario

RUN mkdir -p /usr/authen

COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf
COPY /bin/authen /usr/bin/

EXPOSE 80

ENTRYPOINT ["/usr/bin/supervisord"]
