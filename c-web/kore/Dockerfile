FROM ubuntu:14.04

#ENV PGSQL=1
#ENV NOTLS=1

RUN mkdir /src
WORKDIR /src
VOLUME /src

RUN apt-get update; apt-get install -y wget make gcc libpq-dev

RUN cd /tmp
RUN wget https://github.com/jorisvink/kore/releases/download/2.0.0-release/kore-2.0.0-release.tar.gz
RUN tar -xvf kore-2.0.0-release.tar.gz 
RUN cd kore
RUN make; make install

ENTRYPOINT ["kore"]

