FROM postgres:latest
RUN mkdir -p /usr/share/postgresql/17/tsearch_data
RUN chmod 777 /usr/share/postgresql/17/tsearch_data