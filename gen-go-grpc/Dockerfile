FROM golang

RUN apt-get update && apt-get install -y --no-install-recommends unzip && rm -rf /var/lib/apt/lists/*

RUN mkdir proto && cd proto && \
    wget https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip && \
    unzip protoc-3.5.1-linux-x86_64.zip && rm protoc-3.5.1-linux-x86_64.zip && \
    chmod +x bin/protoc

RUN git clone https://github.com/googleapis/googleapis.git /googleapis
RUN go get -u github.com/gogo/protobuf/protoc-gen-gofast
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

ENV PROTO_PATH /go/proto/include
ENV PATH $PATH:/go/proto/bin

ADD hooks/* /hooks/

ENTRYPOINT ["/hooks/gen.sh"]