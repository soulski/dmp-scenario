FROM dmp

RUN rm -rf /var/lib/apt/lists/*
RUN apt-get update && apt-get install -y supervisor \
    sqlite3 \
    libsqlite3-dev
