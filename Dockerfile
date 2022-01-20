# syntax=docker/dockerfile:1
FROM golang:latest as build

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . /app
RUN go build ./main.go

FROM ubuntu:latest

RUN apt-get -y update && apt-get install -y tzdata
ENV TZ=Russia/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV PGVER 12
RUN apt-get -y update && apt-get install -y postgresql-$PGVER
USER postgres


RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER a_shirshov WITH SUPERUSER PASSWORD 'password';" &&\
    createdb -O a_shirshov bd_tp_V2 &&\
    /etc/init.d/postgresql stop

RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/$PGVER/main/pg_hba.conf
RUN echo "listen_addresses='*'\n\
synchronous_commit = off\n\
fsync = off\n\
shared_buffers = 256MB\n\
effective_cache_size = 1024MB\n\
full_page_writes = off\n\
fsync = off" >> /etc/postgresql/$PGVER/main/postgresql.conf

VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

WORKDIR /app
COPY . .
COPY --from=build /app/main .

USER root
ENV PGPASSWORD password
CMD service postgresql start && psql -h localhost -d bd_tp_V2 -U a_shirshov -p 5432 -a -q -f ./db/db.sql && ./main