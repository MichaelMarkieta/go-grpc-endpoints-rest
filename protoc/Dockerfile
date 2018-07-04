FROM golang:alpine

# Update and install dependencies
RUN apk update && apk add git && apk add protobuf autoconf automake libtool curl make g++ unzip

# Install protobuf
RUN git clone https://github.com/google/protobuf.git
WORKDIR protobuf
RUN git submodule update --init --recursive
RUN ./autogen.sh
RUN ./configure
RUN make
RUN make install
RUN ldconfig /; echo \$?

ENTRYPOINT ["bash"]
