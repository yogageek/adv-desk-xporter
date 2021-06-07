#!/bin/bash

# 正式用
VERSION="1.0.2"
CONTAINER="ifps-xporter"

# docker location
DOCKER_REPO="iiicondor/$CONTAINER"
HARBOR_REPO="iiicondor/ifps-xporter"

docker build --network=host -t $DOCKER_REPO:$VERSION .
docker tag $DOCKER_REPO:$VERSION $HARBOR_REPO:$VERSION
# docker push $DOCKER_REPO:$VERSION
docker push $HARBOR_REPO:$VERSION

# docker repo 
# docker tag $DOCKER_REPO:$VERSION $DOCKER_REPO:latest
# docker push $DOCKER_REPO:latest

# harbor repo
# docker tag $DOCKER_REPO:$VERSION $HARBOR_REPO:latest
# docker push $HARBOR_REPO:latest

# docker rmi -f $(docker images | grep $CONTAINER | awk '{print $3}')

