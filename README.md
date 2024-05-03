# taupatupatu

### How to build image and run single service

To build and run single service as docker container, navigate to service folder and run make commands.
For example, to run auth service:

```
cd ./auth
make  -f ./deployments/Makefile build-image
make  -f ./deployments/Makefile run-service
```

### How to build and start all services with docker-compose

To build and start all services with `docker-compose` navigate to `deployments` directory in repository root and run

```
docker-compose build
docker-compose up
```
