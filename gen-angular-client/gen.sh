#!/bin/bash

rm -rf /source/ui/src/api/*
api-client-generator -s /source/swagger.yaml -o /source/ui/src/api -t all
