FROM scenario

RUN mkdir -p /usr/gateway/assets && \
    mkdir -p /usr/gateway/templates

COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf
COPY /bin/gateway /usr/bin/
COPY /assets/ /usr/gateway/assets/
COPY /templates /usr/gateway/templates/

EXPOSE 80

ENTRYPOINT ["/usr/bin/supervisord"]
