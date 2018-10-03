#!/bin/bash

CWD=$PWD
LOG=${CWD}/../log
MONGODB_DATA=${CWD}/../mongodb


if [ ! -d "${LOG}" ]; then
    mkdir ${LOG}
fi

if [ ! -d "${MONGODB_DATA}" ]; then
    mkdir ${MONGODB_DATA}
fi


if [ $(docker network ls --filter "name=app-net" | wc -l) -eq 1 ]; then
    echo `date -u +"%Y-%m-%d %H:%M:%S"` "create a docker network for communicate between app-server and mongodb"
    docker network create app-net
fi
docker network ls --filter "name=app-net"


if [ $(docker container ls --filter "name=mongodb" | wc -l) -eq 1 ]; then
    echo `date -u +"%Y-%m-%d %H:%M:%S"` "start container mongodb"
    docker run --name mongodb --network app-net --publish 27017:27017 -v ${MONGODB_DATA}:/data/db -itd mongo:3.2
fi
docker container ls --filter "name=mongodb"


echo `date -u +"%Y-%m-%d %H:%M:%S"` "check repository changes"
while true; do
    git fetch
    if [ $(git rev-parse HEAD) != $(git rev-parse @{u}) ]; then
        git reset --hard
        git pull

        IMAGE_TAG=`date -u +"%Y%m%d%H%M%S"`
        docker build -t todolist:${IMAGE_TAG} --rm --no-cache -f docker/Dockerfile .

        # Run the unit test
        docker container stop app-test
        docker container rm -vfl app-test
        docker run --name app-test -it todolist:${IMAGE_TAG} go test ./... > ${LOG}/app-test-${IMAGE_TAG}.log
        docker container stop app-test
        docker container rm -vfl app-test

        # If tests are OK then start the new version container
        if [ $(cat ${LOG}/app-test-${IMAGE_TAG}.log | grep -i "FAIL" | wc -l) -eq 0 ]; then
            docker container stop app-server
            docker container rm -vfl app-server
            docker tag todolist:${IMAGE_TAG} todolist:latest
            docker run --name app-server --network app-net --publish 8181:8181 -itd todolist:latest go run -v app/app_server.go -mongodbUrl mongodb:27017 > ${LOG}/app-server.${IMAGE_TAG}.log
        else
            docker image rm -f todolist:${IMAGE_TAG}
        fi
    else
        echo `date -u +"%Y-%m-%d %H:%M:%S"` "no changes on repository, sleep a while"
        sleep 60
    fi
done