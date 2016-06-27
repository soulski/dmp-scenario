FROM scenario

RUN mkdir -p /usr/user

COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf
COPY /bin/user /usr/bin/

EXPOSE 80

ENTRYPOINT ["/usr/bin/supervisord"]
