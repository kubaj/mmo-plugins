# Changelog

## 1.4.0

* `flowup/mmo-gen-go-grpc`
  * upgraded protoc to 3.6.1
* `flowup/mmo-gen-grpc-gateway`
  * upgraded protoc to 3.6.1
  * grpc-gateway upgraded to 1.5.0
* `flowup/mmo-gen-swagger`
  * upgraded protoc to 3.6.1
  * grpc-gateway upgraded to 1.5.0
* `flowup/mmo-gen-angular-client`
  * api-client-generator upgraded to version 4.0.1
  
## 1.3.0

* `flowup/mmo-gen-angular-client`
  * Angular client is based on the latest generator 

## 1.2.1

* `flowup/mmo-gen-angular-client`
  * Models are removed from project when removed from proto

## 1.0.2
* `flowup/mmo-gen-go-grpc`
  * gogo generator swapped for gofast due to the better compatiblity with grpc-gateway
* `flowup/mmo-gen-grpc-gateway`
  * Support of swagger-extra at root of project

## 1.0.1

* `flowup/mmo-gen-ignore-xxx`
  * Add to all tags `json:"-"` tag `sql:"-"`
* `flowup/mmo-gen-go-grpc`
  * Support of importing protofiles from different service
* `flowup/mmo-gen-grpc-gateway`
  * Support of importing protofiles from different service
* `flowup/mmo-gen-swagger`
  * Support of importing protofiles from different service
* `flowup/mmo-gen-proto-interface`
  * Support of service mock client generation (testability of other services)
