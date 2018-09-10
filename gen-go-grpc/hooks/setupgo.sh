#!/bin/bash

if [ -z "$GO_PREFIX" ]; then
    exit 1
fi

SHADOW_WS=$GOPATH/src/$GO_PREFIX
LINK_NAME=$(dirname "$SHADOW_WS")

mkdir -p ${LINK_NAME}

ln  -s /source "$SHADOW_WS" -T