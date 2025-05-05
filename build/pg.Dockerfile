FROM postgres:16.2

RUN apt-get update && \
    apt-get install -y postgresql-16-cron && \
    echo "shared_preload_libraries = 'pg_cron'" >> /usr/share/postgresql/postgresql.conf.sample

RUN mkdir -p /usr/share/postgresql/16/tsearch_data
RUN chmod 777 /usr/share/postgresql/16/tsearch_data