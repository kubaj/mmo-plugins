#!/bin/bash

for service in "$@"
do
    echo "Generating gRPC server and stubs from protofile for service $service"
    cd /source/$service/protobuf
    protoc -I/usr/local/include -I. -I/googleapis -I${GOPATH}/src --gogo_out=plugins=grpc:/source/$service proto.proto
    protoc -I/usr/local/include -I. -I/googleapis -I${GOPATH}/src --include_imports --include_source_info --descriptor_set_out /source/$service/protobuf/descriptor.pb proto.proto
done