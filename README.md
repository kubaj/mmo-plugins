# MMO plugins repository

This repository contains all officially supported MMO plugins.

## MMO Plugin creation guide

Every MMO Plugin is a Docker container. It has following properties:

* Source code is mounted to `/source` folder
* Entrypoint is set to a program named `gen.sh` (generator)

Upon plugin invocation container is started and name of the service(s) are passed as command arguments to the entrypoint program. 

There are two types of the plugins - global and service plugins. Principle of functionality is same for both of types. Only difference is that all services are passed as arguments upon global plugin invocation. Upon single service plugins invocation is only one name of the service passed.

### Example

**Simple example of the MMO plugin that creates an empty file called `hello-mmo` in service.**

Create a script named `gen.sh` in the root of your plugin folder:
```bash
#!/bin/sh
touch $1
```

Don't forget to add execution permissions to the script:
```
$ chmod +x gen.sh
```

Create `Dockerfile` in your root afterwards:

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
