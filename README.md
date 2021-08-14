# architecting
Just a space to explore different microservice tools in go

## Services

* `user`: A go service abstracting mongo transactions with the `arch.user` collection. [Details here.](doc/user/README.md)

## Launching mongo

A real hack and slash script to spin up the database, probably should use something like `docker-compose`

TODO: Use `docker-compose`


### Usage

```bash
$ source launch-mongo.sh
```

This will:
* Export environment variables
* Make a docker network `architecting`
* Make a docker volume `architectingdbdata`
* Pull the docker image `mongo:5.0.2`
* Create a long lived container `architectingdb`
* Start container `architectingdb` if not already running
* Create an alias `dbconn` to exec into `architectingdb` using `mongosh`


### Environmet Vars Exported

```env
MONGO_CONTAINER_NAME=architectingdb
MONGO_DATA_DIR=$(docker volume inspect --format '{{ .Mountpoint }}' ${MONGO_DATA_VOLUME})
MONGO_DATA_VOLUME=architectingdbdata
MONGO_VERSION=5.0.2
PROJECT_NETWORK=architecting
```

## 2021-08-14

I want to explore the following:

* [x] mongodb containerization with docker
* [x] mongodb go driver
* [ ] liveness probe exposure from program (kube prep)
* [ ] 1:1 mapping between live service and mongo collection
* [ ] jsonrpc
* [ ] interprocess communication
* [ ] user-focused authz & authn