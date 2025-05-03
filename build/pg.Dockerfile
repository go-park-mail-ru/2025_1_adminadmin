FROM postgres:15.0
RUN mkdir -p /usr/share/postgresql/15/tsearch_data
RUN chmod 777 /usr/share/postgresql/15/tsearch_data