FROM golang:1.11

RUN curl -sL https://deb.nodesource.com/setup_10.x | bash -
RUN apt-get update && apt-get install -y --no-install-recommends nodejs unzip && rm -rf /var/lib/apt/lists/*

RUN mkdir proto && cd proto && \
    wget https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip && \
    unzip protoc-3.6.1-linux-x86_64.zip && rm protoc-3.6.1-linux-x86_64.zip && \
    chmod +x bin/protoc

RUN go get -u github.com/gogo/protobuf/protoc-gen-gogo
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
RUN git clone https://github.com/googleapis/googleapis.git /googleapis

ADD scripts/* /scripts/
RUN cd /scripts && npm i

ADD hooks/* /hooks/
RUN npm i -g json2yaml

ENV PROTO_PATH /go/proto/include
ENV PATH $PATH:/go/proto/bin

WORKDIR /source

ENTRYPOINT ["/hooks/gen.sh"]