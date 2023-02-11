FROM postgres:15.1
COPY database/init.sql /docker-entrypoint-initdb.d/
ENV POSTGRES_USER="daymood"\
    POSTGRES_PASSWORD="daymood"