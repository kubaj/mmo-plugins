#!/bin/bash

rm /source/ui/api/models/*
ngx-swag-client -s /source/swagger.yaml -o /source/ui/api