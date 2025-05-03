FROM postgres:15.0

RUN mkdir -p /usr/share/postgresql/tsearch_data

COPY ./build/sql/init.sql /docker-entrypoint-initdb.d/
