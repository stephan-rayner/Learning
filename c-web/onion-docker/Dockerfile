FROM ubuntu:14.04

ENV CMAKE_CXX_COMPILER gcc

RUN apt-get update
RUN apt-get install -y git make cmake gcc libpq-dev libevent-dev

RUN cmake -D CMAKE_C_COMPILER=$(which gcc)

RUN git clone https://github.com/davidmoreno/onion.git
RUN cd onion; mkdir build; cd build; cmake ..
#; make; make install

#RUN mkdir /src
#WORKDIR /src
#VOLUME /src
#CMD echo "CHEESE"
CMD which gcc
