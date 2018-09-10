#!/bin/sh

for service in "$@"
do
    echo "Setting up Go workspace"
    /hooks/setupgo.sh || { echo 'Go workspace setup failed. Make sure prefix is defined in mmo config' ; exit 1; }
    echo "Generating gRPC gateway from protofile for service $service"
    cd /source/$service/protobuf
    protoc -I/usr/local/include -I. -I/source -I/googleapis -I${GOPATH}/src --grpc-gateway_out=logtostderr=true:${GOPATH}/src proto.proto
done