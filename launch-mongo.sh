#!/bin/bash

set -e # Halt on errors

export MONGO_DATA_VOLUME=architectingdbdata
export MONGO_VERSION=5.0.2
export MONGO_CONTAINER_NAME=architectingdb
export PROJECT_NETWORK=architecting

# ensure docker volume

function ensureProjectNetworkCreated() {
    if ! docker network inspect ${PROJECT_NETWORK} &> /dev/null; then
        echo "${PROJECT_NETWORK} does not exist, creating ..."
        docker network create ${PROJECT_NETWORK} &> /dev/null
        echo "${PROJECT_NETWORK} created!"
    fi
}

function ensureMongoVolumeCreated() {
    if ! docker volume inspect ${MONGO_DATA_VOLUME} &> /dev/null; then
        echo "${MONGO_DATA_VOLUME} does not exist, creating ..."
        docker volume create ${MONGO_DATA_VOLUME} &> /dev/null
        echo "${MONGO_DATA_VOLUME} created!"
    fi
}

function ensureMongoImagePulled() {
    if ! docker image inspect mongo:${MONGO_VERSION} &> /dev/null; then
        echo "Pulling mongo:${MONGO_VERSION} ..."
        docker pull mongo:${MONGO_VERSION}
        echo "Pull complete!"
    fi
}

function ensureMongoContainerCreated() {
    if ! docker container inspect ${MONGO_CONTAINER_NAME} &> /dev/null; then
        echo "Created container ${MONGO_CONTAINER_NAME} ..."
        docker container create \
            --volume ${MONGO_DATA_VOLUME}:/data/db \
            --name ${MONGO_CONTAINER_NAME} \
            --publish 27117:27017 \
            --network ${PROJECT_NETWORK} \
            mongo:$MONGO_VERSION \
            &> /dev/null
        echo "${MONGO_CONTAINER_NAME} created!" 
    fi
}

function ensureMongoContainerRunning() {
    if [ $(docker container inspect --format '{{ .State.Running }}' ${MONGO_CONTAINER_NAME}) != "true" ]; then
        docker start ${MONGO_CONTAINER_NAME} &> /dev/null
        echo "${MONGO_CONTAINER_NAME} started!"
    else
        echo "${MONGO_CONTAINER_NAME} already running!"
    fi
}

ensureProjectNetworkCreated
ensureMongoVolumeCreated
export MONGO_DATA_DIR=$(docker volume inspect --format '{{ .Mountpoint }}' ${MONGO_DATA_VOLUME})
ensureMongoImagePulled
ensureMongoContainerCreated
ensureMongoContainerRunning

alias dbconn="docker exec -it ${MONGO_CONTAINER_NAME} mongosh"
echo "Run '\$ dbconn' to connect to mongo"

set +e // Return to normal shell error handling