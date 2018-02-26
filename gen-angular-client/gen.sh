#!/bin/bash

rm /source/ui/api/models/*
api-client-generator -s /source/swagger.yaml -o /source/ui/api
