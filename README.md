# MMO ecosystem repository

This repository will contain all MMO stuff that isn't supposed to be in the core repository

## MMO Plugin creation guide

MMO Plugin is Docker container. It meets following properties:

* Source code is mount to `/source`
* Entrypoint is set to program that does generating

Upon plugin invocation container is started and name of the service(s) are passed as command arguments to the entrypoint program. 

There are two types of the plugins - global and service plugins. Principle of functionality is same for both of types. Only difference is that all services are passed as arguments upon global plugin invocation. Upon single service plugins invocation is only one name of the service passed.

### Example

Simple example of the MMO plugin. Plugin will create empty file `hello-mmo` in service.

Generation script named `gen.sh`:

```bash
#!/bin/sh
touch $1
```

Don't forget add execution permission to script:

```
$ chmod +x gen.sh
```

Then we will create `Dockerfile`:

```Dockerfile
FROM alpine:latest

ADD gen.sh /mmo/gen.sh
ENTRYPOINT /mmo/gen.sh
```

Docker image is built in the last step (`-t` is name of the plugin):
```
$ docker build -t hello-mmo .
```

After acomplishing these step we can put `hello-mmo:latest` to the `plugins` section of some service in the mmo.yaml. Image can be pushed, eg. to Docker hub: 

```
$ docker push hello-mmo:latest
```